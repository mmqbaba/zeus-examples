package handler

import (
	"context"
	"testing"
	gomicro "zeus-examples/sample/proto/samplepb"

	structpb "github.com/golang/protobuf/ptypes/struct"
)

func TestStruct(t *testing.T) {
	type args struct {
		ctx context.Context
		req *gomicro.GetMsgReq
		rsp *structpb.Struct
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
			if err := h.TestStruct(tt.args.ctx, tt.args.req, tt.args.rsp); (err != nil) != tt.wantErr {
				t.Errorf("Sample.TestStruct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
