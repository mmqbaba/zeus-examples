package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	// assetfs "github.com/elazarl/go-bindata-assetfs"

	"github.com/gorilla/mux"
	gruntime "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	gmerrors "github.com/micro/go-micro/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	// "github.com/urfave/negroni"

	"github.com/mmqbaba/zeus/config"
	"github.com/mmqbaba/zeus/engine"
	"github.com/mmqbaba/zeus/engine/etcd"
	"github.com/mmqbaba/zeus/engine/file"
	zeuserrors "github.com/mmqbaba/zeus/errors"
	zgomicro "github.com/mmqbaba/zeus/microsrv/gomicro"
	"github.com/mmqbaba/zeus/plugin/zcontainer"

	// swaggerthirdparty "github.com/mmqbaba/zeus/swagger/third_party"
	swaggerhelper "github.com/mmqbaba/zeus/swagger"
	// swagger "github.com/mmqbaba/zeus/swagger/ui"
	"github.com/mmqbaba/zeus/utils"
)

const (
	retryPeriod       = 5 * time.Second
	changesBufferSize = 10
)

var confEntry *config.Entry
var confEntryPath string
var engineProvidors map[string]engine.NewEngineFn
var swaggerDir string

func init() {
	log.SetPrefix("[zeus] ")
	log.SetFlags(3)
	// 此处使用os.Stdout覆盖回来
	log.SetOutput(os.Stdout)

	os := runtime.GOOS
	if os == "linux" || os == "darwin" {
		confEntryPath = "/etc/zeus/zeus.json"
	} else {
		confEntryPath = "C:\\zeus\\zeus.json"
	}

	engineProvidors = map[string]engine.NewEngineFn{
		"etcd": newEtcdEngine,
		"file": newFileEngine,
	}
}

func newEtcdEngine(cnt zcontainer.Container) (engine.Engine, error) {
	return etcd.New(confEntry, cnt)
}

// newFileEngine fileengine的实现，目前并不完善，不要应用到生产，可简单用在开发测试
func newFileEngine(cnt zcontainer.Container) (engine.Engine, error) {
	return file.New(confEntry, cnt)
}

func Run(cnt zcontainer.Container, conf *Options, opts ...Option) (err error) {
	opt, err := ParseCommandLine()
	if err != nil {
		log.Printf("[zeus] [service.Run] err: %s\n", err)
		return
	}
	if conf != nil {
		opt = *conf
	}
	s := NewService(opt, cnt, opts...)

	if err = s.Init(); err != nil {
		return
	}
	if err = s.RunServer(); err != nil {
		log.Printf("[zeus] [service.Run] err: %s\n", err)
		s.watcherCancelC <- struct{}{} // 服务启动失败，通知停止engine的监听
		return
	}
	return
}

func (s *Service) Init() (err error) {
	s.initConfEntry()

	if fn, ok := engineProvidors[confEntry.EngineType]; ok && fn != nil {
		if s.ng, err = fn(s.container); err != nil {
			log.Printf("[zeus] [service.Run] err: %s\n", err)
		}
	} else {
		err = fmt.Errorf("no newEnginFn providor for engineType: %s", confEntry.EngineType)
		log.Printf("[zeus] [service.Run] err: %s\n", err)
		return
	}

	// 启动engine，监听状态变化，处理更新，各种组件状态变化
	if err = s.loadEngine(); err != nil {
		log.Printf("[zeus] [service.Run] err: %s\n", err)
		return
	}

	if err = s.initServer(); err != nil {
		log.Printf("[zeus] [service.Run] err: %s\n", err)
		s.watcherCancelC <- struct{}{} // 服务启动失败，通知停止engine的监听
		return
	}

	// 触发服务初始化完成事件
	if s.options.InitServiceCompleteFn != nil {
		utils.AsyncFuncSafe(context.Background(), func(args ...interface{}) {
			s.options.InitServiceCompleteFn(s.ng)
		}, nil)
	}

	return
}

type Service struct {
	options        *Options
	container      zcontainer.Container
	ng             engine.Engine
	watcherCancelC chan struct{}
	watcherErrorC  chan struct{}
	watcherWg      sync.WaitGroup
}

func NewService(options Options, container zcontainer.Container, opts ...Option) *Service {
	o := options
	for _, opt := range opts {
		if opt != nil {
			opt(&o)
		}
	}
	s := &Service{
		options:        &o,
		container:      container,
		watcherErrorC:  make(chan struct{}),
		watcherCancelC: make(chan struct{}),
	}
	if !utils.IsEmptyString(options.ConfEntryPath) {
		confEntryPath = options.ConfEntryPath
	}

	return s
}

