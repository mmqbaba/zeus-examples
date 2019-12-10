// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hellodemo.proto

package hellodemo

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
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
	// @inject_tag: form:"age" binding:"required,gt=20,lt=27"
	Age                  int32    `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty" form:"age" binding:"required,gt=20,lt=27"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec69e4919e0e086c, []int{0}
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
	return fileDescriptor_ec69e4919e0e086c, []int{1}
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
	return fileDescriptor_ec69e4919e0e086c, []int{2}
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
	return fileDescriptor_ec69e4919e0e086c, []int{3}
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
	return fileDescriptor_ec69e4919e0e086c, []int{4}
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
	return fileDescriptor_ec69e4919e0e086c, []int{5}
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
	proto.RegisterType((*HelloRequest)(nil), "hellodemo.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "hellodemo.HelloReply")
	proto.RegisterType((*PingRequest)(nil), "hellodemo.PingRequest")
	proto.RegisterType((*PongReply)(nil), "hellodemo.PongReply")
	proto.RegisterType((*UploadReq)(nil), "hellodemo.UploadReq")
	proto.RegisterType((*UploadResp)(nil), "hellodemo.UploadResp")
}

func init() { proto.RegisterFile("hellodemo.proto", fileDescriptor_ec69e4919e0e086c) }

var fileDescriptor_ec69e4919e0e086c = []byte{
	// 546 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xc1, 0x6e, 0xd3, 0x4c,
	0x10, 0xfe, 0x5d, 0x37, 0xa9, 0x3d, 0x8d, 0x7e, 0xda, 0x25, 0x10, 0xcb, 0x41, 0x6a, 0x58, 0x09,
	0x14, 0x45, 0x24, 0xa6, 0x45, 0x42, 0x82, 0x03, 0x12, 0x55, 0x25, 0x5a, 0x55, 0x42, 0xc1, 0x85,
	0x73, 0xe4, 0xc6, 0x53, 0x67, 0x85, 0xed, 0xdd, 0x66, 0x37, 0x8d, 0xd2, 0x23, 0x8f, 0x00, 0xcf,
	0xc1, 0xb3, 0x70, 0xe0, 0x01, 0x90, 0x10, 0x0f, 0x82, 0x76, 0x1d, 0x27, 0xae, 0x7a, 0xa2, 0x27,
	0xef, 0xcc, 0x7c, 0xf3, 0xcd, 0xb7, 0x9e, 0x6f, 0xe1, 0xde, 0x04, 0xd3, 0x94, 0xc7, 0x98, 0xf1,
	0x81, 0x98, 0x72, 0xc5, 0x89, 0xbb, 0x4a, 0xf8, 0x8f, 0x12, 0xce, 0x93, 0x14, 0x83, 0x48, 0xb0,
	0x20, 0xca, 0x73, 0xae, 0x22, 0xc5, 0x78, 0x2e, 0x0b, 0xa0, 0xff, 0x32, 0x61, 0x6a, 0x32, 0x3b,
	0x1f, 0x8c, 0x79, 0x16, 0x64, 0x73, 0xa6, 0x3e, 0xf3, 0x79, 0x90, 0xf0, 0xbe, 0x29, 0xf6, 0xaf,
	0xa2, 0x94, 0xc5, 0x91, 0xe2, 0x53, 0x19, 0xac, 0x8e, 0xcb, 0xbe, 0x67, 0xe6, 0x33, 0xee, 0x27,
	0x98, 0xf7, 0xe5, 0x3c, 0x4a, 0x12, 0x9c, 0x06, 0x5c, 0x18, 0xe6, 0xdb, 0x53, 0xe8, 0x1b, 0x68,
	0x1c, 0x6b, 0x41, 0x21, 0x5e, 0xce, 0x50, 0x2a, 0x42, 0x60, 0x33, 0x8f, 0x32, 0xf4, 0xac, 0x8e,
	0xd5, 0x75, 0x43, 0x73, 0x26, 0x3e, 0xd8, 0x51, 0x82, 0xde, 0x46, 0xc7, 0xea, 0xd6, 0x0e, 0x9d,
	0xdf, 0xbf, 0xf6, 0x36, 0x77, 0x9a, 0x5e, 0x3b, 0xd4, 0x49, 0xfa, 0x14, 0x60, 0xd9, 0x2f, 0xd2,
	0x05, 0xf1, 0x60, 0x2b, 0x43, 0x29, 0x35, 0xba, 0x20, 0x28, 0x43, 0xfa, 0x18, 0xb6, 0x87, 0x2c,
	0x4f, 0x2a, 0x63, 0x04, 0xcb, 0x93, 0x72, 0x8c, 0x3e, 0xd3, 0x3d, 0x70, 0x87, 0x5c, 0x43, 0x34,
	0x93, 0x06, 0xf0, 0x0a, 0x80, 0xe7, 0x09, 0xfd, 0x6e, 0x81, 0xfb, 0x49, 0xa4, 0x3c, 0x8a, 0x43,
	0xbc, 0xd4, 0xb3, 0xc6, 0x3c, 0x57, 0x98, 0x2b, 0x03, 0x6a, 0x84, 0x65, 0x48, 0xda, 0xe0, 0x5e,
	0xb0, 0x14, 0x47, 0xe6, 0x22, 0x1b, 0x86, 0xc0, 0xd1, 0x89, 0xf7, 0xfa, 0x32, 0x65, 0x51, 0xb2,
	0x6b, 0xf4, 0xec, 0x8e, 0xd5, 0xb5, 0x8b, 0xe2, 0x19, 0xbb, 0x36, 0xc5, 0x8c, 0x65, 0x38, 0x52,
	0x0b, 0x81, 0xde, 0x66, 0xd1, 0xa9, 0x13, 0x1f, 0x17, 0x42, 0xff, 0x06, 0xe7, 0x82, 0xe5, 0x4c,
	0x4e, 0x30, 0xf6, 0x6a, 0x1d, 0xab, 0xeb, 0x84, 0xab, 0x98, 0xdc, 0x87, 0x9a, 0x14, 0x23, 0x16,
	0x7b, 0xf5, 0x42, 0xaf, 0x14, 0x27, 0x31, 0x7d, 0x02, 0x50, 0xca, 0x95, 0x82, 0xb4, 0x60, 0xcb,
	0x0c, 0x66, 0xf1, 0xf2, 0x52, 0x75, 0x1d, 0x9e, 0xc4, 0x07, 0x3f, 0x6c, 0x70, 0xcd, 0x3f, 0x3c,
	0xc2, 0x8c, 0x93, 0x21, 0x38, 0x67, 0xd1, 0xc2, 0xc4, 0xa4, 0x35, 0x58, 0xbb, 0xa7, 0xba, 0x25,
	0xff, 0xc1, 0xed, 0x82, 0x48, 0x17, 0xb4, 0xf9, 0xe5, 0xe7, 0x9f, 0x6f, 0x1b, 0xff, 0x53, 0x37,
	0xb8, 0xda, 0x0f, 0x0c, 0xe2, 0xb5, 0xd5, 0x23, 0xc7, 0x60, 0xbf, 0x43, 0xf5, 0xcf, 0x64, 0xbb,
	0x86, 0x6c, 0x9b, 0xac, 0xc9, 0xc8, 0x29, 0xd8, 0xc3, 0x99, 0xba, 0xab, 0x2c, 0xff, 0xa6, 0xac,
	0x53, 0xa8, 0x1f, 0x61, 0x8a, 0x0a, 0xef, 0xaa, 0xac, 0x57, 0x51, 0xf6, 0x01, 0x1c, 0x6d, 0x2f,
	0xed, 0x1f, 0xf2, 0xb0, 0xd2, 0x55, 0xf1, 0x9c, 0xdf, 0xac, 0xe6, 0x4b, 0xa3, 0xd1, 0x96, 0x21,
	0xdb, 0xa5, 0x0d, 0x4d, 0xa6, 0x7d, 0xa8, 0xad, 0xa6, 0xf5, 0xbd, 0x82, 0x7a, 0xb1, 0x3d, 0x52,
	0x6d, 0x5c, 0xf9, 0xef, 0x86, 0xb8, 0xf5, 0x9a, 0xe9, 0x7f, 0x5d, 0xeb, 0xb0, 0xfd, 0xf5, 0xad,
	0x47, 0x76, 0x60, 0xfd, 0xd4, 0x0f, 0xec, 0xfd, 0xc1, 0x73, 0x6a, 0x05, 0x3d, 0xcb, 0x3a, 0xaf,
	0x9b, 0x87, 0xf7, 0xe2, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0a, 0xcf, 0x1e, 0x9d, 0x1a, 0x04,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelloDemoClient is the client API for HelloDemo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloDemoClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	Get(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	Put(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	Delete(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	PingPong(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PongReply, error)
	Upload(ctx context.Context, opts ...grpc.CallOption) (HelloDemo_UploadClient, error)
}

type helloDemoClient struct {
	cc *grpc.ClientConn
}

func NewHelloDemoClient(cc *grpc.ClientConn) HelloDemoClient {
	return &helloDemoClient{cc}
}

func (c *helloDemoClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/hellodemo.HelloDemo/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloDemoClient) Get(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/hellodemo.HelloDemo/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloDemoClient) Put(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/hellodemo.HelloDemo/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloDemoClient) Delete(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/hellodemo.HelloDemo/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloDemoClient) PingPong(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PongReply, error) {
	out := new(PongReply)
	err := c.cc.Invoke(ctx, "/hellodemo.HelloDemo/PingPong", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloDemoClient) Upload(ctx context.Context, opts ...grpc.CallOption) (HelloDemo_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_HelloDemo_serviceDesc.Streams[0], "/hellodemo.HelloDemo/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloDemoUploadClient{stream}
	return x, nil
}

type HelloDemo_UploadClient interface {
	Send(*UploadReq) error
	CloseAndRecv() (*UploadResp, error)
	grpc.ClientStream
}

type helloDemoUploadClient struct {
	grpc.ClientStream
}

func (x *helloDemoUploadClient) Send(m *UploadReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *helloDemoUploadClient) CloseAndRecv() (*UploadResp, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HelloDemoServer is the server API for HelloDemo service.
type HelloDemoServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	Get(context.Context, *HelloRequest) (*HelloReply, error)
	Put(context.Context, *HelloRequest) (*HelloReply, error)
	Delete(context.Context, *HelloRequest) (*HelloReply, error)
	PingPong(context.Context, *PingRequest) (*PongReply, error)
	Upload(HelloDemo_UploadServer) error
}

// UnimplementedHelloDemoServer can be embedded to have forward compatible implementations.
type UnimplementedHelloDemoServer struct {
}

func (*UnimplementedHelloDemoServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (*UnimplementedHelloDemoServer) Get(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedHelloDemoServer) Put(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (*UnimplementedHelloDemoServer) Delete(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedHelloDemoServer) PingPong(ctx context.Context, req *PingRequest) (*PongReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PingPong not implemented")
}
func (*UnimplementedHelloDemoServer) Upload(srv HelloDemo_UploadServer) error {
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}

func RegisterHelloDemoServer(s *grpc.Server, srv HelloDemoServer) {
	s.RegisterService(&_HelloDemo_serviceDesc, srv)
}

func _HelloDemo_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloDemoServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hellodemo.HelloDemo/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloDemoServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloDemo_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloDemoServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hellodemo.HelloDemo/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloDemoServer).Get(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloDemo_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloDemoServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hellodemo.HelloDemo/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloDemoServer).Put(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloDemo_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloDemoServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hellodemo.HelloDemo/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloDemoServer).Delete(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloDemo_PingPong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloDemoServer).PingPong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hellodemo.HelloDemo/PingPong",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloDemoServer).PingPong(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloDemo_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloDemoServer).Upload(&helloDemoUploadServer{stream})
}

type HelloDemo_UploadServer interface {
	SendAndClose(*UploadResp) error
	Recv() (*UploadReq, error)
	grpc.ServerStream
}

type helloDemoUploadServer struct {
	grpc.ServerStream
}

func (x *helloDemoUploadServer) SendAndClose(m *UploadResp) error {
	return x.ServerStream.SendMsg(m)
}

func (x *helloDemoUploadServer) Recv() (*UploadReq, error) {
	m := new(UploadReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _HelloDemo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hellodemo.HelloDemo",
	HandlerType: (*HelloDemoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _HelloDemo_SayHello_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _HelloDemo_Get_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _HelloDemo_Put_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _HelloDemo_Delete_Handler,
		},
		{
			MethodName: "PingPong",
			Handler:    _HelloDemo_PingPong_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _HelloDemo_Upload_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "hellodemo.proto",
}
