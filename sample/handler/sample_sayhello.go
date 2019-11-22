package handler

import (
	"context"

	zeusctx "gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/context"

	gomicro "zeus-examples/sample/proto/samplepb"
	"zeus-examples/sampleservice/logic/user"
)


func (h *Sample) SayHello(ctx context.Context, req *gomicro.Request, rsp *gomicro.Reply) (err error) {
	logger := zeusctx.ExtractLogger(ctx)
	logger.Debug("SayHello")
	info, err := user.GetInfo(ctx, "001")
	if err != nil {
		logger.Error("user.GetInfo err:", err)
		return
	}
	rsp.Message = "hello world, " + info + "."
	return
}
