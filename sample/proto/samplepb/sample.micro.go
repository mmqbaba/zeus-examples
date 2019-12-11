// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: sample.proto

package sample

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/any"
	_ "github.com/golang/protobuf/ptypes/struct"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
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

// Client API for Sample service

type SampleService interface {
	SayHello(ctx context.Context, in *Request, opts ...client.CallOption) (*Reply, error)
	PingPong(ctx context.Context, in *PingRequest, opts ...client.CallOption) (*PongReply, error)
	Upload(ctx context.Context, opts ...client.CallOption) (Sample_UploadService, error)
}

type sampleService struct {
	c    client.Client
	name string
}

func NewSampleService(name string, c client.Client) SampleService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "sample"
	}
	return &sampleService{
		c:    c,
		name: name,
	}
}

func (c *sampleService) SayHello(ctx context.Context, in *Request, opts ...client.CallOption) (*Reply, error) {
	req := c.c.NewRequest(c.name, "Sample.SayHello", in)
	out := new(Reply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sampleService) PingPong(ctx context.Context, in *PingRequest, opts ...client.CallOption) (*PongReply, error) {
	req := c.c.NewRequest(c.name, "Sample.PingPong", in)
	out := new(PongReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sampleService) Upload(ctx context.Context, opts ...client.CallOption) (Sample_UploadService, error) {
	req := c.c.NewRequest(c.name, "Sample.Upload", &UploadReq{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &sampleServiceUpload{stream}, nil
}

type Sample_UploadService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*UploadReq) error
}

type sampleServiceUpload struct {
	stream client.Stream
}

func (x *sampleServiceUpload) Close() error {
	return x.stream.Close()
}

func (x *sampleServiceUpload) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *sampleServiceUpload) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *sampleServiceUpload) Send(m *UploadReq) error {
	return x.stream.Send(m)
}

// Server API for Sample service

type SampleHandler interface {
	SayHello(context.Context, *Request, *Reply) error
	PingPong(context.Context, *PingRequest, *PongReply) error
	Upload(context.Context, Sample_UploadStream) error
}

func RegisterSampleHandler(s server.Server, hdlr SampleHandler, opts ...server.HandlerOption) error {
	type sample interface {
		SayHello(ctx context.Context, in *Request, out *Reply) error
		PingPong(ctx context.Context, in *PingRequest, out *PongReply) error
		Upload(ctx context.Context, stream server.Stream) error
	}
	type Sample struct {
		sample
	}
	h := &sampleHandler{hdlr}
	return s.Handle(s.NewHandler(&Sample{h}, opts...))
}

type sampleHandler struct {
	SampleHandler
}

func (h *sampleHandler) SayHello(ctx context.Context, in *Request, out *Reply) error {
	return h.SampleHandler.SayHello(ctx, in, out)
}

func (h *sampleHandler) PingPong(ctx context.Context, in *PingRequest, out *PongReply) error {
	return h.SampleHandler.PingPong(ctx, in, out)
}

func (h *sampleHandler) Upload(ctx context.Context, stream server.Stream) error {
	return h.SampleHandler.Upload(ctx, &sampleUploadStream{stream})
}

type Sample_UploadStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*UploadReq, error)
}

type sampleUploadStream struct {
	stream server.Stream
}

func (x *sampleUploadStream) Close() error {
	return x.stream.Close()
}

func (x *sampleUploadStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *sampleUploadStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *sampleUploadStream) Recv() (*UploadReq, error) {
	m := new(UploadReq)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
