// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello.proto

package hello

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type HelloRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// @inject_tag: validate:"required,gt=0,max=1000"
	Age                  int32    `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty" validate:"required,gt=0,max=1000"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *HelloRequest) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

type HelloReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReply) Reset()         { *m = HelloReply{} }
func (m *HelloReply) String() string { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()    {}
func (*HelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{1}
}

func (m *HelloReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReply.Unmarshal(m, b)
}
func (m *HelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReply.Marshal(b, m, deterministic)
}
func (m *HelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReply.Merge(m, src)
}
func (m *HelloReply) XXX_Size() int {
	return xxx_messageInfo_HelloReply.Size(m)
}
func (m *HelloReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReply.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReply proto.InternalMessageInfo

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type PingRequest struct {
	Ping                 string   `protobuf:"bytes,1,opt,name=ping,proto3" json:"ping,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{2}
}

func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRequest.Unmarshal(m, b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
}
func (m *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(m, src)
}
func (m *PingRequest) XXX_Size() int {
	return xxx_messageInfo_PingRequest.Size(m)
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

func (m *PingRequest) GetPing() string {
	if m != nil {
		return m.Ping
	}
	return ""
}

type PongReply struct {
	Pong                 string   `protobuf:"bytes,1,opt,name=pong,proto3" json:"pong,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PongReply) Reset()         { *m = PongReply{} }
func (m *PongReply) String() string { return proto.CompactTextString(m) }
func (*PongReply) ProtoMessage()    {}
func (*PongReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{3}
}

func (m *PongReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PongReply.Unmarshal(m, b)
}
func (m *PongReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PongReply.Marshal(b, m, deterministic)
}
func (m *PongReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PongReply.Merge(m, src)
}
func (m *PongReply) XXX_Size() int {
	return xxx_messageInfo_PongReply.Size(m)
}
func (m *PongReply) XXX_DiscardUnknown() {
	xxx_messageInfo_PongReply.DiscardUnknown(m)
}

var xxx_messageInfo_PongReply proto.InternalMessageInfo

func (m *PongReply) GetPong() string {
	if m != nil {
		return m.Pong
	}
	return ""
}

type UploadReq struct {
	Content  []byte `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	FileName string `protobuf:"bytes,2,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	// 文件大小
	FileSize int64  `protobuf:"varint,3,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"`
	MimeType string `protobuf:"bytes,4,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	// 上传完成
	Finished             bool     `protobuf:"varint,5,opt,name=finished,proto3" json:"finished,omitempty"`
	SpId                 string   `protobuf:"bytes,6,opt,name=sp_id,json=spId,proto3" json:"sp_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadReq) Reset()         { *m = UploadReq{} }
func (m *UploadReq) String() string { return proto.CompactTextString(m) }
func (*UploadReq) ProtoMessage()    {}
func (*UploadReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{4}
}

func (m *UploadReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadReq.Unmarshal(m, b)
}
func (m *UploadReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadReq.Marshal(b, m, deterministic)
}
func (m *UploadReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadReq.Merge(m, src)
}
func (m *UploadReq) XXX_Size() int {
	return xxx_messageInfo_UploadReq.Size(m)
}
func (m *UploadReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadReq.DiscardUnknown(m)
}

var xxx_messageInfo_UploadReq proto.InternalMessageInfo

func (m *UploadReq) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *UploadReq) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func (m *UploadReq) GetFileSize() int64 {
	if m != nil {
		return m.FileSize
	}
	return 0
}

func (m *UploadReq) GetMimeType() string {
	if m != nil {
		return m.MimeType
	}
	return ""
}

func (m *UploadReq) GetFinished() bool {
	if m != nil {
		return m.Finished
	}
	return false
}

func (m *UploadReq) GetSpId() string {
	if m != nil {
		return m.SpId
	}
	return ""
}

type UploadResp struct {
	// 文件id
	FileId               string   `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadResp) Reset()         { *m = UploadResp{} }
func (m *UploadResp) String() string { return proto.CompactTextString(m) }
func (*UploadResp) ProtoMessage()    {}
func (*UploadResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{5}
}

func (m *UploadResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadResp.Unmarshal(m, b)
}
func (m *UploadResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadResp.Marshal(b, m, deterministic)
}
func (m *UploadResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadResp.Merge(m, src)
}
func (m *UploadResp) XXX_Size() int {
	return xxx_messageInfo_UploadResp.Size(m)
}
func (m *UploadResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadResp.DiscardUnknown(m)
}

var xxx_messageInfo_UploadResp proto.InternalMessageInfo

func (m *UploadResp) GetFileId() string {
	if m != nil {
		return m.FileId
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "hello.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "hello.HelloReply")
	proto.RegisterType((*PingRequest)(nil), "hello.PingRequest")
	proto.RegisterType((*PongReply)(nil), "hello.PongReply")
	proto.RegisterType((*UploadReq)(nil), "hello.UploadReq")
	proto.RegisterType((*UploadResp)(nil), "hello.UploadResp")
}

func init() { proto.RegisterFile("hello.proto", fileDescriptor_61ef911816e0a8ce) }

var fileDescriptor_61ef911816e0a8ce = []byte{
	// 507 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0x95, 0xe3, 0x24, 0xb5, 0x6f, 0xa2, 0xef, 0x6b, 0xa6, 0x45, 0xb5, 0x5c, 0xa4, 0x86, 0x91,
	0x40, 0x56, 0xa4, 0xc6, 0x02, 0x24, 0x16, 0x5d, 0x20, 0x44, 0x11, 0x34, 0x1b, 0x14, 0xb9, 0xb0,
	0x8e, 0xdc, 0x7a, 0xea, 0x8c, 0xb0, 0x67, 0xa6, 0x9d, 0x49, 0x2b, 0x77, 0xc9, 0x2b, 0xf0, 0x1c,
	0x48, 0xbc, 0x0b, 0x0f, 0x80, 0x84, 0x78, 0x10, 0x74, 0xc7, 0x71, 0xea, 0xee, 0xe8, 0xee, 0xfe,
	0x9c, 0x7b, 0xee, 0x19, 0xdf, 0x63, 0x18, 0x2c, 0x59, 0x51, 0xc8, 0xa9, 0xba, 0x92, 0x46, 0x92,
	0x9e, 0x4d, 0xc2, 0xc7, 0xb9, 0x94, 0x79, 0xc1, 0xe2, 0x54, 0xf1, 0x38, 0x15, 0x42, 0x9a, 0xd4,
	0x70, 0x29, 0x74, 0x0d, 0x0a, 0x5f, 0xe5, 0xdc, 0x2c, 0x57, 0x67, 0xd3, 0x73, 0x59, 0xc6, 0xe5,
	0x0d, 0x37, 0x5f, 0xe4, 0x4d, 0x9c, 0xcb, 0x43, 0xdb, 0x3c, 0xbc, 0x4e, 0x0b, 0x9e, 0xa5, 0x46,
	0x5e, 0xe9, 0x78, 0x13, 0xd6, 0x73, 0xf4, 0x35, 0x0c, 0x4f, 0x90, 0x3e, 0x61, 0x97, 0x2b, 0xa6,
	0x0d, 0x21, 0xd0, 0x15, 0x69, 0xc9, 0x02, 0x67, 0xec, 0x44, 0x7e, 0x62, 0x63, 0x12, 0x82, 0x9b,
	0xe6, 0x2c, 0xe8, 0x8c, 0x9d, 0xa8, 0xf7, 0xd6, 0xfb, 0xfd, 0xeb, 0xa0, 0xbb, 0xbd, 0x1b, 0xec,
	0x27, 0x58, 0xa4, 0xcf, 0x00, 0xd6, 0xf3, 0xaa, 0xa8, 0x48, 0x00, 0x5b, 0x25, 0xd3, 0x1a, 0xd1,
	0x35, 0x41, 0x93, 0xd2, 0x27, 0x30, 0x98, 0x73, 0x91, 0xb7, 0xd6, 0x28, 0x2e, 0xf2, 0x66, 0x0d,
	0xc6, 0xf4, 0x00, 0xfc, 0xb9, 0x44, 0x08, 0x32, 0x21, 0x40, 0xb6, 0x00, 0x52, 0xe4, 0xf4, 0xbb,
	0x03, 0xfe, 0x67, 0x55, 0xc8, 0x34, 0x4b, 0xd8, 0x25, 0xee, 0x3a, 0x97, 0xc2, 0x30, 0x61, 0x2c,
	0x68, 0x98, 0x34, 0x29, 0xd9, 0x07, 0xff, 0x82, 0x17, 0x6c, 0x61, 0x1f, 0xd2, 0xb1, 0x04, 0x1e,
	0x16, 0x3e, 0xe2, 0x63, 0x9a, 0xa6, 0xe6, 0xb7, 0x2c, 0x70, 0xc7, 0x4e, 0xe4, 0xd6, 0xcd, 0x53,
	0x7e, 0x6b, 0x9b, 0x25, 0x2f, 0xd9, 0xc2, 0x54, 0x8a, 0x05, 0xdd, 0x7a, 0x12, 0x0b, 0x9f, 0x2a,
	0x85, 0x9f, 0xc1, 0xbb, 0xe0, 0x82, 0xeb, 0x25, 0xcb, 0x82, 0xde, 0xd8, 0x89, 0xbc, 0x64, 0x93,
	0x93, 0x1d, 0xe8, 0x69, 0xb5, 0xe0, 0x59, 0xd0, 0xaf, 0xf5, 0x6a, 0x35, 0xcb, 0xe8, 0x53, 0x80,
	0x46, 0xae, 0x56, 0x64, 0x0f, 0xb6, 0xec, 0x62, 0x9e, 0xad, 0x1f, 0xd5, 0xc7, 0x74, 0x96, 0xbd,
	0xf8, 0xe1, 0x42, 0xcf, 0x7e, 0x43, 0x72, 0x02, 0xde, 0x69, 0x5a, 0xd5, 0xf1, 0xce, 0xb4, 0xf6,
	0x40, 0xfb, 0x3a, 0xe1, 0xe8, 0x7e, 0x51, 0x15, 0x15, 0xdd, 0xfd, 0xfa, 0xf3, 0xcf, 0xb7, 0xce,
	0x7f, 0xd4, 0x8f, 0xaf, 0x9f, 0xc7, 0xb6, 0x7b, 0xe4, 0x4c, 0xc8, 0x1b, 0x70, 0x3f, 0x30, 0xf3,
	0xcf, 0x24, 0x23, 0x4b, 0x32, 0x20, 0x77, 0x24, 0xe4, 0x18, 0xdc, 0xf9, 0xca, 0x3c, 0x54, 0x46,
	0x78, 0x5f, 0xc6, 0x31, 0xf4, 0xdf, 0xb1, 0x82, 0x19, 0xf6, 0x50, 0x25, 0x93, 0x96, 0x92, 0x19,
	0x78, 0x68, 0x1d, 0xf4, 0x06, 0x21, 0xeb, 0x89, 0x96, 0x97, 0xc2, 0xed, 0xa6, 0xd6, 0x98, 0x87,
	0xee, 0x59, 0x92, 0x11, 0x1d, 0x22, 0x09, 0x7a, 0x0b, 0xed, 0x83, 0x7a, 0xde, 0x43, 0xbf, 0xbe,
	0x08, 0x69, 0x86, 0x36, 0x7e, 0xda, 0x88, 0xb9, 0x3b, 0x19, 0x7d, 0x64, 0x79, 0xfe, 0xa7, 0x80,
	0x3c, 0x2b, 0x5b, 0x3f, 0x72, 0x26, 0x91, 0x73, 0xd6, 0xb7, 0x3f, 0xcf, 0xcb, 0xbf, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xd6, 0x1b, 0xe3, 0x6f, 0xa8, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelloClient is the client API for Hello service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	Get(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	Put(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	Delete(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	PingPong(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PongReply, error)
	Upload(ctx context.Context, opts ...grpc.CallOption) (Hello_UploadClient, error)
}

type helloClient struct {
	cc *grpc.ClientConn
}

func NewHelloClient(cc *grpc.ClientConn) HelloClient {
	return &helloClient{cc}
}

func (c *helloClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/hello.Hello/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloClient) Get(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/hello.Hello/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloClient) Put(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/hello.Hello/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloClient) Delete(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/hello.Hello/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloClient) PingPong(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PongReply, error) {
	out := new(PongReply)
	err := c.cc.Invoke(ctx, "/hello.Hello/PingPong", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloClient) Upload(ctx context.Context, opts ...grpc.CallOption) (Hello_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Hello_serviceDesc.Streams[0], "/hello.Hello/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloUploadClient{stream}
	return x, nil
}

type Hello_UploadClient interface {
	Send(*UploadReq) error
	CloseAndRecv() (*UploadResp, error)
	grpc.ClientStream
}

type helloUploadClient struct {
	grpc.ClientStream
}

func (x *helloUploadClient) Send(m *UploadReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *helloUploadClient) CloseAndRecv() (*UploadResp, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HelloServer is the server API for Hello service.
type HelloServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	Get(context.Context, *HelloRequest) (*HelloReply, error)
	Put(context.Context, *HelloRequest) (*HelloReply, error)
	Delete(context.Context, *HelloRequest) (*HelloReply, error)
	PingPong(context.Context, *PingRequest) (*PongReply, error)
	Upload(Hello_UploadServer) error
}

// UnimplementedHelloServer can be embedded to have forward compatible implementations.
type UnimplementedHelloServer struct {
}

func (*UnimplementedHelloServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (*UnimplementedHelloServer) Get(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedHelloServer) Put(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (*UnimplementedHelloServer) Delete(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedHelloServer) PingPong(ctx context.Context, req *PingRequest) (*PongReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PingPong not implemented")
}
func (*UnimplementedHelloServer) Upload(srv Hello_UploadServer) error {
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}

func RegisterHelloServer(s *grpc.Server, srv HelloServer) {
	s.RegisterService(&_Hello_serviceDesc, srv)
}

func _Hello_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.Hello/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hello_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.Hello/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServer).Get(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hello_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.Hello/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServer).Put(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hello_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.Hello/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServer).Delete(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hello_PingPong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServer).PingPong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.Hello/PingPong",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServer).PingPong(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hello_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloServer).Upload(&helloUploadServer{stream})
}

type Hello_UploadServer interface {
	SendAndClose(*UploadResp) error
	Recv() (*UploadReq, error)
	grpc.ServerStream
}

type helloUploadServer struct {
	grpc.ServerStream
}

func (x *helloUploadServer) SendAndClose(m *UploadResp) error {
	return x.ServerStream.SendMsg(m)
}

func (x *helloUploadServer) Recv() (*UploadReq, error) {
	m := new(UploadReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Hello_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hello.Hello",
	HandlerType: (*HelloServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Hello_SayHello_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Hello_Get_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _Hello_Put_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Hello_Delete_Handler,
		},
		{
			MethodName: "PingPong",
			Handler:    _Hello_PingPong_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _Hello_Upload_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "hello.proto",
}
