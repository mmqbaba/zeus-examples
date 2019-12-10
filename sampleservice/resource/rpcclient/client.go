package rpc

import (
	"context"
	"sync"

	"github.com/micro/go-micro/client"

	zeusctx "gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/context"

	"zeus-examples/sampleservice/proto/hello"
)

type HelloService struct {
	hello.HelloService
	once sync.Once
	name string
}

// HelloSrv
var helloSrv HelloService

func NewHelloService(ctx context.Context) (hello.HelloService, error) {
	var err error
	helloSrv.once.Do(func() {
		var cli client.Client
		cli, err = zeusctx.ExtractGMClient(ctx)
		if err != nil {
			return
		}
		helloSrv.name = "hello"
		helloSrv.HelloService = hello.NewHelloService(helloSrv.name, cli)
	})
	if err != nil {
		return nil, err
	}
	return &helloSrv, nil
}
