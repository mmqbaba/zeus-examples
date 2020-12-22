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

	"zeus-examples/sample/global"
	"zeus-examples/sample/handler"
	gw "zeus-examples/sample/proto/samplepb"
)

const (
	Route_SampleHdlr_SayHello zeusmwhttp.RouteLink = "Route_SampleHdlr_SayHello"
	Route_SampleHdlr_PingPong zeusmwhttp.RouteLink = "Route_SampleHdlr_PingPong"
	Route_SampleHdlr_GetMsg zeusmwhttp.RouteLink = "Route_SampleHdlr_GetMsg"
	Route_SampleHdlr_DelMsg zeusmwhttp.RouteLink = "Route_SampleHdlr_DelMsg"
	Route_SampleHdlr_SendMsg zeusmwhttp.RouteLink = "Route_SampleHdlr_SendMsg"
	Route_SampleHdlr_TestStruct zeusmwhttp.RouteLink = "Route_SampleHdlr_TestStruct"

)

var sampleHdlr = handler.NewSample()
var sampleHdlrRoutes = map[zeusmwhttp.RouteLink]*zeusmwhttp.Route{
	Route_SampleHdlr_SayHello: &zeusmwhttp.Route{
		RLink:  Route_SampleHdlr_SayHello,
		Method: http.MethodPost,
		Path:   "/v1/hello",
		Handle: zeusmwhttp.GenerateGinHandle(sampleHdlr.SayHello),
	},
	Route_SampleHdlr_PingPong: &zeusmwhttp.Route{
		RLink:  Route_SampleHdlr_PingPong,
		Method: http.MethodPost,
		Path:   "/v1/pingpong",
		Handle: zeusmwhttp.GenerateGinHandle(sampleHdlr.PingPong),
	},
	Route_SampleHdlr_GetMsg: &zeusmwhttp.Route{
		RLink:  Route_SampleHdlr_GetMsg,
		Method: http.MethodGet,
		Path:   "/v1/getmsg",
		Handle: zeusmwhttp.GenerateGinHandle(sampleHdlr.GetMsg),
	},
	Route_SampleHdlr_DelMsg: &zeusmwhttp.Route{
		RLink:  Route_SampleHdlr_DelMsg,
		Method: http.MethodDelete,
		Path:   "/v1/delmsg",
		Handle: zeusmwhttp.GenerateGinHandle(sampleHdlr.DelMsg),
	},
	Route_SampleHdlr_SendMsg: &zeusmwhttp.Route{
		RLink:  Route_SampleHdlr_SendMsg,
		Method: http.MethodPost,
		Path:   "/v1/sendmsg",
		Handle: zeusmwhttp.GenerateGinHandle(sampleHdlr.SendMsg),
	},
	Route_SampleHdlr_TestStruct: &zeusmwhttp.Route{
		RLink:  Route_SampleHdlr_TestStruct,
		Method: http.MethodPost,
		Path:   "/v1/teststruct",
		Handle: zeusmwhttp.GenerateGinHandle(sampleHdlr.TestStruct),
	},

}

func init() {
	// grpc gateway
	global.ServiceOpts = append(global.ServiceOpts, service.WithHttpGWhandlerRegisterFnOption(gwHandlerRegister))
	// http handler
	global.ServiceOpts = append(global.ServiceOpts, service.WithHttpHandlerRegisterFnOption(getHandlerRegisterFn()))
	global.ServiceOpts = append(global.ServiceOpts, service.WithSwaggerJSONFileName("sample"))
}

func gwHandlerRegister(ctx context.Context, endpoint string, opts []grpc.DialOption) (m *gruntime.ServeMux, err error) {
	optsTmp := opts
	mux := gruntime.NewServeMux()
	if len(opts) == 0 {
		optsTmp = []grpc.DialOption{grpc.WithInsecure()}
	}
	if err = gw.RegisterSampleHandlerFromEndpoint(ctx, mux, endpoint, optsTmp); err != nil {
		log.Println("gw.RegisterSampleHandlerFromEndpoint err:", err)
		return
	}
	m = mux
	return
}

func getHandlerRegisterFn() service.HttpHandlerRegisterFn {
	return serveHTTPHandler
}

func registerRoutesForSampleHandler(groups map[string]*gin.RouterGroup, customFn ...zeusmwhttp.CustomRouteFn) {
	for _, f := range customFn {
		f(sampleHdlrRoutes)
	}
	for _, r := range sampleHdlrRoutes {
		zeusmwhttp.Method(groups, r)
	}
}
