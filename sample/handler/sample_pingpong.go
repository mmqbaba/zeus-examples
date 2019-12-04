package handler

import (
	"context"

	proto "github.com/golang/protobuf/proto"
	// any "github.com/golang/protobuf/ptypes/any"
	"github.com/golang/protobuf/ptypes"

	zeusctx "gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/context"
	zeuserr "gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/errors"

	"zeus-examples/errdef"
	gomicro "zeus-examples/sample/proto/samplepb"
)

func (h *Sample) PingPong(ctx context.Context, req *gomicro.PingRequest, rsp *gomicro.PongReply) (err error) {
	logger := zeusctx.ExtractLogger(ctx)
	logger.Debug("PingPong")

	// 公共错误码
	zeuserr.ECodeRedisErr.ParseErr("")
	// 本项目错误码
	errdef.ECodeSampleInvalidParams.ParseErr("")

	// pb := &gomicro.Request{
	// 	MetaData: map[string]string{"email": "email", "home_addr": "home_addr"},
	// }
	// data := new(any.Any)
	// data.TypeUrl = "gitlab.digitalgd.com.cn/zeus-examples/" + proto.MessageName(pb)
	// data.Value, _ = proto.Marshal(pb)
	// rsp.Data = data

	var pb proto.Message
	pb = &gomicro.Request{
		MetaData: map[string]string{"email": "email", "home_addr": "home_addr"},
	}
	rsp.Data, err = ptypes.MarshalAny(pb)
	if err != nil {
		logger.Error(err)
		return
	}

	rsp.Pong = "pong"
	return
}
