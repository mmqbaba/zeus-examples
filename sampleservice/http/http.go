package http

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/mmqbaba/zeus/engine"
	zeusmwhttp "github.com/mmqbaba/zeus/middleware/http"
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
	helloGroup := g.Group(pathPrefix, func(c *gin.Context) {
		zeusmwhttp.ExtractLogger(c).Debug("hello group")
		c.Next()
	})
	groups := map[string]*gin.RouterGroup{
		"default": prefixGroup,
		"hello":   helloGroup,
	}
	////

	// TODO: 预留扩展
	// 这里可根据实际需求，为每条路由添加handlerfunc和设置路由组
	////
	customRouteHelloHdlr := zeusmwhttp.CustomRouteFn(func(routes map[zeusmwhttp.RouteLink]*zeusmwhttp.Route) {
		Route_HelloHdlr_PingPong.AddMW(routes, func(c *gin.Context) {
			zeusmwhttp.ExtractLogger(c).Debug("customRouteHelloHdlr: ", Route_HelloHdlr_PingPong)
			c.Next()
		})
		Route_HelloHdlr_PingPong.SetGroup(routes, "hello")
	})
	////

	// register routes for hellohandler
	registerRoutesForHelloHandler(groups, customRouteHelloHdlr)
	return g, nil
}
