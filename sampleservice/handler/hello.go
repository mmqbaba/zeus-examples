package handler

import (
	"bytes"
	"context"
	"log"

	"zeus-examples/sampleservice/logic/user"
	hello "zeus-examples/sampleservice/proto/gomicro"
)

type HelloHDL struct{}

func (h *HelloHDL) SayHello(ctx context.Context, req *hello.HelloRequest, rsp *hello.HelloReply) (err error) {
	err = req.Validate()
	if err != nil {
		return
	}
	info, err := user.GetInfo(ctx, "001")
	if err != nil {
		log.Println("user.GetInfo err:", err)
		return
	}
	rsp.Message = "hello world, " + info + "."
	return
}

func (h *HelloHDL) PingPong(ctx context.Context, req *hello.PingRequest, rsp *hello.PongReply) (err error) {
	err = req.Validate()
	if err != nil {
		return
	}
	log.Println("pingpong")
	rsp.Pong = "ping pong"
	return
}

func (h *HelloHDL) Upload(ctx context.Context, stream hello.Hello_UploadStream) error {
	var err error
	var req *hello.UploadReq
	buf := new(bytes.Buffer)
	rsp := new(hello.UploadResp)

	req, err = stream.Recv()
	if err != nil {
		log.Println(err)
		return err
	}
	if req != nil {
		log.Printf("file_name: %s\n", req.FileName)
		buf.Write(req.Content)
	}

	if err = stream.SendMsg(rsp); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
