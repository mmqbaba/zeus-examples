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

func TestMain(m *testing.M) {
	log.SetPrefix("[zeus_unittest] ")
	log.SetFlags(3)

	opt := service.Options{
		LogLevel:      "debug",
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
	time.Sleep(1 * time.Second)
	os.Exit(code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	global.GetNG().GetContainer().GetHTTPHandler().ServeHTTP(rr, req)
	return rr
}

func checkResponseStatusCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response statuscode %d. Got %d\n", expected, actual)
	}
}

func checkErrCode(t *testing.T, expected errors.ErrorCode, ret *errors.Error) {
	if ret.ErrCode != expected {
		t.Errorf("Expected response errcode %d. Got %d\n",
			expected, ret.ErrCode)
	}
}

func TestHTTP(t *testing.T) {
	// 路径带前缀'/api'
	t.Run("/api/v1/hello", sayHello("/api/v1/hello"))
	t.Run("/api/v1/pingpong", pingPong("/api/v1/pingpong"))
}
func TestGoMicroGrpcGateway(t *testing.T) {
	// 路径不带前缀'/api'
	t.Run("/v1/hello", sayHello("/v1/hello"))
	t.Run("/v1/pingpong", pingPong("/v1/pingpong"))
}

func sayHello(url string) func(t *testing.T) {
	return func(t *testing.T) {
		type args struct {
			req *proto.Request
			ret *errors.Error
		}
		tests := []struct {
			name           string
			args           args
			wantStatusCode int
			wantErrCode    errors.ErrorCode
			wantRsp        *proto.Reply
		}{
			// TODO: Add test cases.
			{
				"name_field",
				args{
					&proto.Request{
						Name: "mark",
						Age:  21,
					},
					&errors.Error{},
				},
				http.StatusOK,
				errors.ECodeSuccessed,
				&proto.Reply{
					Message: "hello world, mark.",
				},
			}, {
				"age_field_validated_failed",
				args{
					&proto.Request{
						Name: "mark",
						Age:  2,
					},
					&errors.Error{},
				},
				http.StatusOK,
				errors.ECodeInvalidParams,
				&proto.Reply{},
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				jsonStr, _ := utils.Marshal(tt.args.req)
				req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
				if err != nil {
					t.Fatal(err)
				}
				req.Header.Set("Content-Type", "application/json")
				rr := executeRequest(req)
				checkResponseStatusCode(t, tt.wantStatusCode, rr.Code)
				rsp := &proto.Reply{}
				tt.args.ret.Data = rsp
				err = utils.Unmarshal(rr.Body.Bytes(), tt.args.ret)
				if err != nil {
					t.Fatal(err)
				}
				checkErrCode(t, tt.wantErrCode, tt.args.ret)
				if rsp.Message != tt.wantRsp.Message {
					t.Errorf("Expected response rsp %v. Got %v\n",
						tt.wantRsp, rsp)
				}
			})
		}
	}
}

func pingPong(url string) func(t *testing.T) {
	return func(t *testing.T) {
		type args struct {
			req *proto.PingRequest
			ret *errors.Error
		}
		tests := []struct {
			name           string
			args           args
			wantStatusCode int
			wantErrCode    errors.ErrorCode
			wantRsp        *proto.PongReply
		}{
			// TODO: Add test cases.
			{
				"pp",
				args{
					&proto.PingRequest{
						Ping: "ping",
					},
					&errors.Error{},
				},
				http.StatusOK,
				errors.ECodeSuccessed,
				&proto.PongReply{
					Pong: "pong",
				},
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				jsonStr, _ := utils.Marshal(tt.args.req)
				req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
				if err != nil {
					t.Fatal(err)
				}
				req.Header.Set("Content-Type", "application/json")
				rr := executeRequest(req)
				checkResponseStatusCode(t, tt.wantStatusCode, rr.Code)
				rsp := &proto.PongReply{}
				tt.args.ret.Data = rsp
				err = utils.Unmarshal(rr.Body.Bytes(), tt.args.ret)
				if err != nil {
					t.Fatal(err)
				}
				checkErrCode(t, tt.wantErrCode, tt.args.ret)
				if rsp.Pong != tt.wantRsp.Pong {
					t.Errorf("Expected response rsp %v. Got %v\n",
						tt.wantRsp, rsp)
				}
			})
		}
	}
}