func (s *Service) initConfEntry() {
	b, err := ioutil.ReadFile(confEntryPath)
	if err != nil {
		panic(err)
	}
	confEntry = &config.Entry{}
	if err := json.Unmarshal(b, confEntry); err != nil {
		panic(err)
	}
	log.Printf("[zeus] [service.initConfEntry] confEntry: %+v", confEntry)
	return
}

// loadEngine 初始化engine，开启监听
func (s *Service) loadEngine() (err error) {
	if err = s.ng.Init(); err != nil {
		log.Println("[zeus] [service.loadEngine] s.ng.Init err:", err)
		return
	}

	configer, err := s.ng.GetConfiger()
	if err != nil {
		return
	}
	// 初始化容器组件
	conf := *configer.Get()
	if !utils.IsEmptyString(s.options.Log) {
		conf.LogConf.Log = s.options.Log
	}
	if !utils.IsEmptyString(s.options.LogFormat) {
		conf.LogConf.Format = s.options.LogFormat
	}
	if !utils.IsEmptyString(s.options.LogLevel) {
		conf.LogConf.Level = s.options.LogLevel
	}
	s.container.Init(&conf)

	changesC := make(chan interface{}, changesBufferSize)
	// 监听配置变化
	go func() {
		defer close(changesC)
		defer close(s.watcherCancelC)
		if err := s.ng.Subscribe(changesC, s.watcherCancelC); err != nil {
			log.Println("[zeus] [s.n.subscribe] err:", err)
			s.watcherErrorC <- struct{}{}
			return
		}
	}()

	// 处理事件
	go func() {
		defer close(s.watcherErrorC)
		for {
			select {
			case change := <-changesC:
				if err := s.processChange(change); err != nil {
					log.Printf("[zeus] failed to processChange, change=%#v, err=%s\n", change, err)
				}
			case <-s.watcherErrorC:
				log.Println("[zeus] watcher error, change processor shutdown")
				return
			}
		}
	}()

	if s.options.LoadEngineFn != nil {
		s.options.LoadEngineFn(s.ng)
	}
	return
}

func (s *Service) processChange(ev interface{}) (err error) {
	switch c := ev.(type) {
	case config.Configer:
		log.Printf("[zeus] config change\n")
		// 重新加载容器组件
		conf := *c.Get()
		if !utils.IsEmptyString(s.options.Log) {
			conf.LogConf.Log = s.options.Log
		}
		if !utils.IsEmptyString(s.options.LogFormat) {
			conf.LogConf.Format = s.options.LogFormat
		}
		if !utils.IsEmptyString(s.options.LogLevel) {
			conf.LogConf.Level = s.options.LogLevel
		}
		s.container.Reload(&conf)
	default:
		log.Printf("[zeus] unsupported event change\n")
	}
	if s.options.ProcessChangeFn != nil {
		utils.AsyncFuncSafe(context.Background(), func(args ...interface{}) {
			s.options.ProcessChangeFn(ev)
		}, nil)
	}
	return
}

type gwOption struct {
	grpcEndpoint        string
	swaggerJSONFile     string
	setSwaggerServiceFn func(name string)
}

