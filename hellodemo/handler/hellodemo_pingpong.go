package handler

import (
	"context"

	zeusctx "github.com/mmqbaba/zeus/context"

	gomicro "github.com/mmqbaba/zeus-examples/hellodemo/proto/hellodemopb"
)

func (h *HelloDemo) PingPong(ctx context.Context, req *gomicro.PingRequest, rsp *gomicro.PongReply) (err error) {
	logger := zeusctx.ExtractLogger(ctx)
	logger.Debug("PingPong")

	return
}
