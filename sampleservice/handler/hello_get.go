package handler

import (
	"context"

	hello "zeus-examples/sampleservice/proto/gomicro"
)

func (h *Hello) Get(ctx context.Context, req *hello.HelloRequest, rsp *hello.HelloReply) (err error) {
	return
}
