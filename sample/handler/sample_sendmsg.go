package handler

import (
	"context"

	zeusctx "gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/context"

	gomicro "zeus-examples/sample/proto/samplepb"
)

func (h *Sample) SendMsg(ctx context.Context, req *gomicro.GetMsgReq, rsp *gomicro.GetMsgResp) (err error) {
	logger := zeusctx.ExtractLogger(ctx)
	logger.Debug("SendMsg")

	return
}
