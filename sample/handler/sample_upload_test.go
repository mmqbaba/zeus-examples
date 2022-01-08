package handler

import (
	"context"
	"testing"

	gomicro "zeus-examples/sample/proto/samplepb"
)

func TestSample_Upload(t *testing.T) {
	type args struct {
		ctx    context.Context
		stream gomicro.Sample_UploadStream
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
			if err := h.Upload(tt.args.ctx, tt.args.stream); (err != nil) != tt.wantErr {
				t.Errorf("Sample.Upload() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
