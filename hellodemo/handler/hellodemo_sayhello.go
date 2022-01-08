package handler

import (
	"context"

	zeusctx "github.com/mmqbaba/zeus/context"

	gomicro "zeus-examples/hellodemo/proto/hellodemopb"
)

func (h *HelloDemo) SayHello(ctx context.Context, req *gomicro.HelloRequest, rsp *gomicro.HelloReply) (err error) {
	logger := zeusctx.ExtractLogger(ctx)
	logger.Debug("SayHello")

	return
}
