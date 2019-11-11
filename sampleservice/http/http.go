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
	apiGroup.GET("/user", getUser)

	return g, nil
}

func getUser(c *gin.Context) {
	zeusmwhttp.ExtractLogger(c).Debug("getUser")
	c.JSON(http.StatusOK, gin.H{
		"errcode": 0,
		"errmsg":  "ok",
		"data":    "hello, zeus engingo.",
	})
}
