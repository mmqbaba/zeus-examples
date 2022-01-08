package handler

import (
	"context"
	"testing"

	zeusctx "github.com/mmqbaba/zeus/context"

	"zeus-examples/sample/global"
	gomicro "zeus-examples/sample/proto/samplepb"
)

func TestSample_SayHello(t *testing.T) {
	cnt := global.GetNG().GetContainer()
	ctx := context.Background()
	logger := cnt.GetLogger()
	ctx = zeusctx.LoggerToContext(ctx, logger.WithField("tag", "sample_handler_sayhello_test"))
	rdc := cnt.GetRedisCli().GetCli()
	ctx = zeusctx.RedisToContext(ctx, rdc)
	type args struct {
		ctx context.Context
		req *gomicro.Request
		rsp *gomicro.Reply
	}
	tests := []struct {
		name    string
		h       *Sample
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"sayhello_name",
			NewSample(),
			args{
				ctx,
				&gomicro.Request{
					Name: "mark",
					Age:  21,
				},
				&gomicro.Reply{},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expected := &gomicro.Reply{Message: "hello world, " + tt.args.req.Name + "."}
			if err := tt.h.SayHello(tt.args.ctx, tt.args.req, tt.args.rsp); (err != nil) != tt.wantErr {
				t.Errorf("Sample.SayHello() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.args.rsp.Message != expected.Message {
				t.Errorf("Sample.SayHello() handler returned unexpected message: got %v want %v",
					tt.args.rsp.Message, expected.Message)
				return
			}
		})
	}
}
