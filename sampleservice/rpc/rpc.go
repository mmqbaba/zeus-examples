package rpc

import (
	"context"

	"github.com/micro/go-micro/server"
	"github.com/sirupsen/logrus"

	zeusctx "gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/context"
	"gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/engine"
	"gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/service"
	"gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/microsrv/gomicro"

	"zeus-examples/sampleservice/global"
)

func init() {
	// gomicro
	global.ServiceOpts = append(global.ServiceOpts, service.WithGoMicroServerWrapGenerateFnOption(customerServerWrap))
	global.ServiceOpts = append(global.ServiceOpts, service.WithGoMicroClientWrapGenerateFnOption(gomicro.GenerateClientWrapTest))
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
