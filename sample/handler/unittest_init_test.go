package handler

import (
	"log"
	"os"
	"testing"
	"time"

	"gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/plugin/container"
	"gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/service"

	"zeus-examples/sample/global"
)

func TestMain(m *testing.M) {
	log.SetPrefix("[zeus_unittest] ")
	log.SetFlags(3)

	opt := service.Options{
		LogLevel:      "debug",
		ConfEntryPath: "../conf/zeus.json",
	}
	s := service.NewService(opt, container.GetContainer(), global.ServiceOpts...)
	if err := s.Init(); err != nil {
		log.Fatalf("Service s.Init exited with error: %s\n", err)
	}
	time.Sleep(1 * time.Second)
	code := m.Run()
	os.Exit(code)
}
