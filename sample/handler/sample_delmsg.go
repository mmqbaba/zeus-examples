package handler

import (
	"context"

	zeusctx "gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/context"

	gomicro "zeus-examples/sample/proto/samplepb"
)

func (h *Sample) DelMsg(ctx context.Context, req *gomicro.GetMsgReq, rsp *gomicro.GetMsgResp) (err error) {
	logger := zeusctx.ExtractLogger(ctx)
	logger.Debug("DelMsg")

	logger.Debug("========================", ctx.Value("a"), ctx.Value("bb"), ctx.Value("wrapctx"), ctx.Value("add-a"), ctx.Value("add-b"))

	return
}
