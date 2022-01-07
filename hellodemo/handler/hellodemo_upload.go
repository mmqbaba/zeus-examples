package handler

import (
	"context"

	zeusctx "github.com/mmqbaba/zeus/context"

	gomicro "github.com/mmqbaba/zeus-examples/hellodemo/proto/hellodemopb"
)

func (h *HelloDemo) Upload(ctx context.Context, stream gomicro.HelloDemo_UploadStream) (err error) {
	logger := zeusctx.ExtractLogger(ctx)
	logger.Debug("Upload")

	return
}
