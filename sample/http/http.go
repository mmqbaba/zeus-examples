package http

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/engine"
	zeuserr "gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/errors"
	zeusmwhttp "gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/middleware/http"

	"zeus-examples/errdef"
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
	sampleGroup := g.Group(pathPrefix, func(c *gin.Context) {
		zeusmwhttp.ExtractLogger(c).Debug("sample group")
		c.Next()
	})
	groups := map[string]*gin.RouterGroup{
		"default": prefixGroup,
		"sample":  sampleGroup,
	}
	////

	// TODO: 预留扩展
	// 这里可根据实际需求，为每条路由添加handlerfunc和设置路由组
	////
	customRouteSampleHdlr := zeusmwhttp.CustomRouteFn(func(routes map[zeusmwhttp.RouteLink]*zeusmwhttp.Route) {
		Route_SampleHdlr_PingPong.AddMW(routes, zeusmwhttp.TagRawRsp(true))
		Route_SampleHdlr_SayHello.AddMW(routes, zeusmwhttp.SetReWriteErrFn(func(c *gin.Context, err error) {
			if err != nil {
				e, ok := err.(*zeuserr.Error)
				if ok && e != nil {
					if e.ErrCode == zeuserr.ECodeInvalidParams {
						// errTmp := zeuserr.New(errdef.ECodeSampleInvalidParams, e.ErrMsg, e.Cause)
						// errTmp.ServiceID = e.ServiceID
						// errTmp.TracerID = e.TracerID
						// e = errTmp
						// e.Write(c.Writer)
						c.JSON(http.StatusBadRequest, gin.H{"ecode": errdef.ECodeSampleInvalidParams, "msg": e.ErrMsg})
						return
					}
				}
			}
			c.JSON(http.StatusOK, gin.H{"ecode": 0, "msg": "ok"})
		}))

		//Route_SampleHdlr_Demo.AddMW(routes, func(c *gin.Context) {
		//	zeusmwhttp.ExtractLogger(c).Debug("customRouteSampleHdlr: ", Route_SampleHdlr_PingPong)
		//	c.Next()
		//})
		//Route_SampleHdlr_Demo.SetGroup(routes, "sample")
	})
	////

	// register routes for Samplehandler
	registerRoutesForSampleHandler(groups, customRouteSampleHdlr)
	return g, nil
}
