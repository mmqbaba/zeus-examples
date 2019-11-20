package handler

import (
	"context"

	zeusctx "gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/context"

	gomicro "zeus-examples/hellodemo/proto/hellodemopb"
)


func (h *HelloDemo) Upload(ctx context.Context, stream gomicro.HelloDemo_UploadStream) (err error) {
	logger := zeusctx.ExtractLogger(ctx)
	logger.Debug("Upload")

	return
}
