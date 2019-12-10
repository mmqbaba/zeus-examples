package rpcclient

import (
	"context"
	"sync"

	"github.com/micro/go-micro/client"

	zeusctx "gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/context"

	gomicro "zeus-examples/hellodemo/proto/hellodemopb"
)

type HelloDemoService struct {
	gomicro.HelloDemoService
	once sync.Once
	name string
}

// HelloDemoSrv
var hellodemoSrv HelloDemoService

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