func (s *Service) initServer() (err error) {
	// gomicro-grpc and gw-http
	log.Println("[zeus] start server ...")
	configer, err := s.ng.GetConfiger()
	if err != nil {
		log.Println("[zeus] [s.startServer] err:", err)
		return
	}
	microConf := configer.Get().GoMicro
	serverPort := microConf.ServerPort
	if s.options.Port > 0 { // 优先使用命令行传递的值
		serverPort = uint32(s.options.Port)
	}
	microConf.ServerPort = serverPort // 最终值为0，使用随机端口

	// 注意：Advertise不要使用etcd里的公共配置
	microConf.Advertise = s.options.Advertise

	serviceName := microConf.ServiceName
	if utils.IsEmptyString(serviceName) {
		serviceName = s.options.ServiceName
	}
	microConf.ServiceName = serviceName
	newHTTPApiServerOpt := micro.AfterStart(func() (err error) {
		addr := s.container.GetGoMicroService().Server().Options().Address
		gw, err := s.newHTTPGateway(gwOption{
			// grpcEndpoint:    fmt.Sprintf("localhost:%d", serverPort),
			grpcEndpoint:        addr,
			swaggerJSONFile:     s.options.SwaggerJSONFileName,
			setSwaggerServiceFn: s.options.SetSwaggerServiceFn,
		})
		if err != nil {
			log.Fatal(err)
			return
		}
		s.container.SetHTTPHandler(gw)

		go func() {
			addr := fmt.Sprintf("%s:%d", s.options.ApiInterface, s.options.ApiPort)
			// log.Printf("http apiserver listen on %s", addr)
			// Start HTTP server (and proxy calls to gRPC server endpoint, serve http, serve swagger)
			// if err := http.ListenAndServe(addr, gw); err != nil {
			// 	log.Fatal(err)
			// }
			srv := &http.Server{
				Handler: gw,
			}
			ln, err := net.Listen("tcp", addr)
			if err != nil {
				log.Fatal(err)
				return
			}
			// host, port, err := net.SplitHostPort(ln.Addr().String())
			// log.Println(host, port)
			log.Printf("http apiserver listen on %s\n", ln.Addr())
			if err := srv.Serve(ln); err != nil {
				log.Fatal(err)
				return
			}
		}()
		return
	})
	gomicroservice, err := s.newGomicroSrv(microConf, newHTTPApiServerOpt)
	if err != nil {
		return
	}
	s.container.SetGoMicroService(gomicroservice)
	return
}

func (s *Service) RunServer() (err error) {
	gomicroservice := s.container.GetGoMicroService()
	// Run the service
	if err = gomicroservice.Run(); err != nil {
		log.Println("[zeus] err:", err)
		return
	}
	return
}

func (s *Service) newGomicroSrv(conf config.GoMicro, srvopts ...micro.Option) (gms micro.Service, err error) {
	var gomicroservice micro.Service
	opts := []micro.Option{
		micro.WrapHandler(zgomicro.GenerateServerLogWrap(s.ng)), // 保证serverlogwrap在最前
	}
	if len(s.options.GoMicroServerWrapGenerateFn) != 0 {
		for _, fn := range s.options.GoMicroServerWrapGenerateFn {
			sw := fn(s.ng)
			if sw != nil {
				opts = append(opts, micro.WrapHandler(sw))
			}
		}
	}
	// new micro client
	cliOpts := []client.Option{
		client.Wrap(zgomicro.GenerateClientLogWrap(s.ng)), // 保证在最前
	}
	if len(s.options.GoMicroClientWrapGenerateFn) != 0 {
		for _, fn := range s.options.GoMicroClientWrapGenerateFn {
			cw := fn(s.ng)
			if cw != nil {
				cliOpts = append(cliOpts, client.Wrap(cw))
			}
		}
	}
	cli, err := zgomicro.NewClient(context.Background(), conf, cliOpts...)
	if err != nil {
		log.Println("[zeus] [s.newGomicroSrv] zgomicro.NewClient err:", err)
		return
	}
	// 把client设置到container
	s.ng.GetContainer().SetGoMicroClient(cli)
	opts = append(opts, micro.Client(cli))
	opts = append(opts, micro.AfterStart(func() error {
		serviceID := gomicroservice.Server().Options().Name + "-" + gomicroservice.Server().Options().Id
		log.Println("[gomicro] afterstart", serviceID)
		s.container.SetServiceID(serviceID)
		return nil
	}))
	opts = append(opts, srvopts...)
	// new micro service
	gomicroservice = zgomicro.NewService(context.Background(), conf, opts...)
	if s.options.GoMicroHandlerRegisterFn != nil {
		if err = s.options.GoMicroHandlerRegisterFn(gomicroservice.Server()); err != nil {
			log.Println("[zeus] [s.newGomicroSrv] GoMicroHandlerRegister err:", err)
			return
		}
		log.Println("[zeus] [s.newGomicroSrv] GoMicroHandlerRegister success.")
	}
	gms = gomicroservice
	return
}

type gwBodyWriter struct {
	http.ResponseWriter
	done    bool
	status  int
	zeusErr *zeuserrors.Error
	body    *bytes.Buffer
}

