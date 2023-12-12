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

package injectors

import (
	"time"

	"github.com/cloudwego/biz-demo/bookinfo/kitex_gen/cwg/bookinfo/ratings/ratingservice"
	"github.com/cloudwego/biz-demo/bookinfo/pkg/constants"
	"github.com/cloudwego/biz-demo/bookinfo/pkg/metadata"
	kclient "github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/xds"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	xdsmanager "github.com/kitex-contrib/xds"
	"github.com/kitex-contrib/xds/core/manager"
	"github.com/kitex-contrib/xds/xdssuite"
)

// RatingsClientOptions client options
type RatingsClientOptions struct {
	Endpoint  string `mapstructure:"endpoint"`
	EnableXDS bool   `mapstructure:"enableXDS"`
	XDSAddr   string `mapstructure:"xdsAddr"`
	XDSAuth   bool   `mapstructure:"xdsAuth"`
}

func DefaultRatingsClientOptions() *RatingsClientOptions {
	return &RatingsClientOptions{
		Endpoint:  "ratings:8083",
		EnableXDS: false,
		XDSAddr:   "istiod.istio-system.svc:15012",
		XDSAuth:   true,
	}
}

// ProvideRatingsClient provide ratings client
// 1、init xds manager: only init once
// 2、enable xds
// 3、enable opentelemetry
func ProvideRatingsClient(opts *RatingsClientOptions) (ratingservice.Client, error) {
	if opts.EnableXDS {
		if err := xdsmanager.Init(
			xdsmanager.WithXDSServerConfig(&manager.XDSServerConfig{
				SvrName: constants.IstiodSvrName,
				SvrAddr: opts.XDSAddr,
				XDSAuth: opts.XDSAuth,
			}),
		); err != nil {
			klog.Fatal(err)
		}
		return ratingservice.NewClient(
			opts.Endpoint,
			kclient.WithSuite(tracing.NewClientSuite()),
			kclient.WithRPCTimeout(time.Second*10),
			kclient.WithXDSSuite(xds.ClientSuite{
				RouterMiddleware: xdssuite.NewXDSRouterMiddleware(
					xdssuite.WithRouterMetaExtractor(metadata.ExtractFromPropagator),
				),
				Resolver: xdssuite.NewXDSResolver(),
			}),
		)
	}

	return ratingservice.NewClient(
		constants.RatingsServiceName,
		kclient.WithHostPorts(opts.Endpoint),
		kclient.WithSuite(tracing.NewClientSuite()),
	)
}
