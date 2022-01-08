package handler

import (
	"context"
	"testing"

	gomicro "zeus-examples/hellodemo/proto/hellodemopb"
)

func TestHelloDemo_Delete(t *testing.T) {
	type args struct {
		ctx context.Context
		req *gomicro.HelloRequest
		rsp *gomicro.HelloReply
	}
	tests := []struct {
		name    string
		h       *HelloDemo
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HelloDemo{}
			if err := h.Delete(tt.args.ctx, tt.args.req, tt.args.rsp); (err != nil) != tt.wantErr {
				t.Errorf("HelloDemo.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
