package handler

import (
	"context"

	structpb "github.com/golang/protobuf/ptypes/struct"

	zeusctx "github.com/mmqbaba/zeus/context"
	zeusutilspb "github.com/mmqbaba/zeus/utils/protobuf"

	gomicro "zeus-examples/sample/proto/samplepb"
)

func (h *Sample) TestStruct(ctx context.Context, req *gomicro.GetMsgReq, rsp *structpb.Struct) (err error) {
	logger := zeusctx.ExtractLogger(ctx)
	logger.Debug("TestStruct")

	val := map[string]interface{}{
		"key": [2]interface{}{[2]interface{}{123, "abc"}, [2]interface{}{123, "bbc"}},
	}
	rsp.Fields = zeusutilspb.ToStruct(val).Fields
	return
}
