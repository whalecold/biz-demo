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
	"github.com/cloudwego/biz-demo/bookinfo/kitex_gen/cwg/bookinfo/details/detailsservice"
	"github.com/cloudwego/biz-demo/bookinfo/pkg/constants"
	kclient "github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	xdsmanager "github.com/kitex-contrib/xds"
	"github.com/kitex-contrib/xds/xdssuite"
)

// DetailsClientOptions detail client options
type DetailsClientOptions struct {
	Endpoint  string `mapstructure:"endpoint"`
	EnableXDS bool   `mapstructure:"enableXDS"`
	XDSAddr   string `mapstructure:"xdsAddr"`
	XDSAuth   bool   `mapstructure:"xdsAuth"`
}

// DefaultDetailsClientOptions default options
func DefaultDetailsClientOptions() *DetailsClientOptions {
	return &DetailsClientOptions{
		Endpoint:  "details:8084",
		EnableXDS: false,
		XDSAddr:   "istiod.istio-system.svc:15012",
		XDSAuth:   true,
	}
}

// ProvideDetailsClient provide details client
// 1、init xds manager: only init once
// 2、enable xds
// 3、enable opentelemetry
func ProvideDetailsClient(opts *DetailsClientOptions) (detailsservice.Client, error) {
	if opts.EnableXDS {
		if err := xdsmanager.Init(); err != nil {
			klog.Fatal(err)
		}
		return detailsservice.NewClient(
			opts.Endpoint,
			kclient.WithSuite(tracing.NewClientSuite()),
			kclient.WithSuite(xdssuite.NewClientSuite(xdssuite.WithRouterMetaExtractor(tracing.ExtractFromPropagator))),
		)
	}

	return detailsservice.NewClient(
		constants.DetailsServiceName,
		kclient.WithHostPorts(opts.Endpoint),
		kclient.WithSuite(tracing.NewClientSuite()),
	)
}
