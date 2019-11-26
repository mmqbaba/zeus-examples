// Code generated by zeus-gen. DO NOT EDIT.
package rpc

import (
	"log"

	"github.com/micro/go-micro/server"

	"gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/service"

	"zeus-examples/sample/global"
	"zeus-examples/sample/handler"
	gomicro "zeus-examples/sample/proto/samplepb"
)

func init() {
	// gomicro
	global.ServiceOpts = append(global.ServiceOpts, service.WithGoMicrohandlerRegisterFnOption(gmSampleHandlerRegister))
	global.ServiceOpts = append(global.ServiceOpts, service.WithServiceNameOption("sample"))
}

func gmSampleHandlerRegister(s server.Server, opts ...server.HandlerOption) (err error) {
	if err = gomicro.RegisterSampleHandler(s, handler.NewSample(), opts...); err != nil {
		log.Println("gomicro.RegisterSampleHandler err:", err)
		return
	}
	return
}
