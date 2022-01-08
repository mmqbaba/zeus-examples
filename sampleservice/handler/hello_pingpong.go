package handler

import (
	"context"

	zeusctx "github.com/mmqbaba/zeus/context"

	"zeus-examples/sampleservice/proto/hello"
)

func (h *Hello) PingPong(ctx context.Context, req *hello.PingRequest, rsp *hello.PongReply) (err error) {
	logger := zeusctx.ExtractLogger(ctx)
	logger.Debug("pingpong")
	rsp.Pong = "ping pong"
	return
}
