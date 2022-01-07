package handler

import (
	"context"

	zeusctx "github.com/mmqbaba/zeus/context"

	gomicro "github.com/mmqbaba/zeus-examples/sample/proto/samplepb"
)

func (h *Sample) SendMsg(ctx context.Context, req *gomicro.GetMsgReq, rsp *gomicro.GetMsgResp) (err error) {
	logger := zeusctx.ExtractLogger(ctx)
	logger.Debug("SendMsg")

	return
}
