// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sample.proto

package sample

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// @inject_tag: form:"age" binding:"required,gt=20,lt=27"
	Age                  int32             `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty" form:"age" binding:"required,gt=20,lt=27"`
	MetaData             map[string]string `protobuf:"bytes,3,rep,name=meta_data,json=metaData,proto3" json:"meta_data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
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

func (m *Request) GetMetaData() map[string]string {
	if m != nil {
		return m.MetaData
	}
	return nil
}

type Reply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
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
	Ping                 string   `protobuf:"bytes,1,opt,name=ping,proto3" json:"ping,omitempty"`
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
	Pong                 string   `protobuf:"bytes,1,opt,name=pong,proto3" json:"pong,omitempty"`
	Data                 *any.Any `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
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

func (m *PongReply) GetData() *any.Any {
	if m != nil {
		return m.Data
	}
	return nil
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
	FileId               string   `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
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
	proto.RegisterMapType((map[string]string)(nil), "sample.Request.MetaDataEntry")
	proto.RegisterType((*Reply)(nil), "sample.Reply")
	proto.RegisterType((*PingRequest)(nil), "sample.PingRequest")
	proto.RegisterType((*PongReply)(nil), "sample.PongReply")
	proto.RegisterType((*UploadReq)(nil), "sample.UploadReq")
	proto.RegisterType((*UploadResp)(nil), "sample.UploadResp")
}

func init() { proto.RegisterFile("sample.proto", fileDescriptor_2141552de9bf11d0) }

var fileDescriptor_2141552de9bf11d0 = []byte{
	// 595 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x53, 0xd1, 0x4e, 0xd4, 0x40,
	0x14, 0x75, 0xe8, 0xee, 0xd2, 0x5e, 0x40, 0x61, 0xd8, 0x84, 0x5a, 0x34, 0x2e, 0x4d, 0x4c, 0x1a,
	0xe2, 0xb6, 0xb2, 0x24, 0xc6, 0xe0, 0x8b, 0x10, 0x4d, 0x24, 0x46, 0x43, 0x8a, 0x3e, 0x6f, 0x06,
	0x3a, 0x94, 0x09, 0xed, 0xcc, 0xb0, 0x33, 0x0b, 0x29, 0x8f, 0x7e, 0x82, 0x7e, 0x87, 0x2f, 0x7e,
	0x87, 0x6f, 0x7e, 0x80, 0x89, 0xf1, 0x43, 0xcc, 0x4c, 0xdb, 0x75, 0xc5, 0xa7, 0xde, 0x7b, 0xcf,
	0xbd, 0xe7, 0xce, 0xb9, 0x27, 0x85, 0x65, 0x45, 0x4a, 0x59, 0xd0, 0x58, 0x4e, 0x84, 0x16, 0xb8,
	0x57, 0x67, 0xc1, 0xfd, 0x5c, 0x88, 0xbc, 0xa0, 0x89, 0xad, 0x9e, 0x4c, 0xcf, 0x12, 0xc2, 0xab,
	0xba, 0x25, 0x78, 0xd0, 0x40, 0x44, 0xb2, 0x84, 0x70, 0x2e, 0x34, 0xd1, 0x4c, 0x70, 0xd5, 0xa0,
	0xcf, 0x72, 0xa6, 0xcf, 0xa7, 0x27, 0xf1, 0xa9, 0x28, 0x93, 0xf2, 0x9a, 0xe9, 0x0b, 0x71, 0x9d,
	0xe4, 0x62, 0x68, 0xc1, 0xe1, 0x15, 0x29, 0x58, 0x46, 0xb4, 0x98, 0xa8, 0x64, 0x16, 0x36, 0x73,
	0x4f, 0xec, 0xe7, 0x74, 0x98, 0x53, 0x3e, 0x54, 0xd7, 0x24, 0xcf, 0xe9, 0x24, 0x11, 0xd2, 0x32,
	0xff, 0xbf, 0x25, 0xfc, 0x86, 0x60, 0x31, 0xa5, 0x97, 0x53, 0xaa, 0x34, 0xc6, 0xd0, 0xe1, 0xa4,
	0xa4, 0x3e, 0x1a, 0xa0, 0xc8, 0x4b, 0x6d, 0x8c, 0x03, 0x70, 0x48, 0x4e, 0xfd, 0x85, 0x01, 0x8a,
	0xba, 0x07, 0xee, 0xaf, 0x9f, 0x8f, 0x3a, 0xab, 0x7d, 0x7f, 0x33, 0x35, 0x45, 0xbc, 0x07, 0x5e,
	0x49, 0x35, 0x19, 0x67, 0x44, 0x13, 0xdf, 0x19, 0x38, 0xd1, 0xd2, 0xe8, 0x61, 0xdc, 0x1c, 0xa1,
	0xe1, 0x8c, 0xdf, 0x51, 0x4d, 0x5e, 0x11, 0x4d, 0x5e, 0x73, 0x3d, 0xa9, 0x52, 0xb7, 0x6c, 0xd2,
	0xe0, 0x05, 0xac, 0xfc, 0x03, 0xe1, 0x55, 0x70, 0x2e, 0x68, 0xd5, 0xec, 0x36, 0x21, 0xee, 0x43,
	0xf7, 0x8a, 0x14, 0xd3, 0x7a, 0xb9, 0x97, 0xd6, 0xc9, 0xde, 0xc2, 0x73, 0x14, 0x6e, 0x41, 0x37,
	0xa5, 0xb2, 0xa8, 0xb0, 0x0f, 0x8b, 0x25, 0x55, 0xca, 0xbc, 0xb0, 0x1e, 0x6c, 0xd3, 0x70, 0x0b,
	0x96, 0x8e, 0x18, 0xcf, 0xe7, 0xa4, 0x49, 0xc6, 0xf3, 0x56, 0x9a, 0x89, 0xc3, 0x43, 0xf0, 0x8e,
	0x84, 0x69, 0x31, 0x4c, 0xa6, 0x41, 0xcc, 0x35, 0x08, 0x9e, 0xe3, 0x08, 0x3a, 0x56, 0x9a, 0xd9,
	0xbf, 0x34, 0xea, 0xc7, 0xb5, 0x5d, 0x71, 0xeb, 0x64, 0xbc, 0xcf, 0xab, 0xd4, 0x76, 0x84, 0x5f,
	0x11, 0x78, 0x1f, 0x65, 0x21, 0x48, 0x96, 0xd2, 0x4b, 0xf3, 0xaa, 0x53, 0xc1, 0x35, 0xe5, 0xda,
	0xd2, 0x2d, 0xa7, 0x6d, 0x8a, 0x37, 0xc1, 0x3b, 0x63, 0x05, 0x1d, 0xdb, 0x33, 0xd7, 0xb2, 0x5c,
	0x53, 0x78, 0x6f, 0x4e, 0xdd, 0x82, 0x8a, 0xdd, 0x50, 0xdf, 0x19, 0xa0, 0xc8, 0xa9, 0xc1, 0x63,
	0x76, 0x63, 0xc1, 0x92, 0x95, 0x74, 0xac, 0x2b, 0x49, 0xfd, 0x4e, 0x3d, 0x69, 0x0a, 0x1f, 0x2a,
	0x69, 0x4c, 0x72, 0xcf, 0x18, 0x67, 0xea, 0x9c, 0x66, 0x7e, 0x77, 0x80, 0x22, 0x37, 0x9d, 0xe5,
	0x78, 0x1d, 0xba, 0x4a, 0x8e, 0x59, 0xe6, 0xf7, 0x6a, 0x65, 0x4a, 0x1e, 0x66, 0xe1, 0x63, 0x80,
	0xf6, 0xb9, 0x4a, 0xe2, 0x0d, 0x58, 0xb4, 0x8b, 0x59, 0xd6, 0xc8, 0xef, 0x99, 0xf4, 0x30, 0x1b,
	0x7d, 0x47, 0xd0, 0x3b, 0xb6, 0x7e, 0xe2, 0x97, 0xe0, 0x1e, 0x93, 0xea, 0x0d, 0x2d, 0x0a, 0x81,
	0xef, 0xdd, 0x32, 0x39, 0x58, 0xf9, 0x5b, 0x90, 0x45, 0x15, 0xf6, 0x3f, 0xfd, 0xf8, 0xfd, 0x65,
	0xe1, 0x6e, 0xe8, 0x25, 0x57, 0x3b, 0xc9, 0xb9, 0x19, 0xd9, 0x43, 0xdb, 0xf8, 0x2d, 0xb8, 0xc6,
	0x11, 0x73, 0x72, 0xbc, 0xde, 0x0e, 0xcc, 0x79, 0x14, 0xac, 0xcd, 0x8a, 0xad, 0x2b, 0xe1, 0x86,
	0x65, 0x5a, 0x0b, 0x97, 0x0d, 0x93, 0x31, 0xcd, 0xf8, 0x62, 0xc8, 0x76, 0xa1, 0x57, 0x0b, 0xc0,
	0xb3, 0xa9, 0xd9, 0xfd, 0x03, 0x7c, 0xbb, 0xa4, 0x64, 0x78, 0x27, 0x42, 0x07, 0x9b, 0x9f, 0xf7,
	0x7d, 0xbc, 0x02, 0xcd, 0x9f, 0x39, 0x72, 0x76, 0xe2, 0xa7, 0x61, 0xc7, 0xfc, 0x7d, 0xdb, 0x08,
	0x9d, 0xf4, 0xac, 0xad, 0xbb, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x9e, 0x01, 0x0c, 0xc6,
	0x03, 0x00, 0x00,
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
