package http

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/engine"
	// zeuserr "gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/errors"
	zeusmwhttp "gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/middleware/http"
	// "zeus-examples/errdef"
)

func init() {
	// zeusmwhttp.SuccessResponse = customSsuccessResponse // 可初始化设置为自定义
	// zeusmwhttp.ErrorResponse = customErrorResponse

	zeusmwhttp.SetDefaultValidator(nil)
}

func serveHTTPHandler(ctx context.Context, pathPrefix string, ng engine.Engine) (http.Handler, error) {
	log.Println("serveHTTPHandler pathPrefix:", pathPrefix)
	g := gin.New()

	// TODO: 预留扩展
	// 这里可根据实际需求添加全局handlerfunc
	g.NoRoute(zeusmwhttp.NotFound(ng))
	g.Use(zeusmwhttp.Access(ng))
	g.Use(zeusmwhttp.Recovery())

	prefixGroup := g.Group(pathPrefix)

	prefixGroup.Use(zeusmwhttp.WrapHandlerCtx(func(c *gin.Context, ctx context.Context) context.Context {
		return context.WithValue(ctx, "a", "a1")
	}))
	prefixGroup.Use(zeusmwhttp.WrapHandlerCtx(func(c *gin.Context, ctx context.Context) context.Context {
		return context.WithValue(ctx, "bb", "bb1")
	}))
	prefixGroup.Use(func(c *gin.Context) {
		zeusmwhttp.AddWrapHandlerCtxFn(c, func(c *gin.Context, zctx context.Context) context.Context {
			zeusmwhttp.ExtractLogger(c).Debug("test AddWrapHandlerCtxFn")
			ctx := context.WithValue(zctx, "add-a", "add-a")
			ctx = context.WithValue(ctx, "add-b", "add-b")
			return ctx
		})
		c.Next()
	})

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
		Route_SampleHdlr_PingPong.AddMW(routes, zeusmwhttp.TagRawRsp(false))
		Route_SampleHdlr_SayHello.AddMW(routes, zeusmwhttp.UseGinBindValidateForPB(false))
		Route_SampleHdlr_SayHello.AddMW(routes, zeusmwhttp.DisablePBValidate(false))
		Route_SampleHdlr_SayHello.AddMW(routes, zeusmwhttp.WrapHandlerCtx(func(c *gin.Context, handlerCtx context.Context) context.Context {
			ctx := context.WithValue(handlerCtx, "wrapctx", c.Request.URL.String())
			ctx = context.WithValue(ctx, "ginctx", c)
			return ctx
		}))
		Route_SampleHdlr_SendMsg.AddMW(routes, zeusmwhttp.UseGinBindValidateForPB(true))
		Route_SampleHdlr_GetMsg.AddMW(routes, zeusmwhttp.UseGinBindValidateForPB(true))
		Route_SampleHdlr_DelMsg.AddMW(routes, zeusmwhttp.WrapHandlerCtx(func(c *gin.Context, handlerCtx context.Context) context.Context {
			ctx := context.WithValue(handlerCtx, "wrapctx", c.Request.URL.String())
			ctx = context.WithValue(ctx, "ginctx", c)
			return ctx
		}))
		// Route_SampleHdlr_SayHello.AddMW(routes, zeusmwhttp.SetReWriteErrFn(func(c *gin.Context, err error) {
		// 	if err != nil {
		// 		e, ok := err.(*zeuserr.Error)
		// 		if ok && e != nil {
		// 			if e.ErrCode == zeuserr.ECodeInvalidParams {
		// 				c.JSON(http.StatusBadRequest, gin.H{"ecode": errdef.ECodeSampleInvalidParams, "msg": e.ErrMsg})
		// 				return
		// 			}
		// 		}
		// 	}
		// 	c.JSON(http.StatusOK, gin.H{"ecode": zeuserr.ECodeSuccessed, "msg": err.Error()})
		// }))
		// Route_SampleHdlr_SayHello.AddMW(routes, zeusmwhttp.SetReWriteResponseFn(func(c *gin.Context, rsp interface{}) {
		// 	c.JSON(http.StatusOK, gin.H{"ecode": zeuserr.ECodeSuccessed, "data": rsp})
		// }))

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
