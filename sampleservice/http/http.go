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
	g.NoRoute(zeusmwhttp.NotFound(ng))
	g.Use(zeusmwhttp.Access(ng))

	prefixGroup := g.Group(pathPrefix)

	prefixGroup.GET("/ping", func(c *gin.Context) {
		zeusmwhttp.ExtractLogger(c).Debug("ping")
		zeusmwhttp.SuccessResponse(c, gin.H{"message": "hello, zeus enginego."})
		// zeusmwhttp.ErrorResponse(c, nil)
	})

	// register routes for hellohandler
	addMWForRouteHelloHdlrPingPong := func(routes map[zeusmwhttp.RouteLink]*zeusmwhttp.Route) {
		Route_HelloHdlr_PingPong.AddMW(routes, func(c *gin.Context) {
			zeusmwhttp.ExtractLogger(c).Debug("begin ======", Route_HelloHdlr_PingPong)
			c.Next()
			zeusmwhttp.ExtractLogger(c).Debug("end ======", Route_HelloHdlr_PingPong)
		})
	}
	subGroup := map[string]*gin.RouterGroup{}
	registerRoutesForHelloHandler(prefixGroup, subGroup, addMWForRouteHelloHdlrPingPong)

	return g, nil
}
