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

func checkRsp(t *testing.T, expected, actual interface{}) {
	r, err := utils.Marshal(actual)
	if err != nil {
		t.Fatal(err)
	}
	w, err := utils.Marshal(expected)
	if err != nil {
		t.Fatal(err)
	}
	if string(r) != string(w) {
		t.Errorf("Expected response rsp %v. Got %v\n",
			expected, actual)
	}
}

func TestHTTP(t *testing.T) {
	// 路径带前缀'/api'
	runSayHello(t, "/api/v1/hello")
	runPingPong(t, "/api/v1/pingpong")
}
func TestGoMicroGrpcGateway(t *testing.T) {
	// 路径不带前缀'/api'
	runSayHello(t, "/v1/hello")
	runPingPong(t, "/v1/pingpong")
}

func runSayHello(t *testing.T, url string) {
	t.Run(url, func(t *testing.T) {
		type args struct {
			req interface{}
			ret *errors.Error
		}
		tests := []struct {
			name           string
			args           args
			wantStatusCode int
			wantErrCode    errors.ErrorCode
			wantRsp        interface{}
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
				checkRsp(t, tt.wantRsp, rsp)
			})
		}
	})
}

func runPingPong(t *testing.T, url string) {
	t.Run(url, func(t *testing.T) {
		type args struct {
			req interface{}
			ret *errors.Error
		}
		tests := []struct {
			name           string
			args           args
			wantStatusCode int
			wantErrCode    errors.ErrorCode
			wantRsp        interface{}
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
				checkRsp(t, tt.wantRsp, rsp)
			})
		}
	})
}
