package rpcclient
import (
	"context"
	"sync"

	"github.com/micro/go-micro/client"

	zeusctx "gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/context"

	gomicro "zeus_app/hello/proto/gomicro"
)

var cli client.Client

type helloService struct {
	mux    sync.RWMutex
	name   string
	client gomicro.HelloService
}

// helloSrv
var helloSrv helloService

func NewHelloService(ctx context.Context) (gomicro.HelloService, error) {
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
	helloSrv.client = gomicro.NewHelloService(helloSrv.name, cli)
	return helloSrv.client, nil
}

