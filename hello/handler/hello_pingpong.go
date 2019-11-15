package handler

import (
	"context"

	gomicro "zeus_app/hello/proto/gomicro"
)


func (h *Hello) PingPong(ctx context.Context, req *gomicro.PingRequest, rsp *gomicro.PongReply) (err error) {

	return
}