// 这里使用指针实现，传递指针，保证done status值的变更传递
func (w *gwBodyWriter) WriteHeader(status int) {
	w.done = true
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *gwBodyWriter) Write(b []byte) (l int, err error) {
	// 处理错误
	if w.zeusErr != nil && w.zeusErr.ErrCode != zeuserrors.ECodeSuccessed {
		return w.ResponseWriter.Write(b)
	}

	if w.done {
		// Something already wrote a response
		// status already wrote
		return w.ResponseWriter.Write(b)
	}
	// 正常返回
	buf := bytes.NewBufferString(`{"errcode":` + strconv.Itoa(int(zeuserrors.ECodeSuccessed)) + `,"errmsg":"ok","data":`)
	buf.Write(b)
	buf.WriteString(`}`)
	return w.ResponseWriter.Write(buf.Bytes())
}

func grpcGatewayHTTPError(ctx context.Context, mux *gruntime.ServeMux, marshaler gruntime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
	const fallback = `{"error": "failed to marshal error message"}`

	s, ok := status.FromError(err)
	if !ok {
		s = status.New(codes.Unknown, err.Error())
	}

	w.Header().Del("Trailer")

	contentType := marshaler.ContentType()
	// Check marshaler on run time in order to keep backwards compatability
	// An interface param needs to be added to the ContentType() function on
	// the Marshal interface to be able to remove this check
	if httpBodyMarshaler, ok := marshaler.(*gruntime.HTTPBodyMarshaler); ok {
		pb := s.Proto()
		contentType = httpBodyMarshaler.ContentTypeFromMessage(pb)
	}
	w.Header().Set("Content-Type", contentType)

	msg := s.Message()
	if !utils.IsEmptyString(msg) && s.Code() != 0 {
		gmErr := gmerrors.Error{}
		if e := utils.Unmarshal([]byte(msg), &gmErr); e != nil {
			log.Println("utils.Unmarshal err:", e)
		}
		if gmErr.Code != 0 {
			// w.Header().Set("x-zeus-errcode", strconv.Itoa(int(gmErr.Code)))
			zeusErr := zeuserrors.New(zeuserrors.ErrorCode(gmErr.Code), gmErr.Detail, gmErr.Status)
			ww, ok := w.(*gwBodyWriter)
			if ok {
				ww.zeusErr = zeusErr
			}
			err = zeusErr.Write(w)
			return
		}
	}

	zeuserrors.ECodeProxyFailed.ParseErr(msg).Write(w)

	// body := &struct {
	// 	Error   string     `protobuf:"bytes,100,name=error" json:"error"`
	// 	Code    int32      `protobuf:"varint,1,name=code" json:"code"`
	// 	Message string     `protobuf:"bytes,2,name=message" json:"message"`
	// 	Details []*any.Any `protobuf:"bytes,3,rep,name=details" json:"details,omitempty"`
	// }{
	// 	Error:   s.Message(),
	// 	Message: s.Message(),
	// 	Code:    int32(s.Code()),
	// 	Details: s.Proto().GetDetails(),
	// }

	// _, ok = gruntime.ServerMetadataFromContext(ctx)
	// if !ok {
	// 	log.Println("Failed to extract ServerMetadata from context")
	// }
}

