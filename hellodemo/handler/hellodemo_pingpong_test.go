package handler

import (
	"context"
	"testing"

	gomicro "github.com/mmqbaba/zeus-examples/hellodemo/proto/hellodemopb"
)

func TestHelloDemo_PingPong(t *testing.T) {
	type args struct {
		ctx context.Context
		req *gomicro.PingRequest
		rsp *gomicro.PongReply
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
			if err := h.PingPong(tt.args.ctx, tt.args.req, tt.args.rsp); (err != nil) != tt.wantErr {
				t.Errorf("HelloDemo.PingPong() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
