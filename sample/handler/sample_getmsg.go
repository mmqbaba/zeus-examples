package handler

import (
	"context"

	zeusctx "gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/context"

	gomicro "zeus-examples/sample/proto/samplepb"

	zeushttpclient "gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/httpclient"
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

	hc, err := zeushttpclient.GetClient(ctx, "test")
	if err != nil {
		logger.Debug(err)
		return
	}
	hcret, err := hc.Get(ctx, "/json/?lang=zh-CN", nil)
	if err != nil {
		logger.Debug(err)
		return
	}
	logger.Debug(string(hcret))

	return
}
