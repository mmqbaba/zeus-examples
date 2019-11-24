package handler

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
	pb "zeus-examples/sample/proto/samplepb"
)

func TestMain(m *testing.M) {
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

func TestSayHello(t *testing.T) {
	cnt := global.GetNG().GetContainer()
	ctx := context.Background()
	logger := cnt.GetLogger()
	ctx = zeusctx.LoggerToContext(ctx, logger.WithField("tag", "sample_handler_sayhello_test"))
	rdc := cnt.GetRedisCli().GetCli()
	ctx = zeusctx.RedisToContext(ctx, rdc)
	// ctx = zeusctx.GMClientToContext(ctx, cnt.GetGoMicroClient())
	expected := &pb.Reply{Message: "hello world, mark."}
	sampleSrv := NewSample()
	req := &pb.Request{
		Name: "mark",
		Age:  21,
	}
	rsp := &pb.Reply{}
	err := sampleSrv.SayHello(ctx, req, rsp)
	if err != nil {
		t.Fatal(err)
	}
	if rsp.Message != expected.Message {
		t.Errorf("handler returned unexpected message: got %v want %v",
			rsp.Message, expected)
		return
	}
}
