package rpc

import (
	"context"
	"sync"

	"github.com/micro/go-micro/client"

	zeusctx "gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/context"

	"zeus-examples/sampleservice/proto/hello"
)

var cli client.Client

type helloService struct {
	mux    sync.RWMutex
	name   string
	client hello.HelloService
}

// HelloSrv
var helloSrv helloService

func GetHelloService(ctx context.Context) (hello.HelloService, error) {
	helloSrv.mux.RLock()
	if helloSrv.client != nil {
		defer helloSrv.mux.RUnlock()
		return helloSrv.client, nil
	}
	helloSrv.mux.RUnlock()

	helloSrv.mux.Lock()
	defer helloSrv.mux.Unlock()
	cli, err := zeusctx.ExtractGMClient(ctx)
	if err != nil {
		return nil, err
	}
	helloSrv.name = "zeus"
	helloSrv.client = hello.NewHelloService(helloSrv.name, cli)
	return helloSrv.client, nil
}
