package handler

import (
	"context"
	"reflect"
	"testing"

	gomicro "zeus-examples/sample/proto/samplepb"

	any "github.com/golang/protobuf/ptypes/any"
	structpb "github.com/golang/protobuf/ptypes/struct"
)

func TestSample_PingPong(t *testing.T) {
	type args struct {
		ctx context.Context
		req *gomicro.PingRequest
		rsp *gomicro.PongReply
	}
	tests := []struct {
		name    string
		h       *Sample
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Sample{}
			if err := h.PingPong(tt.args.ctx, tt.args.req, tt.args.rsp); (err != nil) != tt.wantErr {
				t.Errorf("Sample.PingPong() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_toAny(t *testing.T) {
	tests := []struct {
		name    string
		want    *any.Any
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := toAny()
			if (err != nil) != tt.wantErr {
				t.Errorf("toAny() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toAny() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toStruct(t *testing.T) {
	tests := []struct {
		name string
		want *structpb.Struct
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toStruct(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toStruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
