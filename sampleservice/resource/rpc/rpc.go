package rpc

import (
	"context"
	"sync"

	"github.com/micro/go-micro/client"

	"gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/microsrv/gomicro"

	"zeus-examples/sampleservice/global"
	hello "zeus-examples/sampleservice/proto/gomicro"
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
	cli, err := gomicro.NewClient(ctx, global.GetConfig().GoMicro)
	if err != nil {
		return nil, err
	}
	helloSrv.name = "zeus"
	helloSrv.client = hello.NewHelloService("zeus", cli)
	return helloSrv.client, nil
}
