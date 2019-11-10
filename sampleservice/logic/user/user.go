package user

import (
	"context"
	"log"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"

	hello "zeus-examples/sampleservice/proto/gomicro"
	"zeus-examples/sampleservice/resource/cache"
	"zeus-examples/sampleservice/resource/rpc"
)

func GetInfo(ctx context.Context, id string) (info string, err error) {
	info, err = cache.GetUser(ctx, id)
	if err != nil {
		return
	}
	cli, err := rpc.GetHelloService(ctx)
	if err != nil {
		log.Println("rpc.GetHelloService err:", err)
		return
	}
	rsp, err := cli.PingPong(ctx, &hello.PingRequest{Ping: "ping"}, client.WithCallWrapper(func(c client.CallFunc) client.CallFunc {
		return func(ctx context.Context, node *registry.Node, req client.Request, rsp interface{}, opts client.CallOptions) error {
			log.Println("wrapcall PingPong begin")
			err := c(ctx, node, req, rsp, opts)
			log.Println("wrapcall PingPong end")
			return err
		}
	}))
	if err != nil {
		log.Println("cli.PingPong err:", err)
		return
	}
	log.Println("cli.PingPong success. rsp.Pong:", rsp.Pong)
	return
}
