// Code generated by zeus-gen. DO NOT EDIT.

package rpc

import (
	"log"

	"github.com/micro/go-micro/server"

	"github.com/mmqbaba/zeus/service"

	"github.com/mmqbaba/zeus-examples/sampleservice/global"
	"github.com/mmqbaba/zeus-examples/sampleservice/handler"
	"github.com/mmqbaba/zeus-examples/sampleservice/proto/hello"
)

func init() {
	// gomicro
	global.ServiceOpts = append(global.ServiceOpts, service.WithGoMicrohandlerRegisterFnOption(gmHelloHandlerRegister))
	global.ServiceOpts = append(global.ServiceOpts, service.WithServiceNameOption("hello"))
}

func gmHelloHandlerRegister(s server.Server, opts ...server.HandlerOption) (err error) {
	if err = hello.RegisterHelloHandler(s, &handler.Hello{}, opts...); err != nil {
		log.Println("hello.RegisterHelloHandler err:", err)
		return
	}
	return
}
