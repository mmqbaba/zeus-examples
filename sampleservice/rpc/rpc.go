package rpc

import (
	"context"

	"github.com/micro/go-micro/server"
	"github.com/sirupsen/logrus"

	zeusctx "github.com/mmqbaba/zeus/context"
	"github.com/mmqbaba/zeus/engine"
	zgomicro "github.com/mmqbaba/zeus/microsrv/gomicro"
	"github.com/mmqbaba/zeus/service"

	"github.com/mmqbaba/zeus-examples/sampleservice/global"
)

func init() {
	// gomicro
	global.ServiceOpts = append(global.ServiceOpts, service.WithGoMicroServerWrapGenerateFnOption(customerServerWrap))
	global.ServiceOpts = append(global.ServiceOpts, service.WithGoMicroClientWrapGenerateFnOption(zgomicro.GenerateClientWrapTest))
}

func customerServerWrap(ng engine.Engine) func(fn server.HandlerFunc) server.HandlerFunc {
	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) (err error) {
			logger := zeusctx.ExtractLogger(ctx)
			logger.WithFields(logrus.Fields{
				"tag": "gomicro-serverlogwrap-customer",
			}).Debug("====== test customer server wrap ======")
			err = fn(ctx, req, rsp)
			return
		}
	}
}
