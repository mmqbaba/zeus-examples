package handler

import (
	"context"

	zeusctx "github.com/mmqbaba/zeus/context"

	gomicro "zeus-examples/sample/proto/samplepb"
)

func (h *Sample) Upload(ctx context.Context, stream gomicro.Sample_UploadStream) (err error) {
	logger := zeusctx.ExtractLogger(ctx)
	logger.Debug("Upload")

	return
}
