package handler

import (
	"context"

	zeusctx "github.com/mmqbaba/zeus/context"

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

	hc, err := zeusctx.ExtractHttpClient(ctx)
	if err != nil {
		logger.Debug(err)
		return
	}
	hcc, err := hc.GetHttpClient("test")
	if err != nil {
		logger.Debug(err)
		return
	}
	hccret, err := hcc.Get(ctx, "/json/?lang=zh-CN", nil)
	if err != nil {
		logger.Debug(err)
		return
	}
	logger.Debug(string(hccret))

	return
}
