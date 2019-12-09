package rpcclient

import (
	"context"
	"sync"

	"github.com/micro/go-micro/client"

	zeusctx "gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/context"

	hellodemo "zeus-examples/hellodemo/proto/hellodemopb"
	gomicro "zeus-examples/sample/proto/samplepb"
)

var cli client.Client

type SampleService struct {
	mux    sync.RWMutex
	name   string
	client gomicro.SampleService
}

// sampleSrv
var sampleSrv SampleService

func NewSampleService(ctx context.Context) (gomicro.SampleService, error) {
	sampleSrv.mux.RLock()
	if sampleSrv.client != nil {
		defer sampleSrv.mux.RUnlock()
		return sampleSrv.client, nil
	}
	sampleSrv.mux.RUnlock()

	sampleSrv.mux.Lock()
	defer sampleSrv.mux.Unlock()
	cli, err := zeusctx.ExtractGMClient(ctx)
	if err != nil {
		return nil, err
	}
	sampleSrv.name = "sample"
	sampleSrv.client = gomicro.NewSampleService(sampleSrv.name, cli)
	return sampleSrv.client, nil
}

type helloService struct {
	mux    sync.RWMutex
	name   string
	client hellodemo.HelloDemoService
}

// HelloSrv
var helloSrv helloService

func NewHelloService(ctx context.Context) (hellodemo.HelloDemoService, error) {
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
	helloSrv.name = "hellodemo"
	helloSrv.client = hellodemo.NewHelloDemoService(helloSrv.name, cli)
	return helloSrv.client, nil
}
