package handler

import (
	"context"

	zeusctx "github.com/mmqbaba/zeus/context"

	gomicro "github.com/mmqbaba/zeus-examples/hellodemo/proto/hellodemopb"
)

func (h *HelloDemo) Put(ctx context.Context, req *gomicro.HelloRequest, rsp *gomicro.HelloReply) (err error) {
	logger := zeusctx.ExtractLogger(ctx)
	logger.Debug("Put")

	return
}
