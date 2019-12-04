package handler

import (
	"context"
	"fmt"
	"reflect"

	proto "github.com/golang/protobuf/proto"
	// any "github.com/golang/protobuf/ptypes/any"
	"github.com/golang/protobuf/ptypes"
	st "github.com/golang/protobuf/ptypes/struct"

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

	type ss struct {
		AbcVal string `json:"abc_val,omitempty"`
	}
	val := map[string]interface{}{
		"req":   &ss{AbcVal: "abcval"},
		"email": "emailemail",
		"phone": 12345678901,
		"my_data": map[string]interface{}{
			"a":   "a",
			"b":   true,
			"req": req,
		},
	}

	rsp.MetaData = ToStruct(val)

	rsp.Pong = "pong"
	return
}

// ToStruct converts a map[string]interface{} to a ptypes.Struct
func ToStruct(v map[string]interface{}) *st.Struct {
	size := len(v)
	if size == 0 {
		return nil
	}
	fields := make(map[string]*st.Value, size)
	for k, v := range v {
		fields[k] = ToValue(v)
	}
	return &st.Struct{
		Fields: fields,
	}
}

// ToValue converts an interface{} to a ptypes.Value
func ToValue(v interface{}) *st.Value {
	switch v := v.(type) {
	case nil:
		return nil
	case bool:
		return &st.Value{
			Kind: &st.Value_BoolValue{
				BoolValue: v,
			},
		}
	case int:
		return &st.Value{
			Kind: &st.Value_NumberValue{
				NumberValue: float64(v),
			},
		}
	case int8:
		return &st.Value{
			Kind: &st.Value_NumberValue{
				NumberValue: float64(v),
			},
		}
	case int32:
		return &st.Value{
			Kind: &st.Value_NumberValue{
				NumberValue: float64(v),
			},
		}
	case int64:
		return &st.Value{
			Kind: &st.Value_NumberValue{
				NumberValue: float64(v),
			},
		}
	case uint:
		return &st.Value{
			Kind: &st.Value_NumberValue{
				NumberValue: float64(v),
			},
		}
	case uint8:
		return &st.Value{
			Kind: &st.Value_NumberValue{
				NumberValue: float64(v),
			},
		}
	case uint32:
		return &st.Value{
			Kind: &st.Value_NumberValue{
				NumberValue: float64(v),
			},
		}
	case uint64:
		return &st.Value{
			Kind: &st.Value_NumberValue{
				NumberValue: float64(v),
			},
		}
	case float32:
		return &st.Value{
			Kind: &st.Value_NumberValue{
				NumberValue: float64(v),
			},
		}
	case float64:
		return &st.Value{
			Kind: &st.Value_NumberValue{
				NumberValue: v,
			},
		}
	case string:
		return &st.Value{
			Kind: &st.Value_StringValue{
				StringValue: v,
			},
		}
	case error:
		return &st.Value{
			Kind: &st.Value_StringValue{
				StringValue: v.Error(),
			},
		}
	default:
		// Fallback to reflection for other types
		return toValue(reflect.ValueOf(v))
	}
}

func toValue(v reflect.Value) *st.Value {
	switch v.Kind() {
	case reflect.Bool:
		return &st.Value{
			Kind: &st.Value_BoolValue{
				BoolValue: v.Bool(),
			},
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return &st.Value{
			Kind: &st.Value_NumberValue{
				NumberValue: float64(v.Int()),
			},
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return &st.Value{
			Kind: &st.Value_NumberValue{
				NumberValue: float64(v.Uint()),
			},
		}
	case reflect.Float32, reflect.Float64:
		return &st.Value{
			Kind: &st.Value_NumberValue{
				NumberValue: v.Float(),
			},
		}
	case reflect.Ptr:
		if v.IsNil() {
			return nil
		}
		return toValue(reflect.Indirect(v))
	case reflect.Array, reflect.Slice:
		size := v.Len()
		if size == 0 {
			return nil
		}
		values := make([]*st.Value, size)
		for i := 0; i < size; i++ {
			values[i] = toValue(v.Index(i))
		}
		return &st.Value{
			Kind: &st.Value_ListValue{
				ListValue: &st.ListValue{
					Values: values,
				},
			},
		}
	case reflect.Struct:
		t := v.Type()
		size := v.NumField()
		if size == 0 {
			return nil
		}
		fields := make(map[string]*st.Value, size)
		for i := 0; i < size; i++ {
			name := t.Field(i).Name
			// Better way?
			if len(name) > 0 && 'A' <= name[0] && name[0] <= 'Z' {
				fields[name] = toValue(v.Field(i))
			}
		}
		if len(fields) == 0 {
			return nil
		}
		return &st.Value{
			Kind: &st.Value_StructValue{
				StructValue: &st.Struct{
					Fields: fields,
				},
			},
		}
	case reflect.Map:
		keys := v.MapKeys()
		if len(keys) == 0 {
			return nil
		}
		fields := make(map[string]*st.Value, len(keys))
		for _, k := range keys {
			if k.Kind() == reflect.String {
				fields[k.String()] = toValue(v.MapIndex(k))
			}
		}
		if len(fields) == 0 {
			return nil
		}
		return &st.Value{
			Kind: &st.Value_StructValue{
				StructValue: &st.Struct{
					Fields: fields,
				},
			},
		}
	default:
		// Last resort
		return &st.Value{
			Kind: &st.Value_StringValue{
				StringValue: fmt.Sprint(v),
			},
		}
	}
}

// DecodeToMap converts a pb.Struct to a map from strings to Go types.
// DecodeToMap panics if s is invalid.
func DecodeToMap(s *st.Struct) map[string]interface{} {
	if s == nil {
		return nil
	}
	m := map[string]interface{}{}
	for k, v := range s.Fields {
		m[k] = decodeValue(v)
	}
	return m
}

func decodeValue(v *st.Value) interface{} {
	switch k := v.Kind.(type) {
	case *st.Value_NullValue:
		return nil
	case *st.Value_NumberValue:
		return k.NumberValue
	case *st.Value_StringValue:
		return k.StringValue
	case *st.Value_BoolValue:
		return k.BoolValue
	case *st.Value_StructValue:
		return DecodeToMap(k.StructValue)
	case *st.Value_ListValue:
		s := make([]interface{}, len(k.ListValue.Values))
		for i, e := range k.ListValue.Values {
			s[i] = decodeValue(e)
		}
		return s
	default:
		panic("protostruct: unknown kind")
	}
}
