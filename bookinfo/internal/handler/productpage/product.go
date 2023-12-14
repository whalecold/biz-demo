// Copyright 2022 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package productpage

import (
	"context"
	"net/http"
	"strings"

	"github.com/cloudwego/biz-demo/bookinfo/kitex_gen/base"
	"github.com/cloudwego/biz-demo/bookinfo/kitex_gen/cwg/bookinfo/details"
	"github.com/cloudwego/biz-demo/bookinfo/kitex_gen/cwg/bookinfo/details/detailsservice"
	"github.com/cloudwego/biz-demo/bookinfo/kitex_gen/cwg/bookinfo/reviews"
	"github.com/cloudwego/biz-demo/bookinfo/kitex_gen/cwg/bookinfo/reviews/reviewsservice"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/pkg/klog"
	"golang.org/x/sync/errgroup"
)

// Handler productpage hertz handler
type Handler struct {
	reviewsClient reviewsservice.Client
	detailsClient detailsservice.Client
}

// New create handler
func New(reviewsClient reviewsservice.Client, detailsClient detailsservice.Client) *Handler {
	return &Handler{reviewsClient: reviewsClient, detailsClient: detailsClient}
}

// GetProduct get product info
func (h *Handler) GetProduct(ctx context.Context, c *app.RequestContext) {
	productID := c.Param("productID")
	output := c.Query("output")

	var (
		reviewsResp *reviews.ReviewResp
		detailsResp *details.GetProductResp
	)

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		res, err := h.reviewsClient.ReviewProduct(ctx, &reviews.ReviewReq{ProductID: productID})
		if err != nil {
			klog.CtxErrorf(ctx, "call reviews error: %s", err.Error())
			return err
		}
		reviewsResp = res
		return nil
	})

	eg.Go(func() error {
		res, err := h.detailsClient.GetProduct(ctx, &details.GetProductReq{ID: productID})
		if err != nil {
			klog.CtxErrorf(ctx, "call details error: %s", err.Error())
			return err
		}

		detailsResp = res
		return nil
	})

	if err := eg.Wait(); err != nil {
		c.JSON(http.StatusInternalServerError, &base.BaseResp{
			StatusMessage: err.Error(),
			StatusCode:    http.StatusInternalServerError,
			Extra:         nil,
		})
		return
	}

	if output == "html" {
		c.HTML(http.StatusOK, "productpage.html", buildData(productID, reviewsResp, detailsResp))
		return
	}
	c.JSON(http.StatusOK, buildData(productID, reviewsResp, detailsResp))
}

func buildData(id string, reviews *reviews.ReviewResp, detailResp *details.GetProductResp) *productData {
	tr := &productData{
		ID:              id,
		Title:           detailResp.GetProduct().GetTitle(),
		Author:          detailResp.GetProduct().GetAuthor(),
		Description:     detailResp.GetProduct().GetDescription(),
		FullStarts:      map[int8]string{},
		EmptyStarts:     map[int8]string{},
		Color:           strings.ToLower(reviews.GetReview().GetType().String()),
		ReviewsInstance: reviews.GetReview().GetReviewsInstance(),
		RatingsInstance: reviews.GetReview().GetRatingsInstance(),
		Link:            detailResp.GetProduct().GetLink(),
		AuthorLink:      detailResp.GetProduct().GetAuthorLink(),
	}
	for i := int8(0); i < reviews.GetReview().GetRating(); i++ {
		tr.FullStarts[i] = "*"
	}
	for i := int8(5); i > reviews.GetReview().GetRating(); i-- {
		tr.EmptyStarts[i] = "*"
	}
	return tr
}

type productData struct {
	ReviewsInstance string
	RatingsInstance string
	ID              string
	Color           string
	FullStarts      map[int8]string
	EmptyStarts     map[int8]string
	Author          string
	Description     string
	Title           string
	Link            string
	AuthorLink      string
}