func (s *Service) newHTTPGateway(opt gwOption) (h http.Handler, err error) {
	configer, err := s.ng.GetConfiger()
	if err != nil {
		log.Println("[zeus] [s.newHTTPGateway] s.ng.GetConfiger err:", err)
		return
	}
	conf := configer.Get()

	// 必须注意r.PathPrefix的顺序问题
	r := mux.NewRouter()

	// swagger handler
	// r.PathPrefix("/swagger/").HandlerFunc(serveSwaggerFileHandle)
	serveSwaggerFile("/swagger/", r, http.FS(swaggerhelper.Swaggerfile))

	// if opt.setSwaggerServiceFn != nil {
	// 	opt.setSwaggerServiceFn(opt.swaggerJSONFile)
	// }

	// serveSwaggerUI("/swagger-ui/", true, r, swaggerhelper.SwaggerAssetFS) // go-bindata

	r.Path("/swagger-ui/").Methods("GET").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p := strings.TrimPrefix(r.URL.Path, "/") + "index.html"
		data, err := swaggerhelper.SwaggerUI.ReadFile(p)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		s := strings.Replace(string(data), "{DEFAULT_SERVICE}", opt.swaggerJSONFile, 1)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(s))
	})
	serveSwaggerUI("/swagger-ui/", false, r, http.FS(swaggerhelper.SwaggerUI)) // embed

	log.Println("[zeus] [s.newHTTPGateway] swaggerRegister success.")

	// http handler
	if s.options.HttpHandlerRegisterFn != nil {
		var handler http.Handler
		handlerPrefix := ""
		if conf != nil {
			if v, ok := conf.Ext["httphandler_pathprefix"]; ok {
				handlerPrefix = fmt.Sprint(v)
			}
		}
		if utils.IsEmptyString(handlerPrefix) {
			// handlerPrefix = "/zeus/"
			handlerPrefix = "/api/"
		}
		if handler, err = s.options.HttpHandlerRegisterFn(context.Background(), handlerPrefix, s.ng); err != nil {
			log.Println("[zeus] [s.newHTTPGateway] HttpHandlerRegister err:", err)
			return
		}
		if handler != nil {
			r.PathPrefix(handlerPrefix).Handler(handler)
			log.Println("[zeus] [s.newHTTPGateway] HttpHandlerRegister success.")
		}
	}

	// gateway handler
	if s.options.HttpGWHandlerRegisterFn != nil {
		var gwmux *gruntime.ServeMux
		if gwmux, err = s.options.HttpGWHandlerRegisterFn(context.Background(), opt.grpcEndpoint, nil); err != nil {
			log.Println("[zeus] [s.newHTTPGateway] HttpGWHandlerRegister err:", err)
			return
		}
		gruntime.HTTPError = grpcGatewayHTTPError // 覆盖默认的错误处理函数
		if gwmux != nil {
			gwPrefix := ""
			if conf != nil {
				if v, ok := conf.Ext["grpcgateway_pathprefix"]; ok {
					gwPrefix = fmt.Sprint(v)
				}
			}
			if utils.IsEmptyString(gwPrefix) {
				gwPrefix = "/"
			}
			r.PathPrefix(gwPrefix).HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
				rr := r.WithContext(r.Context())
				rr.URL.Path = strings.Replace(r.URL.Path, gwPrefix, "/", 1)
				bwriter := &gwBodyWriter{body: bytes.NewBufferString(""), ResponseWriter: rw}
				gwmux.ServeHTTP(bwriter, rr)
			})
			log.Println("[zeus] [s.newHTTPGateway] HttpGWHandlerRegister success.")
		}
	}

	// r := http.NewServeMux()
	// r.HandleFunc("/swagger/", serveSwaggerFileHandle)
	// serveGin(r)
	// r.Handle("/", gwmux)

	// mux + negroni 可实现完整的前后置处理和路由
	// r.WithContext + r.Context() 实现上下文传递
	// n := negroni.New()
	// n.UseFunc(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc){
	// 	c := r.Context()
	// 	rr := r.WithContext(context.WithValue(c, "a", "b")) // 上下文传递
	// 	log.Println("========================1 begin")
	// 	next(rw, rr)
	// 	log.Println("========================1 end")
	// })
	// n.UseFunc(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc){
	// 	_ := r.Context().Value("a") // 读取
	// 	log.Println("========================2 begin")
	// 	next(rw, r)
	// 	log.Println("========================2 end")
	// })
	// n.UseHandlerFunc(func(rw http.ResponseWriter, r *http.Request){
	// 	log.Println("========================process begin")
	// 	rw.Write([]byte("hello, abc user."))
	// 	log.Println("========================process end")
	// })
	// r.Path("/abc/user").Methods("GET").HandlerFunc(n.ServeHTTP)

	h = r
	return
}

func serveSwaggerFileHandle(w http.ResponseWriter, r *http.Request) {
	if !strings.HasSuffix(r.URL.Path, "swagger.json") {
		log.Printf("Not Found: %s", r.URL.Path)
		http.NotFound(w, r)
		return
	}

	swaggerDir = "proto"
	p := strings.TrimPrefix(r.URL.Path, "/swagger/")
	p = path.Join(swaggerDir, p)

	log.Printf("Serving swagger-file: %s", p)

	http.ServeFile(w, r, p)
}

func serveSwaggerFile(prefix string, mux *mux.Router, f http.FileSystem) {
	fileServer := http.FileServer(f)
	mux.PathPrefix(prefix).Handler(http.StripPrefix(prefix, fileServer))
}

func serveSwaggerUI(prefix string, stripPrefix bool, mux *mux.Router, f http.FileSystem) {
	fileServer := http.FileServer(f)
	if stripPrefix {
		mux.PathPrefix(prefix).Handler(http.StripPrefix(prefix, fileServer))
		return
	}
	mux.PathPrefix(prefix).Handler(fileServer)
}
