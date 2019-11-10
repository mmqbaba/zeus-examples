// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: hello.proto

package hello

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Hello service

type HelloService interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...client.CallOption) (*HelloReply, error)
	// PingPong
	PingPong(ctx context.Context, in *PingRequest, opts ...client.CallOption) (*PongReply, error)
	// 上传
	Upload(ctx context.Context, opts ...client.CallOption) (Hello_UploadService, error)
}

type helloService struct {
	c    client.Client
	name string
}

func NewHelloService(name string, c client.Client) HelloService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "hello"
	}
	return &helloService{
		c:    c,
		name: name,
	}
}

func (c *helloService) SayHello(ctx context.Context, in *HelloRequest, opts ...client.CallOption) (*HelloReply, error) {
	req := c.c.NewRequest(c.name, "Hello.SayHello", in)
	out := new(HelloReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloService) PingPong(ctx context.Context, in *PingRequest, opts ...client.CallOption) (*PongReply, error) {
	req := c.c.NewRequest(c.name, "Hello.PingPong", in)
	out := new(PongReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloService) Upload(ctx context.Context, opts ...client.CallOption) (Hello_UploadService, error) {
	req := c.c.NewRequest(c.name, "Hello.Upload", &UploadReq{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &helloServiceUpload{stream}, nil
}

type Hello_UploadService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*UploadReq) error
}

type helloServiceUpload struct {
	stream client.Stream
}

func (x *helloServiceUpload) Close() error {
	return x.stream.Close()
}

func (x *helloServiceUpload) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *helloServiceUpload) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *helloServiceUpload) Send(m *UploadReq) error {
	return x.stream.Send(m)
}

// Server API for Hello service

type HelloHandler interface {
	SayHello(context.Context, *HelloRequest, *HelloReply) error
	// PingPong
	PingPong(context.Context, *PingRequest, *PongReply) error
	// 上传
	Upload(context.Context, Hello_UploadStream) error
}

func RegisterHelloHandler(s server.Server, hdlr HelloHandler, opts ...server.HandlerOption) error {
	type hello interface {
		SayHello(ctx context.Context, in *HelloRequest, out *HelloReply) error
		PingPong(ctx context.Context, in *PingRequest, out *PongReply) error
		Upload(ctx context.Context, stream server.Stream) error
	}
	type Hello struct {
		hello
	}
	h := &helloHandler{hdlr}
	return s.Handle(s.NewHandler(&Hello{h}, opts...))
}

type helloHandler struct {
	HelloHandler
}

func (h *helloHandler) SayHello(ctx context.Context, in *HelloRequest, out *HelloReply) error {
	return h.HelloHandler.SayHello(ctx, in, out)
}

func (h *helloHandler) PingPong(ctx context.Context, in *PingRequest, out *PongReply) error {
	return h.HelloHandler.PingPong(ctx, in, out)
}

func (h *helloHandler) Upload(ctx context.Context, stream server.Stream) error {
	return h.HelloHandler.Upload(ctx, &helloUploadStream{stream})
}

type Hello_UploadStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*UploadReq, error)
}

type helloUploadStream struct {
	stream server.Stream
}

func (x *helloUploadStream) Close() error {
	return x.stream.Close()
}

func (x *helloUploadStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *helloUploadStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *helloUploadStream) Recv() (*UploadReq, error) {
	m := new(UploadReq)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
