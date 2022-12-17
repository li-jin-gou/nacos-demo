package main

import (
	"context"
	"sync"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	nacos_demo "github.com/longlihale/nacos-demo"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

var wg sync.WaitGroup

type Test struct {
	A int `json:"a"`
	B int `json:"b"`
}

func main() {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig("127.0.0.1", 8848),
	}

	cc := constant.ClientConfig{
		NamespaceId:         "public",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "info",
	}

	cli, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		panic(err)
	}
	wg.Add(2)
	go func() {
		defer wg.Done()
		addr := "127.0.0.1:8888"
		r := nacos_demo.NewNacosRegistry(cli)
		h := server.Default(
			server.WithHostPorts(addr),
			server.WithRegistry(r, &registry.Info{
				ServiceName: "hertz.test.demo",
				Addr:        utils.NewNetAddr("tcp", addr),
				Weight:      10,
				Tags:        nil,
			}))

		h.POST("/ping", func(c context.Context, ctx *app.RequestContext) {
			t := Test{}
			if err := ctx.Bind(&t); err != nil {
				ctx.String(consts.StatusOK, err.Error())
				return
			}
			ctx.JSON(consts.StatusOK, t)
		})
		h.Spin()
	}()
	go func() {
		defer wg.Done()
		addr := "127.0.0.1:8889"
		r := nacos_demo.NewNacosRegistry(cli)
		h := server.Default(
			server.WithHostPorts(addr),
			server.WithRegistry(r, &registry.Info{
				ServiceName: "hertz.test.demo",
				Addr:        utils.NewNetAddr("tcp", addr),
				Weight:      10,
				Tags:        nil,
			}))
		h.POST("/ping", func(c context.Context, ctx *app.RequestContext) {
			t := Test{}
			if err := ctx.Bind(&t); err != nil {
				ctx.String(consts.StatusOK, err.Error())
				return
			}
			ctx.JSON(consts.StatusOK, t)
		})
		h.Spin()
	}()

	wg.Wait()
}
