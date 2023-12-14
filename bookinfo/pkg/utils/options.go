package utils

import (
	"os"

	"github.com/kitex-contrib/registry-nacos/nacos"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

// WithAuth sets the username .
func WithAuth() nacos.Option {
	return nacos.Option{F: func(param *vo.NacosClientParam) {
		if param != nil && param.ClientConfig != nil {
			param.ClientConfig.Username = os.Getenv("NACOS_ENV_USERNAME")
			param.ClientConfig.Password = os.Getenv("NACOS_ENV_PASSWARD")
		}
	}}
}
