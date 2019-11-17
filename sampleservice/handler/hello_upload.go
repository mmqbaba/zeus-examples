package handler

import (
	"bytes"
	"context"

	zeusctx "gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/context"

	"zeus-examples/sampleservice/proto/hello"
)

func (h *Hello) Upload(ctx context.Context, stream hello.Hello_UploadStream) error {
	logger := zeusctx.ExtractLogger(ctx)
	var err error
	var req *hello.UploadReq
	buf := new(bytes.Buffer)
	rsp := new(hello.UploadResp)

	req, err = stream.Recv()
	if err != nil {
		logger.Error(err)
		return err
	}
	if req != nil {
		logger.Infof("file_name: %s\n", req.FileName)
		buf.Write(req.Content)
	}

	if err = stream.SendMsg(rsp); err != nil {
		logger.Error(err)
		return err
	}

	return nil
}
