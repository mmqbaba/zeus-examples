// Code generated by zeus-gen. DO NOT EDIT.

package http

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	gruntime "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	zeusmwhttp "gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/middleware/http"
	"gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/service"

	"zeus-examples/sampleservice/global"
	"zeus-examples/sampleservice/handler"
	"zeus-examples/sampleservice/proto/hello"
)

const (
	Route_HelloHdlr_SayHello zeusmwhttp.RouteLink = "Route_HelloHdlr_SayHello"
	Route_HelloHdlr_PingPong zeusmwhttp.RouteLink = "Route_HelloHdlr_PingPong"
)

var helloHdlr = &handler.Hello{}
var helloHdlrRoutes = map[zeusmwhttp.RouteLink]*zeusmwhttp.Route{
	Route_HelloHdlr_SayHello: &zeusmwhttp.Route{
		RLink:  Route_HelloHdlr_SayHello,
		Method: http.MethodPost,
		Path:   "/v1/hello",
		Handle: zeusmwhttp.GenerateGinHandle(helloHdlr.SayHello),
	},
	Route_HelloHdlr_PingPong: &zeusmwhttp.Route{
		RLink:  Route_HelloHdlr_PingPong,
		Method: http.MethodPost,
		Path:   "/v1/pingpong",
		Handle: zeusmwhttp.GenerateGinHandle(helloHdlr.PingPong),
	},
}

func init() {
	// grpc gateway
	global.ServiceOpts = append(global.ServiceOpts, service.WithHttpGWhandlerRegisterFnOption(gwHandlerRegister))

	// http handler
	global.ServiceOpts = append(global.ServiceOpts, service.WithHttpHandlerRegisterFnOption(getHandlerRegisterFn()))
	global.ServiceOpts = append(global.ServiceOpts, service.WithSwaggerJSONFileName("hello"))
}

func gwHandlerRegister(ctx context.Context, endpoint string, opts []grpc.DialOption) (m *gruntime.ServeMux, err error) {
	optsTmp := opts
	mux := gruntime.NewServeMux()
	if len(opts) == 0 {
		optsTmp = []grpc.DialOption{grpc.WithInsecure()}
	}
	if err = hello.RegisterHelloHandlerFromEndpoint(ctx, mux, endpoint, optsTmp); err != nil {
		log.Println("hellogw.RegisterHelloHandlerFromEndpoint err:", err)
		return
	}
	m = mux
	return
}

func getHandlerRegisterFn() service.HttpHandlerRegisterFn {
	return serveHTTPHandler
}

func registerRoutesForHelloHandler(groups map[string]*gin.RouterGroup, customFn ...zeusmwhttp.CustomRouteFn) {
	for _, f := range customFn {
		f(helloHdlrRoutes)
	}
	for _, r := range helloHdlrRoutes {
		zeusmwhttp.Method(groups, r)
	}
}
