// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sample.proto

package sample

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

type Request struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	// @inject_tag: form:"age" binding:"required,gt=20,lt=27"
	Age                  int32    `protobuf:"varint,2,opt,name=age,proto3" json:"age" form:"age" binding:"required,gt=20,lt=27"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_2141552de9bf11d0, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Request) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

type Reply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Reply) Reset()         { *m = Reply{} }
func (m *Reply) String() string { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()    {}
func (*Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_2141552de9bf11d0, []int{1}
}

func (m *Reply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Reply.Unmarshal(m, b)
}
func (m *Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Reply.Marshal(b, m, deterministic)
}
func (m *Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reply.Merge(m, src)
}
func (m *Reply) XXX_Size() int {
	return xxx_messageInfo_Reply.Size(m)
}
func (m *Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_Reply.DiscardUnknown(m)
}

var xxx_messageInfo_Reply proto.InternalMessageInfo

func (m *Reply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type PingRequest struct {
	Ping                 string   `protobuf:"bytes,1,opt,name=ping,proto3" json:"ping"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2141552de9bf11d0, []int{2}
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
	Pong                 string   `protobuf:"bytes,1,opt,name=pong,proto3" json:"pong"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PongReply) Reset()         { *m = PongReply{} }
func (m *PongReply) String() string { return proto.CompactTextString(m) }
func (*PongReply) ProtoMessage()    {}
func (*PongReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_2141552de9bf11d0, []int{3}
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
	Content  []byte `protobuf:"bytes,1,opt,name=content,proto3" json:"content"`
	FileName string `protobuf:"bytes,2,opt,name=file_name,json=fileName,proto3" json:"file_name"`
	// 文件大小
	FileSize int64  `protobuf:"varint,3,opt,name=file_size,json=fileSize,proto3" json:"file_size"`
	MimeType string `protobuf:"bytes,4,opt,name=mime_type,json=mimeType,proto3" json:"mime_type"`
	// 上传完成
	Finished             bool     `protobuf:"varint,5,opt,name=finished,proto3" json:"finished"`
	SpId                 string   `protobuf:"bytes,6,opt,name=sp_id,json=spId,proto3" json:"sp_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadReq) Reset()         { *m = UploadReq{} }
