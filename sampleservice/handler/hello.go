package handler

import (
	"bytes"
	"context"

	zeusctx "gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/context"

	"zeus-examples/sampleservice/logic/user"
	hello "zeus-examples/sampleservice/proto/gomicro"
)

type HelloHDL struct{}

func (h *HelloHDL) SayHello(ctx context.Context, req *hello.HelloRequest, rsp *hello.HelloReply) (err error) {
	logger := zeusctx.ExtractLogger(ctx)
	info, err := user.GetInfo(ctx, "001")
	if err != nil {
		logger.Error("user.GetInfo err:", err)
		return
	}
	rsp.Message = "hello world, " + info + "."
	return
}

func (h *HelloHDL) PingPong(ctx context.Context, req *hello.PingRequest, rsp *hello.PongReply) (err error) {
	logger := zeusctx.ExtractLogger(ctx)
	logger.Debug("pingpong")
	rsp.Pong = "ping pong"
	return
}

func (h *HelloHDL) Upload(ctx context.Context, stream hello.Hello_UploadStream) error {
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
