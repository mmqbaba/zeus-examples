package handler

import (
	"context"

	zeusctx "github.com/mmqbaba/zeus/context"

	"zeus-examples/sampleservice/logic/user"
	"zeus-examples/sampleservice/proto/hello"
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
