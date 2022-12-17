package main

import (
	"context"
	"encoding/json"
	"github.com/cloudwego/hertz/pkg/protocol"

	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/app/middlewares/client/sd"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	nacos_demo "github.com/longlihale/nacos-demo"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

type Test struct {
	A int `json:"a"`
	B int `json:"b"`
}

func main() {
	cli, err := client.NewClient()
	if err != nil {
		panic(err)
	}
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

	naocsCli, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		})
	if err != nil {
		panic(err)
	}
	r := nacos_demo.NewNacosResolver(naocsCli)
	cli.Use(sd.Discovery(r))
	for i := 0; i < 10; i++ {
		// set request config、method、request uri.
		req := protocol.AcquireRequest()
		req.SetOptions(config.WithSD(true))
		req.SetMethod("POST")
		req.SetRequestURI("http://hertz.test.demo/ping")
		t := Test{A: i, B: i}
		bytes, _ := json.Marshal(t)
		// set body and content type
		req.SetBody(bytes)
		req.Header.SetContentTypeBytes([]byte("application/json"))
		resp := protocol.AcquireResponse()
		// send request
		err := cli.Do(context.Background(), req, resp)
		if err != nil {
			hlog.Fatal(err)
		}
		hlog.Infof("code=%d,body=%s", resp.StatusCode(), string(resp.Body()))
	}
}
