package cache

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	zeusctx "gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/context"
	"gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/plugin/container"
	"gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/service"

	"zeus-examples/sample/global"
)

func TestMain(m *testing.M) {
	opt := service.Options{
		LogLevel:      "debug",
		ConfEntryPath: "../../conf/zeus.json",
	}
	s := service.NewService(opt, container.GetContainer(), global.ServiceOpts...)
	if err := s.Init(); err != nil {
		log.Fatalf("Service s.Init exited with error: %s\n", err)
	}
	time.Sleep(1 * time.Second)
	code := m.Run()
	os.Exit(code)
}

func TestGetUser(t *testing.T) {
	cnt := global.GetNG().GetContainer()
	ctx := context.Background()
	logger := cnt.GetLogger()
	ctx = zeusctx.LoggerToContext(ctx, logger.WithField("tag", "sample_reource_cache_getuser_test"))
	rdc := cnt.GetRedisCli().GetCli()
	ctx = zeusctx.RedisToContext(ctx, rdc)
	id := "001"
	expected := id + "@" + "mark"
	info, err := GetUser(ctx, "001")
	if err != nil {
		t.Fatal(err)
	}
	if id+"@"+info != expected {
		t.Errorf("handler returned unexpected info: got %v want %v",
			id+"@"+info, expected)
		return
	}
}

func TestSetUser(t *testing.T) {
	cnt := global.GetNG().GetContainer()
	ctx := context.Background()
	logger := cnt.GetLogger()
	ctx = zeusctx.LoggerToContext(ctx, logger.WithField("tag", "sample_reource_cache_setuser_test"))
	rdc := cnt.GetRedisCli().GetCli()
	ctx = zeusctx.RedisToContext(ctx, rdc)
	id := "001"
	info := "mark"
	rdc.Del(id)
	err := SetUser(ctx, id, info)
	if err != nil {
		t.Fatal(err)
	}
}
