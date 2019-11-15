package handler

import (
	"context"

	zeusctx "gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/context"

	"zeus_app/hello/logic/user"
	hello "zeus_app/hello/proto/gomicro"
)

func (h *Hello) SayHello(ctx context.Context, req *hello.HelloRequest, rsp *hello.HelloReply) (err error) {
	logger := zeusctx.ExtractLogger(ctx)
	info, err := user.GetInfo(ctx, "001")
	if err != nil {
		logger.Error("user.GetInfo err:", err)
		return
	}
	rsp.Message = "hello world, " + info + "."
	return
}
