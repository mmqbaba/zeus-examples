package rpcclient

import (
	"context"
	"sync"

	"github.com/micro/go-micro/client"

	zeusctx "gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/context"

	gomicro "zeus-examples/sample/proto/samplepb"
)

type SampleService struct {
	gomicro.SampleService
	once sync.Once
	name string
}

// sampleSrv
var sampleSrv SampleService

func NewSampleService(ctx context.Context) (gomicro.SampleService, error) {
	var err error
	sampleSrv.once.Do(func() {
		var cli client.Client
		cli, err = zeusctx.ExtractGMClient(ctx)
		if err != nil {
			return
		}
		sampleSrv.name = "sample"
		sampleSrv.SampleService = gomicro.NewSampleService(sampleSrv.name, cli)
	})
	if err != nil {
		return nil, err
	}
	return &sampleSrv, nil
}
