package handler

import (
	"context"

	zeusctx "gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/context"

	gomicro "zeus-examples/sample/proto/samplepb"
)

func (h *Sample) GetMsg(ctx context.Context, req *gomicro.GetMsgReq, rsp *gomicro.GetMsgResp) (err error) {
	logger := zeusctx.ExtractLogger(ctx)
	logger.Debug("GetMsg")
	redisCli, err := zeusctx.ExtractRedis(ctx)
	if err != nil {
		logger.Debug(err)
		return
	}
	ret := redisCli.Get("001")
	err = ret.Err()
	if err != nil {
		logger.Debug(err)
		return
	}
	logger.Debug(ret.Val())

	return
}