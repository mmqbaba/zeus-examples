package handler

import (
	"context"

	zeusctx "gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/context"

	// "zeus-examples/sample/logic/user"
	gomicro "zeus-examples/sample/proto/samplepb"
)

func (h *Sample) SayHello(ctx context.Context, req *gomicro.Request, rsp *gomicro.Reply) (err error) {
	logger := zeusctx.ExtractLogger(ctx)
	logger.Debug("SayHello")

	// info, err := user.GetInfo(ctx, "001")
	// if err != nil {
	// 	logger.Error("user.GetInfo err:", err)
	// 	return
	// }
	// rsp.Message = "hello world, " + info + "."

	rsp.Message = "hello world, " + req.Name + "."
	return
}