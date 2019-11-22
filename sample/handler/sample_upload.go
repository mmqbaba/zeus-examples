package handler

import (
	"context"

	zeusctx "gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/context"

	gomicro "zeus-examples/sample/proto/samplepb"
)


func (h *Sample) Upload(ctx context.Context, stream gomicro.Sample_UploadStream) (err error) {
	logger := zeusctx.ExtractLogger(ctx)
	logger.Debug("Upload")

	return
}
