package user

import (
	"context"
	"testing"

	zeusctx "gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/context"

	"zeus-examples/sample/global"
)

func TestGetInfoEx(t *testing.T) {
	cnt := global.GetNG().GetContainer()
	ctx := context.Background()
	logger := cnt.GetLogger()
	ctx = zeusctx.LoggerToContext(ctx, logger.WithField("tag", "sample_logic_user_getinfoex_test"))
	rdc := cnt.GetRedisCli().GetCli()
	ctx = zeusctx.RedisToContext(ctx, rdc)
	ctx = zeusctx.GMClientToContext(ctx, cnt.GetGoMicroClient())
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name     string
		args     args
		wantInfo string
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			"001",
			args{
				ctx,
				"001",
			},
			"mark",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotInfo, err := GetInfoEx(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetInfoEx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotInfo != tt.wantInfo {
				t.Errorf("GetInfoEx() = %v, want %v", gotInfo, tt.wantInfo)
			}
		})
	}
}
