package user

import (
	"context"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"

	zeusctx "github.com/mmqbaba/zeus/context"

	sample "zeus-examples/sample/proto/samplepb"
	"zeus-examples/sample/resource/cache"
	rpc "zeus-examples/sample/resource/rpcclient"
	"zeus-examples/sampleservice/proto/hello"
)

func GetInfo(ctx context.Context, id string) (info string, err error) {
	info, err = cache.GetUser(ctx, id)
	if err != nil {
		return
	}
	cli, err := rpc.NewHelloService(ctx)
	if err != nil {
		zeusctx.ExtractLogger(ctx).Error("rpc.GetHelloService err:", err)
		return
	}
	rsp, err := cli.PingPong(ctx, &hello.PingRequest{Ping: "ping"}, client.WithCallWrapper(func(c client.CallFunc) client.CallFunc {
		return func(ctx context.Context, node *registry.Node, req client.Request, rsp interface{}, opts client.CallOptions) error {
			zeusctx.ExtractLogger(ctx).Debug("wrapcall PingPong begin")
			err := c(ctx, node, req, rsp, opts)
			zeusctx.ExtractLogger(ctx).Debug("wrapcall PingPong end")
			return err
		}
	}))
	if err != nil {
		zeusctx.ExtractLogger(ctx).Error("cli.PingPong err:", err)
		return
	}
	zeusctx.ExtractLogger(ctx).Debug("cli.PingPong success. rsp.Pong:", rsp.Pong)

	sampleSrv, _ := rpc.NewSampleService(ctx)
	sampleSrv.PingPong(ctx, &sample.PingRequest{Ping: "ping"})
	return
}

func GetInfoEx(ctx context.Context, id string) (info string, err error) {
	info, err = cache.GetUser(ctx, id)
	if err != nil {
		return
	}
	zeusctx.ExtractLogger(ctx).Debugf("info:%s@%s", id, info)
	return
}