func (m *UploadReq) String() string { return proto.CompactTextString(m) }
func (*UploadReq) ProtoMessage()    {}
func (*UploadReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_2141552de9bf11d0, []int{4}
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
	FileId               string   `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadResp) Reset()         { *m = UploadResp{} }
func (m *UploadResp) String() string { return proto.CompactTextString(m) }
func (*UploadResp) ProtoMessage()    {}
func (*UploadResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_2141552de9bf11d0, []int{5}
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
	proto.RegisterType((*Request)(nil), "sample.Request")
	proto.RegisterType((*Reply)(nil), "sample.Reply")
	proto.RegisterType((*PingRequest)(nil), "sample.PingRequest")
	proto.RegisterType((*PongReply)(nil), "sample.PongReply")
	proto.RegisterType((*UploadReq)(nil), "sample.UploadReq")
	proto.RegisterType((*UploadResp)(nil), "sample.UploadResp")
}

func init() { proto.RegisterFile("sample.proto", fileDescriptor_2141552de9bf11d0) }

var fileDescriptor_2141552de9bf11d0 = []byte{
	// 495 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x52, 0xcb, 0x6e, 0x13, 0x31,
	0x14, 0xc5, 0x79, 0x4c, 0x26, 0x26, 0x05, 0xea, 0x16, 0x75, 0x34, 0x45, 0x6a, 0x6a, 0x09, 0x29,
	0xaa, 0x48, 0x86, 0xb6, 0x12, 0x12, 0xac, 0xa0, 0x2b, 0x2a, 0x24, 0x54, 0x4d, 0x60, 0x1d, 0xb9,
	0x19, 0xd7, 0xb1, 0x98, 0xb1, 0xdd, 0xd8, 0x6d, 0x94, 0x2e, 0xf9, 0x04, 0xf8, 0x0e, 0xfe, 0x84,
	0x1d, 0x1f, 0x80, 0x84, 0xf8, 0x10, 0x74, 0x3d, 0x0f, 0x22, 0x58, 0xcd, 0x7d, 0x9c, 0x73, 0xae,
	0xef, 0x99, 0x8b, 0x07, 0x96, 0x15, 0x26, 0xe7, 0x13, 0xb3, 0xd4, 0x4e, 0x93, 0xa0, 0xcc, 0xe2,
	0x27, 0x42, 0x6b, 0x91, 0xf3, 0x84, 0x19, 0x99, 0x30, 0xa5, 0xb4, 0x63, 0x4e, 0x6a, 0x65, 0x4b,
	0x54, 0xfc, 0x42, 0x48, 0xb7, 0xb8, 0xb9, 0x9c, 0xcc, 0x75, 0x91, 0x14, 0x2b, 0xe9, 0x3e, 0xe9,
	0x55, 0x22, 0xf4, 0xd8, 0x37, 0xc7, 0xb7, 0x2c, 0x97, 0x19, 0x73, 0x7a, 0x69, 0x93, 0x26, 0xac,
	0x78, 0xcf, 0xfc, 0x67, 0x3e, 0x16, 0x5c, 0x8d, 0xed, 0x8a, 0x09, 0xc1, 0x97, 0x89, 0x36, 0x5e,
	0xf9, 0xff, 0x29, 0xf4, 0x25, 0xee, 0xa5, 0xfc, 0xfa, 0x86, 0x5b, 0x47, 0x08, 0xee, 0x28, 0x56,
	0xf0, 0x08, 0x0d, 0xd1, 0xa8, 0x9f, 0xfa, 0x98, 0xc4, 0xb8, 0xcd, 0x04, 0x8f, 0x5a, 0x43, 0x34,
	0xea, 0x9e, 0x85, 0xbf, 0x7e, 0x1e, 0x74, 0x1e, 0xed, 0x46, 0xfb, 0x29, 0x14, 0xe9, 0x21, 0xee,
	0xa6, 0xdc, 0xe4, 0x6b, 0x12, 0xe1, 0x5e, 0xc1, 0xad, 0x05, 0x60, 0xc9, 0xad, 0x53, 0x7a, 0x88,
	0xef, 0x5f, 0x48, 0x25, 0x36, 0x26, 0x18, 0xa9, 0x44, 0x3d, 0x01, 0x62, 0x7a, 0x80, 0xfb, 0x17,
	0x1a, 0x20, 0xa0, 0x04, 0x00, 0xbd, 0x01, 0xd0, 0x4a, 0xd0, 0x6f, 0x08, 0xf7, 0x3f, 0x9a, 0x5c,
	0xb3, 0x2c, 0xe5, 0xd7, 0x30, 0x6b, 0xae, 0x95, 0xe3, 0xca, 0x79, 0xd0, 0x20, 0xad, 0x53, 0xb2,
	0x8f, 0xfb, 0x57, 0x32, 0xe7, 0x33, 0xbf, 0x43, 0xcb, 0x0b, 0x84, 0x50, 0x78, 0x0f, 0x7b, 0xd4,
	0x4d, 0x2b, 0xef, 0x78, 0xd4, 0x1e, 0xa2, 0x51, 0xbb, 0x6c, 0x4e, 0xe5, 0x9d, 0x6f, 0x16, 0xb2,
	0xe0, 0x33, 0xb7, 0x36, 0x3c, 0xea, 0x94, 0x4c, 0x28, 0x7c, 0x58, 0x1b, 0x70, 0x20, 0xbc, 0x92,
	0x4a, 0xda, 0x05, 0xcf, 0xa2, 0xee, 0x10, 0x8d, 0xc2, 0xb4, 0xc9, 0xc9, 0x0e, 0xee, 0x5a, 0x33,
	0x93, 0x59, 0x14, 0x94, 0xef, 0xb5, 0xe6, 0x3c, 0xa3, 0x4f, 0x31, 0xae, 0x9f, 0x6b, 0x0d, 0xd9,
	0xc3, 0x3d, 0x3f, 0x58, 0x66, 0xd5, 0x52, 0x01, 0xa4, 0xe7, 0xd9, 0xc9, 0x77, 0x84, 0x83, 0xa9,
	0xbf, 0x03, 0xf2, 0x1a, 0x87, 0x53, 0xb6, 0x7e, 0xcb, 0xf3, 0x5c, 0x93, 0x87, 0x93, 0xea, 0x54,
	0x2a, 0xcf, 0xe2, 0xad, 0xbf, 0x05, 0x93, 0xaf, 0xe9, 0xee, 0xe7, 0x1f, 0xbf, 0xbf, 0xb6, 0x1e,
	0xd0, 0x7e, 0x72, 0x7b, 0x9c, 0x2c, 0x80, 0xf2, 0x0a, 0x1d, 0x91, 0x77, 0x38, 0x04, 0x9f, 0xc1,
	0x48, 0xb2, 0x53, 0x13, 0x36, 0x9c, 0x8f, 0xb7, 0x9b, 0x62, 0xed, 0x35, 0xdd, 0xf3, 0x4a, 0xdb,
	0x74, 0x00, 0x4a, 0xf0, 0x2b, 0xc0, 0x6d, 0x10, 0x3b, 0xc5, 0x41, 0xb9, 0x00, 0x69, 0x58, 0x8d,
	0xff, 0x31, 0xf9, 0xb7, 0x64, 0x0d, 0xbd, 0x37, 0x42, 0x67, 0xd1, 0x97, 0x37, 0x8f, 0xc9, 0x16,
	0xae, 0x4e, 0xfb, 0xa4, 0x7d, 0x3c, 0x79, 0x4e, 0x51, 0x72, 0x84, 0xd0, 0x65, 0xe0, 0x0f, 0xed,
	0xf4, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc8, 0x13, 0x74, 0x33, 0x04, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SampleClient is the client API for Sample service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SampleClient interface {
	SayHello(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error)
	PingPong(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PongReply, error)
	Upload(ctx context.Context, opts ...grpc.CallOption) (Sample_UploadClient, error)
}

type sampleClient struct {
	cc *grpc.ClientConn
}

func NewSampleClient(cc *grpc.ClientConn) SampleClient {
	return &sampleClient{cc}
}

func (c *sampleClient) SayHello(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/sample.Sample/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sampleClient) PingPong(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PongReply, error) {
	out := new(PongReply)
	err := c.cc.Invoke(ctx, "/sample.Sample/PingPong", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sampleClient) Upload(ctx context.Context, opts ...grpc.CallOption) (Sample_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Sample_serviceDesc.Streams[0], "/sample.Sample/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &sampleUploadClient{stream}
	return x, nil
}

type Sample_UploadClient interface {
	Send(*UploadReq) error
	CloseAndRecv() (*UploadResp, error)
	grpc.ClientStream
}

type sampleUploadClient struct {
	grpc.ClientStream
}

func (x *sampleUploadClient) Send(m *UploadReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sampleUploadClient) CloseAndRecv() (*UploadResp, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SampleServer is the server API for Sample service.
type SampleServer interface {
	SayHello(context.Context, *Request) (*Reply, error)
	PingPong(context.Context, *PingRequest) (*PongReply, error)
	Upload(Sample_UploadServer) error
}

// UnimplementedSampleServer can be embedded to have forward compatible implementations.
type UnimplementedSampleServer struct {
}

func (*UnimplementedSampleServer) SayHello(ctx context.Context, req *Request) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (*UnimplementedSampleServer) PingPong(ctx context.Context, req *PingRequest) (*PongReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PingPong not implemented")
}
func (*UnimplementedSampleServer) Upload(srv Sample_UploadServer) error {
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}

func RegisterSampleServer(s *grpc.Server, srv SampleServer) {
	s.RegisterService(&_Sample_serviceDesc, srv)
}

func _Sample_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SampleServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sample.Sample/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SampleServer).SayHello(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sample_PingPong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SampleServer).PingPong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sample.Sample/PingPong",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SampleServer).PingPong(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sample_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SampleServer).Upload(&sampleUploadServer{stream})
}

type Sample_UploadServer interface {
	SendAndClose(*UploadResp) error
	Recv() (*UploadReq, error)
	grpc.ServerStream
}

type sampleUploadServer struct {
	grpc.ServerStream
}

func (x *sampleUploadServer) SendAndClose(m *UploadResp) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sampleUploadServer) Recv() (*UploadReq, error) {
	m := new(UploadReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Sample_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sample.Sample",
	HandlerType: (*SampleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Sample_SayHello_Handler,
		},
		{
			MethodName: "PingPong",
			Handler:    _Sample_PingPong_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _Sample_Upload_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "sample.proto",
}
