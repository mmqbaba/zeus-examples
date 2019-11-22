package handler

import (
	"context"

	zeusctx "gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/context"

	gomicro "zeus-examples/sample/proto/samplepb"
)


func (h *Sample) PingPong(ctx context.Context, req *gomicro.PingRequest, rsp *gomicro.PongReply) (err error) {
	logger := zeusctx.ExtractLogger(ctx)
	logger.Debug("PingPong")

	return
}
