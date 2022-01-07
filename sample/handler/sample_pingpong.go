package handler

import (
	"context"

	proto "github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	any "github.com/golang/protobuf/ptypes/any"
	structpb "github.com/golang/protobuf/ptypes/struct"

	zeusctx "github.com/mmqbaba/zeus/context"
	zeuserr "github.com/mmqbaba/zeus/errors"
	zeusutilspb "github.com/mmqbaba/zeus/utils/protobuf"

	"zeus-examples/errdef"
	gomicro "zeus-examples/sample/proto/samplepb"
)

func (h *Sample) PingPong(ctx context.Context, req *gomicro.PingRequest, rsp *gomicro.PongReply) (err error) {
	logger := zeusctx.ExtractLogger(ctx)
	logger.Debug("PingPong")

	// "any_data": {
	// 	"@type": "type.googleapis.com/sample.Request",
	// 	"name": "mark",
	// 	"age": 10,
	// 	"meta_data": {
	// 		"email": "email",
	// 		"home_addr": "home_addr"
	// 	},
	// 	"sex_type": true
	// }
	var pbRequest proto.Message = &gomicro.Request{}
	// var pbRequest proto.Message
	ptypes.UnmarshalAny(req.AnyData, pbRequest)
	logger.Debug(pbRequest)
	// logger.Debug(pbRequest.Name)
	// logger.Debug(pbRequest.MetaData)

	// helloSrv, err := hellodemorpc.NewHelloDemoService(ctx)
	// if err != nil {
	// 	logger.Error(err)
	// 	return
	// }
	// _, err = helloSrv.SayHello(ctx, &hellodemopb.HelloRequest{Age: 21})
	// if err != nil {
	// 	logger.Error(err)
	// 	return
	// }

	// 公共错误码
	zeuserr.ECodeRedisErr.ParseErr("")
	// 本项目错误码
	errdef.ECodeSampleInvalidParams.ParseErr("")

	// to any
	rsp.AnyData, err = toAny()
	if err != nil {
		logger.Error(err)
		return
	}

	// to struct
	rsp.StCustomData = toStruct()

	rsp.StMetaData = zeusutilspb.ToStruct(zeusutilspb.DecodeToMap(req.StMetaData))

	rsp.Pong = "pong"
	return
}

func toAny() (*any.Any, error) {
	var pb proto.Message
	pb = &gomicro.Request{
		MetaData: map[string]string{"email": "email", "home_addr": "home_addr"},
	}
	return ptypes.MarshalAny(pb)

	// pb := &gomicro.Request{
	// 	MetaData: map[string]string{"email": "email", "home_addr": "home_addr"},
	// }
	// data := new(any.Any)
	// data.TypeUrl = "gitlab.digitalgd.com.cn/zeus-examples/" + proto.MessageName(pb)
	// data.Value, _ = proto.Marshal(pb)
	// rsp.Data = data
}

func toStruct() *structpb.Struct {
	type ss struct {
		AbcVal string                 `json:"abc_val,omitempty"`
		SsData *gomicro.PingRequest   `json:"ss_data"`
		SsOk   bool                   `json:"ss_ok"`
		MySs   *ss                    `json:"my_ss"`
		MyMeta map[string]interface{} `json:"my_meta"`
	}
	pingReq := &gomicro.PingRequest{
		Ping: "ping",
	}
	req := &gomicro.Request{
		Name:    "mark",
		Age:     21,
		SexType: true,
		MetaData: map[string]string{
			"t1": "t1",
		},
	}
	val := map[string]interface{}{
		"struct_pingrequest": pingReq,
		"struct_ss":          &ss{AbcVal: "abcval", SsData: pingReq, SsOk: true, MySs: &ss{AbcVal: "abcval", SsData: pingReq, SsOk: true, MySs: &ss{}}},
		"string":             "emailemail",
		"number":             12345678901,
		"map_string_interface": map[string]interface{}{
			"a":   "a",
			"b":   true,
			"req": req,
		},
		"string_list":         []string{"l1", "l2"},
		"struct_list_request": []*gomicro.Request{req, req},
		"struct_list_ss": []*ss{
			&ss{AbcVal: "abcval", SsData: pingReq, SsOk: true, MySs: &ss{AbcVal: "abcval", SsData: pingReq, SsOk: true, MySs: &ss{}}, MyMeta: map[string]interface{}{"m1": true, "m2": "m2", "struct": req}},
			&ss{AbcVal: "abcval", SsData: pingReq, SsOk: true, MySs: &ss{AbcVal: "abcval", SsData: pingReq, SsOk: true, MySs: &ss{}}},
		},
	}

	return zeusutilspb.ToStruct(val)

	// stval := &st.Struct{
	// 	Fields: map[string]*st.Value{
	// 		"email": &st.Value{Kind: &st.Value_StringValue{StringValue: "email"}},
	// 		"phone": &st.Value{Kind: &st.Value_NumberValue{NumberValue: 123456789}},
	// 		"mystruct": &st.Value{Kind: &st.Value_StructValue{
	// 			StructValue: &st.Struct{
	// 				Fields: map[string]*st.Value{
	// 					"email": &st.Value{Kind: &st.Value_StringValue{StringValue: "email"}},
	// 					"phone": &st.Value{Kind: &st.Value_NumberValue{NumberValue: 123456789}},
	// 				},
	// 			},
	// 		}},
	// 		// "list":   &st.Value{Kind: &st.Value_ListValue{}},
	// 	},
	// }
	// val := DecodeToMap(stval)
}
