package http

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/engine"
	zeusmwhttp "gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/middleware/http"
)

// func init() {
// 	zeusmwhttp.SuccessResponse = customSsuccessResponse // 可初始化设置为自定义
// 	zeusmwhttp.ErrorResponse = customErrorResponse
// }

func serveHttpHandler(ctx context.Context, pathPrefix string, ng engine.Engine) (http.Handler, error) {
	log.Println("serveHttpHandler pathPrefix:", pathPrefix)
	g := gin.New()
	g.NoRoute(zeusmwhttp.NotFound(ng))
	g.Use(zeusmwhttp.Access(ng))

	prefixGroup := g.Group(pathPrefix)

	apiGroup := prefixGroup.Group("api")
	apiGroup.GET("/echo", echo)

	return g, nil
}

func echo(c *gin.Context) {
	zeusmwhttp.ExtractLogger(c).Debug("echo")
	zeusmwhttp.SuccessResponse(c, gin.H{"message": "hello, zeus enginego."})
	// zeusmwhttp.ErrorResponse(c, err)
}
