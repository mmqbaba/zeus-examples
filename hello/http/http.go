package http

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/engine"
	zeusmwhttp "gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/middleware/http"
)

func serveHttpHandler(ctx context.Context, pathPrefix string, ng engine.Engine) (http.Handler, error) {
	log.Println("serveHttpHandler pathPrefix:", pathPrefix)
	g := gin.New()
	g.NoRoute(zeusmwhttp.NotFound(ng))
	g.Use(zeusmwhttp.Access(ng))

	prefixGroup := g.Group(pathPrefix)

	apiGroup := prefixGroup.Group("api")
	apiGroup.GET("/echo", getEcho)

	return g, nil
}

func getEcho(c *gin.Context) {
	zeusmwhttp.ExtractLogger(c).Debug("echo")
	zeusmwhttp.SuccessResponse(c, gin.H{"message": "hello, zeus enginego."})
}

