package sample

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/errors"
	"gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/plugin/container"
	"gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/service"
	"gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/utils"

	_ "zeus-examples/sample/http"
	_ "zeus-examples/sample/rpc"

	"zeus-examples/sample/global"
	proto "zeus-examples/sample/proto/samplepb"
)

func init() {
	log.SetPrefix("[zeus] ")
	log.SetFlags(3)
}

func TestMain(m *testing.M) {
	opt := service.Options{
		LogLevel:      "error",
		ConfEntryPath: "./conf/zeus.json",
	}
	s := service.NewService(opt, container.GetContainer(), global.ServiceOpts...)
	if err := s.Init(); err != nil {
		log.Fatalf("Service s.Init exited with error: %s\n", err)
	}
	go func() {
		if err := s.RunServer(); err != nil {
			log.Fatalf("Service s.RunServer exited with error: %s\n", err)
		}
	}()
	time.Sleep(1 * time.Second)
	code := m.Run()
	os.Exit(code)
}

func TestHTTP_SayHello(t *testing.T) {
	reqVal := &proto.Request{
		Name: "mark",
		Age:  21,
	}
	jsonStr, _ := utils.Marshal(reqVal)

	req, err := http.NewRequest("POST", "/api/v1/hello", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := global.GetNG().GetContainer().GetHTTPHandler()
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
		return
	}
	expected := `{"message":"hello world, mark."}`
	ret := errors.ECodeSuccessed.ParseErr("")
	rsp := &proto.Reply{}
	ret.Data = rsp
	err = utils.Unmarshal(rr.Body.Bytes(), ret)
	if err != nil {
		t.Fatal(err)
	}
	if ret.ErrCode != errors.ECodeSuccessed {
		t.Errorf("handler returned wrong err code: got %v want %v",
			ret.ErrCode, errors.ECodeSuccessed)
		return
	}
	result, err := utils.Marshal(rsp)
	if err != nil {
		t.Fatal(err)
	}
	if string(result) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			string(result), expected)
		return
	}
}

func TestHTTP_PingPong(t *testing.T) {
	reqVal := &proto.PingRequest{
		Ping: "ping",
	}
	jsonStr, _ := utils.Marshal(reqVal)

	req, err := http.NewRequest("POST", "/api/v1/pingpong", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := global.GetNG().GetContainer().GetHTTPHandler()
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
		return
	}
	expected := `{"pong":"pong"}`
	ret := errors.ECodeSuccessed.ParseErr("")
	rsp := &proto.PongReply{}
	ret.Data = rsp
	err = utils.Unmarshal(rr.Body.Bytes(), ret)
	if err != nil {
		t.Fatal(err)
	}
	if ret.ErrCode != errors.ECodeSuccessed {
		t.Errorf("handler returned wrong err code: got %v want %v",
			ret.ErrCode, errors.ECodeSuccessed)
		return
	}
	result, err := utils.Marshal(rsp)
	if err != nil {
		t.Fatal(err)
	}
	if string(result) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			string(result), expected)
		return
	}
}

func TestGoMicro_SayHello(t *testing.T) {
	reqVal := &proto.Request{
		Name: "mark",
		Age:  21,
	}
	jsonStr, _ := utils.Marshal(reqVal)

	req, err := http.NewRequest("POST", "/v1/hello", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := global.GetNG().GetContainer().GetHTTPHandler()
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
		return
	}
	expected := `{"message":"hello world, mark."}`
	ret := errors.ECodeSuccessed.ParseErr("")
	rsp := &proto.Reply{}
	ret.Data = rsp
	err = utils.Unmarshal(rr.Body.Bytes(), ret)
	if err != nil {
		t.Fatal(err)
	}
	if ret.ErrCode != errors.ECodeSuccessed {
		t.Errorf("handler returned wrong err code: got %v want %v",
			ret.ErrCode, errors.ECodeSuccessed)
		return
	}
	result, err := utils.Marshal(rsp)
	if err != nil {
		t.Fatal(err)
	}
	if string(result) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			string(result), expected)
		return
	}
}

func TestGoMicro_PingPong(t *testing.T) {
	reqVal := &proto.PingRequest{
		Ping: "ping",
	}
	jsonStr, _ := utils.Marshal(reqVal)

	req, err := http.NewRequest("POST", "/v1/pingpong", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := global.GetNG().GetContainer().GetHTTPHandler()
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
		return
	}
	expected := `{"pong":"pong"}`
	ret := errors.ECodeSuccessed.ParseErr("")
	rsp := &proto.PongReply{}
	ret.Data = rsp
	err = utils.Unmarshal(rr.Body.Bytes(), ret)
	if err != nil {
		t.Fatal(err)
	}
	if ret.ErrCode != errors.ECodeSuccessed {
		t.Errorf("handler returned wrong err code: got %v want %v",
			ret.ErrCode, errors.ECodeSuccessed)
		return
	}
	result, err := utils.Marshal(rsp)
	if err != nil {
		t.Fatal(err)
	}
	if string(result) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			string(result), expected)
		return
	}
}
