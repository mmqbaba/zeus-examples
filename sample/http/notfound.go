package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mmqbaba/zeus/engine"
	zeusmwhttp "github.com/mmqbaba/zeus/middleware/http"
)

func readHtml(filepath string) []byte {
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Errorf("读取html文件[%s]错误：", filepath, err)
	}
	return bytes
}

func NotFound(ng engine.Engine, websiteroot string, skipprefixpaths ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// log.Printf("url not found url: %s\n", c.Request.URL)
		zeusmwhttp.ExtractLogger(c).Debugf("url not found url: %s\n", c.Request.URL)
		path := c.Request.URL.Path
		if strings.HasPrefix(path, websiteroot) {
			for _, p := range skipprefixpaths {
				if strings.HasPrefix(path, p) {
					c.String(http.StatusNotFound, "not found")
					return
				}
			}
			// c.Redirect(302, "index.html")
			c.Header("Content-Type", "text/html; charset=utf-8")
			c.String(http.StatusOK, string(readHtml("./website/index.html")))
		} else {
			c.String(http.StatusNotFound, "not found")
		}
	}
}
