package rpcclient
import (
	"context"
	"sync"

	"github.com/micro/go-micro/client"

	zeusctx "gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/context"

	gomicro "zeus-examples/hellodemo/proto/hellodemopb"
)

var cli client.Client

type HelloDemoService struct {
	mux    sync.RWMutex
	name   string
	client gomicro.HelloDemoService
}

// hellodemoSrv
var hellodemoSrv HelloDemoService

func NewHelloDemoService(ctx context.Context) (gomicro.HelloDemoService, error) {
	hellodemoSrv.mux.RLock()
	if hellodemoSrv.client != nil {
		defer hellodemoSrv.mux.RUnlock()
		return hellodemoSrv.client, nil
	}
	hellodemoSrv.mux.RUnlock()

	hellodemoSrv.mux.Lock()
	defer hellodemoSrv.mux.Unlock()
	cli, err := zeusctx.ExtractGMClient(ctx)
	if err != nil {
		return nil, err
	}
	hellodemoSrv.name = "hellodemo"
	hellodemoSrv.client = gomicro.NewHelloDemoService(hellodemoSrv.name, cli)
	return hellodemoSrv.client, nil
}

