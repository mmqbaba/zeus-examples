package handler

import (
	"context"
	"testing"

	gomicro "github.com/mmqbaba/zeus-examples/hellodemo/proto/hellodemopb"
)

func TestHelloDemo_Upload(t *testing.T) {
	type args struct {
		ctx    context.Context
		stream gomicro.HelloDemo_UploadStream
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
			if err := h.Upload(tt.args.ctx, tt.args.stream); (err != nil) != tt.wantErr {
				t.Errorf("HelloDemo.Upload() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
