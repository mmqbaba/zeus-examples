package rpcclient

import (
	"context"
	"sync"

	"github.com/micro/go-micro/client"

	zeusctx "github.com/mmqbaba/zeus/context"

	gomicro "zeus-examples/hellodemo/proto/hellodemopb"
)

var hellodemoSrv HelloDemoService

type HelloDemoService struct {
	gomicro.HelloDemoService
	once sync.Once
	name string
}

func NewHelloDemoService(ctx context.Context) (gomicro.HelloDemoService, error) {
	var err error
	hellodemoSrv.once.Do(func() {
		var cli client.Client
		cli, err = zeusctx.ExtractGMClient(ctx)
		if err != nil {
			return
		}
		hellodemoSrv.name = "hellodemo"
		hellodemoSrv.HelloDemoService = gomicro.NewHelloDemoService(hellodemoSrv.name, cli)
	})
	if err != nil {
		return nil, err
	}
	return &hellodemoSrv, nil
}
