// Code generated by zeus-gen. DO NOT EDIT.
package global

import (
	"log"

	"github.com/mmqbaba/zeus/config"
	"github.com/mmqbaba/zeus/engine"
	"github.com/mmqbaba/zeus/service"
)

var ng engine.Engine
var ServiceOpts []service.Option

func init() {
	// load engine
	loadEngineFnOpt := service.WithLoadEngineFnOption(func(ng engine.Engine) {
		log.Println("WithLoadEngineFnOption: SetNG success.")
		SetNG(ng)
	})
	ServiceOpts = append(ServiceOpts, loadEngineFnOpt)
	// // server wrap
	// ServiceOpts = append(ServiceOpts, service.WithGoMicroServerWrapGenerateFnOption(gomicro.GenerateServerLogWrap))
}

// GetNG ...
func GetNG() engine.Engine {
	return ng
}

// SetNG ...
func SetNG(n engine.Engine) {
	ng = n
}

// GetConfig ...
func GetConfig() (conf *config.AppConf) {
	c, err := ng.GetConfiger()
	if err != nil {
		log.Println("global.GetConfig err:", err)
		return
	}
	conf = c.Get()
	return
}
