package handler

import (
	"context"

	zeusctx "gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/context"

	gomicro "zeus-examples/hellodemo/proto/hellodemopb"
)


func (h *HelloDemo) Delete(ctx context.Context, req *gomicro.HelloRequest, rsp *gomicro.HelloReply) (err error) {
	logger := zeusctx.ExtractLogger(ctx)
	logger.Debug("Delete")

	return
}
