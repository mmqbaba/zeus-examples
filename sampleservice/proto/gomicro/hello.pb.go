// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello.proto

package hello

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	// @inject_tag: form:"age" binding:"required,gt=20,lt=27"
	Age                  int32    `protobuf:"varint,2,opt,name=age,proto3" json:"age" form:"age" binding:"required,gt=20,lt=27"`
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
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message"`
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
	Ping                 string   `protobuf:"bytes,1,opt,name=ping,proto3" json:"ping"`
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
	Pong                 string   `protobuf:"bytes,1,opt,name=pong,proto3" json:"pong"`
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
	FileId               string   `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id"`
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
	// 455 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x52, 0xcf, 0x6e, 0x13, 0x3f,
	0x10, 0xfe, 0x6d, 0x92, 0x4d, 0x37, 0xd3, 0xfe, 0xa4, 0xe0, 0x82, 0xb0, 0x36, 0x48, 0x0d, 0x96,
	0x40, 0x2b, 0xa4, 0x66, 0x55, 0x90, 0x38, 0x70, 0xe0, 0xc0, 0x89, 0x4a, 0x80, 0xaa, 0x2d, 0x9c,
	0x23, 0xb7, 0xeb, 0x6e, 0x2c, 0x76, 0x6d, 0xb7, 0x76, 0x5b, 0x6d, 0x8f, 0xbc, 0x02, 0xcf, 0xc1,
	0xd3, 0x70, 0xe0, 0x88, 0x84, 0x78, 0x10, 0x34, 0x76, 0x1c, 0xc2, 0x6d, 0xfe, 0x7c, 0xf3, 0xcd,
	0xcc, 0x37, 0x03, 0xbb, 0x2b, 0xd1, 0xb6, 0x7a, 0x61, 0xae, 0xb4, 0xd3, 0x24, 0xf5, 0x4e, 0xfe,
	0xa8, 0xd1, 0xba, 0x69, 0x45, 0xc9, 0x8d, 0x2c, 0xb9, 0x52, 0xda, 0x71, 0x27, 0xb5, 0xb2, 0x01,
	0x94, 0xbf, 0x6c, 0xa4, 0x5b, 0x5d, 0x9f, 0x2d, 0xce, 0x75, 0x57, 0x76, 0xb7, 0xd2, 0x7d, 0xd6,
	0xb7, 0x65, 0xa3, 0x0f, 0x7d, 0xf2, 0xf0, 0x86, 0xb7, 0xb2, 0xe6, 0x4e, 0x5f, 0xd9, 0x72, 0x63,
	0x86, 0x3a, 0xf6, 0x1a, 0xf6, 0xde, 0x22, 0x7d, 0x25, 0x2e, 0xaf, 0x85, 0x75, 0x84, 0xc0, 0x48,
	0xf1, 0x4e, 0xd0, 0x64, 0x9e, 0x14, 0x93, 0xca, 0xdb, 0x24, 0x87, 0x21, 0x6f, 0x04, 0x1d, 0xcc,
	0x93, 0x22, 0x7d, 0x93, 0xfd, 0xfa, 0x79, 0x30, 0x9a, 0xde, 0xa7, 0xb3, 0x0a, 0x83, 0xec, 0x29,
	0xc0, 0xba, 0xde, 0xb4, 0x3d, 0xa1, 0xb0, 0xd3, 0x09, 0x6b, 0x11, 0x1d, 0x08, 0xa2, 0xcb, 0x1e,
	0xc3, 0xee, 0x89, 0x54, 0xcd, 0x56, 0x1b, 0x23, 0x55, 0x13, 0xdb, 0xa0, 0xcd, 0x0e, 0x60, 0x72,
	0xa2, 0x11, 0x82, 0x4c, 0x08, 0xd0, 0x5b, 0x00, 0xad, 0x1a, 0xf6, 0x2d, 0x81, 0xc9, 0x27, 0xd3,
	0x6a, 0x5e, 0x57, 0xe2, 0x12, 0x7b, 0x9d, 0x6b, 0xe5, 0x84, 0x72, 0x1e, 0xb4, 0x57, 0x45, 0x97,
	0xcc, 0x60, 0x72, 0x21, 0x5b, 0xb1, 0xf4, 0x8b, 0x0c, 0x3c, 0x41, 0x86, 0x81, 0x0f, 0xb8, 0x4c,
	0x4c, 0x5a, 0x79, 0x27, 0xe8, 0x70, 0x9e, 0x14, 0xc3, 0x90, 0x3c, 0x95, 0x77, 0x3e, 0xd9, 0xc9,
	0x4e, 0x2c, 0x5d, 0x6f, 0x04, 0x1d, 0x85, 0x4a, 0x0c, 0x7c, 0xec, 0x0d, 0xca, 0x90, 0x5d, 0x48,
	0x25, 0xed, 0x4a, 0xd4, 0x34, 0x9d, 0x27, 0x45, 0x56, 0x6d, 0x7c, 0xb2, 0x0f, 0xa9, 0x35, 0x4b,
	0x59, 0xd3, 0x71, 0x98, 0xd7, 0x9a, 0xe3, 0x9a, 0x3d, 0x01, 0x88, 0xe3, 0x5a, 0x43, 0x1e, 0xc2,
	0x8e, 0x6f, 0x2c, 0xeb, 0xf5, 0x52, 0x63, 0x74, 0x8f, 0xeb, 0xe7, 0x3f, 0x12, 0x48, 0xbd, 0x86,
	0xe4, 0x1d, 0x64, 0xa7, 0xbc, 0x0f, 0xf6, 0xfe, 0x22, 0xfc, 0xc0, 0xf6, 0x75, 0xf2, 0x7b, 0xff,
	0x06, 0x4d, 0xdb, 0x33, 0xfa, 0xe5, 0xfb, 0xef, 0xaf, 0x03, 0xc2, 0xfe, 0xf7, 0x8f, 0x71, 0x73,
	0x54, 0x7a, 0xc4, 0xab, 0xe4, 0x19, 0x79, 0x0f, 0x19, 0x4a, 0x8e, 0x9a, 0x12, 0xb2, 0x2e, 0xdc,
	0xba, 0x41, 0x3e, 0x8d, 0xb1, 0x28, 0x3a, 0x9b, 0x79, 0xae, 0x07, 0x6c, 0x1a, 0xb9, 0xf0, 0x2e,
	0x28, 0x3d, 0xd2, 0x1d, 0xc1, 0x38, 0x6c, 0x43, 0x62, 0xe1, 0xe6, 0x16, 0x9b, 0xb9, 0xfe, 0xae,
	0xcb, 0xfe, 0x2b, 0x92, 0xb3, 0xb1, 0xff, 0xb1, 0x17, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xe6,
	0x1c, 0xc5, 0x74, 0xcf, 0x02, 0x00, 0x00,
}
