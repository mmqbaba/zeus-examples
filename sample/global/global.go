package global
import (
    "gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/config"
    "gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/engine"
)

func loadConfig(conf *config.AppConf) {
    // 加载配置
    // TODO: do something here
}

func loadEngineSuccess(ng engine.Engine) {
    loadConfig(GetConfig())
    // 加载engine成功
    // TODO: do something here
}

func processChange(event interface{}) {
    loadConfig(GetConfig())
    // 配置变更
    // TODO: do something here
}

