package cache

import (
	"context"
    "testing"

    zeusctx "gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/context"

    "zeus-examples/sample/global"
)

func TestGetUser(t *testing.T) {
	cnt := global.GetNG().GetContainer()
	ctx := context.Background()
	logger := cnt.GetLogger()
	ctx = zeusctx.LoggerToContext(ctx, logger.WithField("tag", "sample_reource_cache_getuser_test"))
	rdc := cnt.GetRedisCli().GetCli()
    ctx = zeusctx.RedisToContext(ctx, rdc)

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
			gotInfo, err := GetUser(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotInfo != tt.wantInfo {
				t.Errorf("GetUser() = %v, want %v", gotInfo, tt.wantInfo)
			}
		})
	}
}

func TestSetUser(t *testing.T) {
	cnt := global.GetNG().GetContainer()
	ctx := context.Background()
	logger := cnt.GetLogger()
	ctx = zeusctx.LoggerToContext(ctx, logger.WithField("tag", "sample_reource_cache_setuser_test"))
	rdc := cnt.GetRedisCli().GetCli()
    ctx = zeusctx.RedisToContext(ctx, rdc)

	type args struct {
		ctx  context.Context
		id   string
		info string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
        // TODO: Add test cases.
        {
            "001",
            args{
                ctx,
                "001",
                "mark",
            },
            false,
        },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SetUser(tt.args.ctx, tt.args.id, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("SetUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
