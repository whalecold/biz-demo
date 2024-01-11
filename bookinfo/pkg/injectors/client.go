package injectors

import (
	"context"
	"time"

	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/biz-demo/bookinfo/pkg/metadata"
	kclient "github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/xds"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	xdsmanager "github.com/kitex-contrib/xds"
	"github.com/kitex-contrib/xds/xdssuite"
	"go.opentelemetry.io/otel"

	"github.com/cloudwego/kitex-examples/kitex_gen/api"
	"github.com/cloudwego/kitex-examples/kitex_gen/api/echo"
)

func ExtractFromPropagator(ctx context.Context) map[string]string {
	metadata := metainfo.GetAllValues(ctx)
	if metadata == nil {
		metadata = make(map[string]string)
	}
	otel.GetTextMapPropagator().Inject(ctx, &metadataProvider{metadata: metadata})
	return metadata
}

func main() {

	// 初始化 xds manager
	if err := xdsmanager.Init(); err != nil {
		klog.Fatal(err)
	}

	// 使用 xds 扩展初始化 kitex client
	echoClient, err := echo.NewClient(
		"echo:80",
		// tracing 扩展，用于染色信息透传
		kclient.WithSuite(tracing.NewClientSuite()),
		kclient.WithXDSSuite(xds.ClientSuite{
			// 使用 xds 路由
			RouterMiddleware: xdssuite.NewXDSRouterMiddleware(
				xdssuite.WithRouterMetaExtractor(metadata.ExtractFromPropagator),
			),
			Resolver: xdssuite.NewXDSResolver(),
		}),
	)
	if err != nil {
		panic(err)
	}

	for {
		req := &api.Request{Message: "my request"}
		resp, err := echoClient.Echo(context.Background(), req)
		if err != nil {
			klog.Errorf("take request error: %v", err)
		} else {
			klog.Infof("receive response %v", resp)
		}
		time.Sleep(time.Second * 10)
	}
}
