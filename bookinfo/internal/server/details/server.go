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

package details

import (
	"context"
	"net"

	"code.byted.org/kite/kitex/pkg/rpcinfo"
	"github.com/cloudwego/biz-demo/bookinfo/kitex_gen/cwg/bookinfo/details"
	"github.com/cloudwego/biz-demo/bookinfo/kitex_gen/cwg/bookinfo/details/detailsservice"
	"github.com/cloudwego/biz-demo/bookinfo/pkg/constants"
	"github.com/cloudwego/biz-demo/bookinfo/pkg/utils"
	"github.com/cloudwego/biz-demo/bookinfo/pkg/utils/logutils"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"github.com/kitex-contrib/registry-nacos/nacos"
	"github.com/kitex-contrib/registry-nacos/registry"
)

// Server kitex server
type Server struct {
	opts *ServerOptions
	svc  details.DetailsService
}

// ServerOptions server options
type ServerOptions struct {
	Addr          string         `mapstructure:"addr"`
	LogLevel      logutils.Level `mapstructure:"logLevel"`
	EnableNacos   bool           `mapstructure:"enableNacos"`
	EnableTracing bool           `mapstructure:"enableTracing"`
}

// DefaultServerOptions default opts
func DefaultServerOptions() *ServerOptions {
	return &ServerOptions{
		Addr:     ":8084",
		LogLevel: logutils.LevelInfo,
	}
}

// Run kitex server
func (s *Server) Run(ctx context.Context) error {
	klog.SetLogger(kitexlogrus.NewLogger())
	klog.SetLevel(s.opts.LogLevel.KitexLogLevel())

	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(constants.DetailsServiceName),
		provider.WithInsecure(),
	)
	defer func(p provider.OtelProvider, ctx context.Context) {
		_ = p.Shutdown(ctx)
	}(p, ctx)

	addr, err := net.ResolveTCPAddr("tcp", s.opts.Addr)
	if err != nil {
		klog.Fatal(err)
	}
	opts := []server.Option{
		server.WithServiceAddr(addr),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: constants.DetailsServiceName,
			Tags:        utils.ExtractInstanceMeta(),
		}),
	}

	if s.opts.EnableNacos {
		cli, err := nacos.NewDefaultNacosClient(utils.WithAuth())
		if err != nil {
			klog.Fatal(err)
		}
		opts = append(opts, server.WithRegistry(registry.NewNacosRegistry(cli)))
	}
	if s.opts.EnableTracing {
		opts = append(opts, server.WithSuite(tracing.NewServerSuite()))
	}

	klog.Infof("init nacos registry client completed.")
	svr := detailsservice.NewServer(s.svc, opts...)
	if err := svr.Run(); err != nil {
		klog.Fatal(err)
	}

	return nil
}
