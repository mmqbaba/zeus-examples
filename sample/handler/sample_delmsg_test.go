package handler

import (
	"context"
	"testing"

	gomicro "github.com/mmqbaba/zeus-examples/sample/proto/samplepb"
)

func TestSample_DelMsg(t *testing.T) {
	type args struct {
		ctx context.Context
		req *gomicro.GetMsgReq
		rsp *gomicro.GetMsgResp
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
			if err := h.DelMsg(tt.args.ctx, tt.args.req, tt.args.rsp); (err != nil) != tt.wantErr {
				t.Errorf("Sample.DelMsg() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
