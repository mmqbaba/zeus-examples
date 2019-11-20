package http

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/engine"
	zeusmwhttp "gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/middleware/http"
)

func init() {
	// zeusmwhttp.SuccessResponse = customSsuccessResponse // 可初始化设置为自定义
	// zeusmwhttp.ErrorResponse = customErrorResponse
}

func serveHTTPHandler(ctx context.Context, pathPrefix string, ng engine.Engine) (http.Handler, error) {
	log.Println("serveHTTPHandler pathPrefix:", pathPrefix)
	g := gin.New()

	// TODO: 预留扩展
	// 这里可根据实际需求添加全局handlerfunc
	g.NoRoute(zeusmwhttp.NotFound(ng))
	g.Use(zeusmwhttp.Access(ng))

	prefixGroup := g.Group(pathPrefix)
	prefixGroup.GET("/ping", func(c *gin.Context) {
		zeusmwhttp.ExtractLogger(c).Debug("ping")
		zeusmwhttp.SuccessResponse(c, gin.H{"message": "hello, zeus enginego."})
		// zeusmwhttp.ErrorResponse(c, nil)
	})

	// TODO: 预留扩展
	// 这里可根据实际需求，添加grouprouter
	////
	hellodemoGroup := g.Group(pathPrefix, func(c *gin.Context) {
		zeusmwhttp.ExtractLogger(c).Debug("hellodemo group")
		c.Next()
	})
	groups := map[string]*gin.RouterGroup{
		"default": prefixGroup,
		"hellodemo":   hellodemoGroup,
	}
	////

	// TODO: 预留扩展
	// 这里可根据实际需求，为每条路由添加handlerfunc和设置路由组
	////
	customRouteHelloDemoHdlr := zeusmwhttp.CustomRouteFn(func(routes map[zeusmwhttp.RouteLink]*zeusmwhttp.Route) {
		//Route_HelloDemoHdlr_Demo.AddMW(routes, func(c *gin.Context) {
		//	zeusmwhttp.ExtractLogger(c).Debug("customRouteHelloDemoHdlr: ", Route_HelloDemoHdlr_PingPong)
		//	c.Next()
		//})
		//Route_HelloDemoHdlr_Demo.SetGroup(routes, "hellodemo")
	})
	////

	// register routes for HelloDemohandler
	registerRoutesForHelloDemoHandler(groups, customRouteHelloDemoHdlr)
	return g, nil
}

