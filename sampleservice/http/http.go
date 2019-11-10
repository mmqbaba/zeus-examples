package http

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func serveHttpHandler(ctx context.Context, pathPrefix string) (http.Handler, error) {
	log.Println("serveHttpHandler pathPrefix:", pathPrefix)
	g := gin.New()
	g.NoRoute(notFound)
	g.Use(access)

	prefixGroup := g.Group(pathPrefix)

	apiGroup := prefixGroup.Group("api")
	apiGroup.GET("/user", getUser)

	return g, nil
}

func notFound(c *gin.Context) {
	log.Printf("url not found url: %s\n", c.Request.URL)
	c.JSON(http.StatusNotFound, "not found")
}

func access(ctx *gin.Context) {
	log.Println("[gin] access start", ctx.Request.URL.Path)
	ctx.Next()
	log.Println("[gin] access end", ctx.Request.URL.Path)
}

func getUser(ctx *gin.Context) {
	log.Println("[gin] user")
	ctx.JSON(http.StatusOK, gin.H{
		"errcode": 0,
		"errmsg":  "ok",
		"data":    "hello, zeus.",
	})
}
