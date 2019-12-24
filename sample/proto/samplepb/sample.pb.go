// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sample.proto

package sample

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	_struct "github.com/golang/protobuf/ptypes/struct"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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
	SexType              bool              `protobuf:"varint,4,opt,name=sex_type,json=sexType,proto3" json:"sex_type,omitempty"`
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

func (m *Request) GetSexType() bool {
	if m != nil {
		return m.SexType
	}
	return false
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
	Ping                 string                `protobuf:"bytes,1,opt,name=ping,proto3" json:"ping,omitempty"`
	StMetaData           *_struct.Struct       `protobuf:"bytes,2,opt,name=st_meta_data,json=stMetaData,proto3" json:"st_meta_data,omitempty"`
	StCustomData         *_struct.Struct       `protobuf:"bytes,3,opt,name=st_custom_data,json=stCustomData,proto3" json:"st_custom_data,omitempty"`
	DoubleValue          *wrappers.DoubleValue `protobuf:"bytes,5,opt,name=double_value,json=doubleValue,proto3" json:"double_value,omitempty"`
	StringValue          *wrappers.StringValue `protobuf:"bytes,6,opt,name=string_value,json=stringValue,proto3" json:"string_value,omitempty"`
	BoolValue            *wrappers.BoolValue   `protobuf:"bytes,7,opt,name=bool_value,json=boolValue,proto3" json:"bool_value,omitempty"`
	Int64Value           *wrappers.Int64Value  `protobuf:"bytes,8,opt,name=int64_value,json=int64Value,proto3" json:"int64_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
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

func (m *PingRequest) GetStMetaData() *_struct.Struct {
	if m != nil {
		return m.StMetaData
	}
	return nil
}

func (m *PingRequest) GetStCustomData() *_struct.Struct {
	if m != nil {
		return m.StCustomData
	}
	return nil
}

func (m *PingRequest) GetDoubleValue() *wrappers.DoubleValue {
	if m != nil {
		return m.DoubleValue
	}
	return nil
}

func (m *PingRequest) GetStringValue() *wrappers.StringValue {
	if m != nil {
		return m.StringValue
	}
	return nil
}

func (m *PingRequest) GetBoolValue() *wrappers.BoolValue {
	if m != nil {
		return m.BoolValue
	}
	return nil
}

func (m *PingRequest) GetInt64Value() *wrappers.Int64Value {
	if m != nil {
		return m.Int64Value
	}
	return nil
}

type PongReply struct {
	Pong                 string          `protobuf:"bytes,1,opt,name=pong,proto3" json:"pong,omitempty"`
	StMetaData           *_struct.Struct `protobuf:"bytes,2,opt,name=st_meta_data,json=stMetaData,proto3" json:"st_meta_data,omitempty"`
	StCustomData         *_struct.Struct `protobuf:"bytes,3,opt,name=st_custom_data,json=stCustomData,proto3" json:"st_custom_data,omitempty"`
	AnyData              *any.Any        `protobuf:"bytes,4,opt,name=any_data,json=anyData,proto3" json:"any_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
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

func (m *PongReply) GetStMetaData() *_struct.Struct {
	if m != nil {
		return m.StMetaData
	}
	return nil
}

func (m *PongReply) GetStCustomData() *_struct.Struct {
	if m != nil {
		return m.StCustomData
	}
	return nil
}

func (m *PongReply) GetAnyData() *any.Any {
	if m != nil {
		return m.AnyData
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

type GetMsgReq struct {
	// @inject_tag: form:"msg"
	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty" form:"msg"`
	// @inject_tag: form:"count"
	Count                int32    `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty" form:"count"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMsgReq) Reset()         { *m = GetMsgReq{} }
func (m *GetMsgReq) String() string { return proto.CompactTextString(m) }
func (*GetMsgReq) ProtoMessage()    {}
func (*GetMsgReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_2141552de9bf11d0, []int{6}
}

func (m *GetMsgReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMsgReq.Unmarshal(m, b)
}
func (m *GetMsgReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMsgReq.Marshal(b, m, deterministic)
}
func (m *GetMsgReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMsgReq.Merge(m, src)
}
func (m *GetMsgReq) XXX_Size() int {
	return xxx_messageInfo_GetMsgReq.Size(m)
}
func (m *GetMsgReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMsgReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetMsgReq proto.InternalMessageInfo

func (m *GetMsgReq) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *GetMsgReq) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type GetMsgResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMsgResp) Reset()         { *m = GetMsgResp{} }
func (m *GetMsgResp) String() string { return proto.CompactTextString(m) }
func (*GetMsgResp) ProtoMessage()    {}
func (*GetMsgResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_2141552de9bf11d0, []int{7}
}

func (m *GetMsgResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMsgResp.Unmarshal(m, b)
}
func (m *GetMsgResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMsgResp.Marshal(b, m, deterministic)
}
func (m *GetMsgResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMsgResp.Merge(m, src)
}
func (m *GetMsgResp) XXX_Size() int {
	return xxx_messageInfo_GetMsgResp.Size(m)
}
func (m *GetMsgResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMsgResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetMsgResp proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Request)(nil), "sample.Request")
	proto.RegisterMapType((map[string]string)(nil), "sample.Request.MetaDataEntry")
	proto.RegisterType((*Reply)(nil), "sample.Reply")
	proto.RegisterType((*PingRequest)(nil), "sample.PingRequest")
	proto.RegisterType((*PongReply)(nil), "sample.PongReply")
	proto.RegisterType((*UploadReq)(nil), "sample.UploadReq")
	proto.RegisterType((*UploadResp)(nil), "sample.UploadResp")
	proto.RegisterType((*GetMsgReq)(nil), "sample.GetMsgReq")
	proto.RegisterType((*GetMsgResp)(nil), "sample.GetMsgResp")
}

func init() { proto.RegisterFile("sample.proto", fileDescriptor_2141552de9bf11d0) }

var fileDescriptor_2141552de9bf11d0 = []byte{
	// 855 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x95, 0xdd, 0x6e, 0x1b, 0x45,
	0x14, 0x80, 0xd9, 0xf8, 0x6f, 0xf7, 0xd8, 0x29, 0xe9, 0x34, 0x52, 0xdc, 0x4d, 0x81, 0xb0, 0x12,
	0x92, 0x15, 0x11, 0x2f, 0x4d, 0x50, 0x45, 0x03, 0x08, 0x9a, 0x04, 0x41, 0x40, 0x45, 0xd5, 0x1a,
	0xb8, 0xb5, 0xc6, 0xde, 0xc9, 0x66, 0xd4, 0xdd, 0x99, 0xa9, 0x67, 0x9c, 0x74, 0x73, 0xc9, 0x23,
	0xc0, 0x2d, 0xaf, 0xc0, 0x8b, 0x20, 0x71, 0x85, 0xc4, 0x2d, 0x12, 0xe2, 0x41, 0xd0, 0x99, 0xd9,
	0xb5, 0x83, 0x2d, 0x90, 0xb8, 0xea, 0x95, 0xcf, 0xcf, 0x7c, 0x67, 0xe6, 0xfc, 0xad, 0xa1, 0xa7,
	0x69, 0xa1, 0x72, 0x36, 0x54, 0x33, 0x69, 0x24, 0x69, 0x3b, 0x2d, 0xbc, 0x9f, 0x49, 0x99, 0xe5,
	0x2c, 0xb6, 0xd6, 0xc9, 0xfc, 0x22, 0xa6, 0xa2, 0x74, 0x47, 0xc2, 0x07, 0xab, 0x2e, 0x6d, 0x66,
	0xf3, 0xa9, 0xa9, 0xbc, 0x6f, 0xae, 0x7a, 0xaf, 0x67, 0x54, 0x29, 0x36, 0xd3, 0x2b, 0x34, 0x55,
	0x3c, 0xa6, 0x42, 0x48, 0x43, 0x0d, 0x97, 0xa2, 0xf6, 0x3e, 0xca, 0xb8, 0xb9, 0x9c, 0x4f, 0x86,
	0x53, 0x59, 0xc4, 0xc5, 0x35, 0x37, 0xcf, 0xe5, 0x75, 0x9c, 0xc9, 0x03, 0xeb, 0x3c, 0xb8, 0xa2,
	0x39, 0x4f, 0xa9, 0x91, 0x33, 0x1d, 0x2f, 0xc4, 0x8a, 0x7b, 0xd7, 0xfe, 0x4c, 0x0f, 0x32, 0x26,
	0x0e, 0xf4, 0x35, 0xcd, 0x32, 0x36, 0x8b, 0xa5, 0xb2, 0x91, 0xd7, 0x6f, 0x89, 0x7e, 0xf5, 0xa0,
	0x93, 0xb0, 0x17, 0x73, 0xa6, 0x0d, 0x21, 0xd0, 0x14, 0xb4, 0x60, 0x7d, 0x6f, 0xcf, 0x1b, 0x04,
	0x89, 0x95, 0x49, 0x08, 0x0d, 0x9a, 0xb1, 0xfe, 0xc6, 0x9e, 0x37, 0x68, 0x9d, 0xf8, 0x7f, 0xfe,
	0xf1, 0x56, 0x73, 0x6b, 0xbb, 0xbf, 0x9b, 0xa0, 0x91, 0x1c, 0x43, 0x50, 0x30, 0x43, 0xc7, 0x29,
	0x35, 0xb4, 0xdf, 0xd8, 0x6b, 0x0c, 0xba, 0x87, 0x6f, 0x0c, 0xab, 0x12, 0x56, 0x31, 0x87, 0x4f,
	0x99, 0xa1, 0x67, 0xd4, 0xd0, 0xcf, 0x84, 0x99, 0x95, 0x89, 0x5f, 0x54, 0x2a, 0xb9, 0x0f, 0xbe,
	0x66, 0x2f, 0xc7, 0xa6, 0x54, 0xac, 0xdf, 0xdc, 0xf3, 0x06, 0x7e, 0xd2, 0xd1, 0xec, 0xe5, 0x37,
	0xa5, 0x62, 0xe1, 0x87, 0xb0, 0xf9, 0x0f, 0x8a, 0x6c, 0x41, 0xe3, 0x39, 0x2b, 0xab, 0x67, 0xa1,
	0x48, 0xb6, 0xa1, 0x75, 0x45, 0xf3, 0xb9, 0x7b, 0x57, 0x90, 0x38, 0xe5, 0x78, 0xe3, 0x03, 0x2f,
	0x7a, 0x1b, 0x5a, 0x09, 0x53, 0x79, 0x49, 0xfa, 0xd0, 0x29, 0x98, 0xd6, 0xf8, 0x78, 0x07, 0xd6,
	0x6a, 0xf4, 0x53, 0x03, 0xba, 0xcf, 0xb8, 0xc8, 0x6e, 0xa5, 0xad, 0xb8, 0xc8, 0xea, 0xb4, 0x51,
	0x26, 0x8f, 0xa1, 0xa7, 0xcd, 0x78, 0x99, 0x1d, 0xde, 0xd3, 0x3d, 0xdc, 0x19, 0xba, 0x8e, 0x0d,
	0xeb, 0x8e, 0x0e, 0x47, 0xb6, 0xdf, 0x09, 0x68, 0x53, 0x3f, 0x99, 0x7c, 0x0c, 0x77, 0xb4, 0x19,
	0x4f, 0xe7, 0xda, 0xc8, 0xa2, 0x2e, 0xcd, 0x7f, 0xc2, 0x3d, 0x6d, 0x4e, 0xed, 0x69, 0x8b, 0x7f,
	0x02, 0xbd, 0x54, 0xce, 0x27, 0x39, 0x1b, 0xbb, 0x0c, 0x5b, 0x16, 0x7e, 0xb0, 0x06, 0x9f, 0xd9,
	0x43, 0xdf, 0xe1, 0x99, 0xa4, 0x9b, 0x2e, 0x15, 0x0c, 0xa0, 0xcd, 0x8c, 0x8b, 0xac, 0x0a, 0xd0,
	0xfe, 0x97, 0x00, 0x23, 0x7b, 0xa8, 0x0a, 0xa0, 0x97, 0x0a, 0x79, 0x0c, 0x30, 0x91, 0x32, 0xaf,
	0xf0, 0x8e, 0xc5, 0xc3, 0x35, 0xfc, 0x44, 0xca, 0xdc, 0xc1, 0xc1, 0xa4, 0x16, 0xc9, 0x47, 0xd0,
	0xe5, 0xc2, 0x3c, 0x7a, 0xbf, 0x62, 0x7d, 0xcb, 0xee, 0xae, 0xb1, 0xe7, 0x78, 0xc6, 0xc1, 0xc0,
	0x17, 0xf2, 0x97, 0x4d, 0xbf, 0xb9, 0xd5, 0x8a, 0x7e, 0xf1, 0x20, 0x78, 0x26, 0xb1, 0x3d, 0xd8,
	0x46, 0x6c, 0x8e, 0xbc, 0xd5, 0x1c, 0xf9, 0x4a, 0x9b, 0x13, 0x83, 0x4f, 0x45, 0xe9, 0xc0, 0xa6,
	0x05, 0xb7, 0xd7, 0xc0, 0x27, 0xa2, 0x4c, 0x3a, 0x54, 0x94, 0x08, 0x44, 0x3f, 0x7b, 0x10, 0x7c,
	0xab, 0x72, 0x49, 0xd3, 0x84, 0xbd, 0xc0, 0x99, 0x9c, 0x4a, 0x61, 0x98, 0x30, 0x36, 0x9f, 0x5e,
	0x52, 0xab, 0x64, 0x17, 0x82, 0x0b, 0x9e, 0xb3, 0xb1, 0xdd, 0x3f, 0x37, 0xd4, 0x3e, 0x1a, 0xbe,
	0xc6, 0x1d, 0xac, 0x9d, 0x9a, 0xdf, 0x30, 0xfb, 0xde, 0x86, 0x73, 0x8e, 0xf8, 0x8d, 0x75, 0x16,
	0xbc, 0x60, 0xcb, 0x4d, 0x0a, 0x12, 0x1f, 0x0d, 0xb8, 0x4a, 0x24, 0x04, 0xff, 0x82, 0x0b, 0xae,
	0x2f, 0x59, 0x6a, 0x07, 0xc9, 0x4f, 0x16, 0x3a, 0xb9, 0x07, 0x2d, 0xad, 0xc6, 0x3c, 0xb5, 0x03,
	0x12, 0x24, 0x4d, 0xad, 0xce, 0xd3, 0xe8, 0x1d, 0x80, 0xfa, 0xb9, 0x5a, 0x91, 0x1d, 0xe8, 0xd8,
	0x8b, 0x79, 0x5a, 0xd5, 0xbf, 0x8d, 0xea, 0x79, 0x1a, 0x1d, 0x41, 0xf0, 0x39, 0x33, 0x4f, 0x35,
	0xee, 0x10, 0xae, 0x67, 0xa1, 0xeb, 0x0e, 0xa1, 0x88, 0xeb, 0x39, 0x95, 0x73, 0x61, 0xdc, 0x67,
	0x23, 0x71, 0x4a, 0xd4, 0x03, 0xa8, 0x21, 0xad, 0x0e, 0x7f, 0xdf, 0x80, 0xf6, 0xc8, 0x7e, 0x2b,
	0xc8, 0xa7, 0xe0, 0x8f, 0x68, 0xf9, 0x05, 0xcb, 0x73, 0x49, 0x5e, 0x5f, 0xf9, 0x80, 0x84, 0x9b,
	0x4b, 0x83, 0xca, 0xcb, 0x68, 0xfb, 0xfb, 0xdf, 0xfe, 0xfa, 0x71, 0xe3, 0x4e, 0x14, 0xc4, 0x57,
	0x0f, 0xe3, 0x4b, 0x44, 0x8e, 0xbd, 0x7d, 0xf2, 0x15, 0xf8, 0xb8, 0xd1, 0x38, 0x36, 0xe4, 0x5e,
	0x0d, 0xdc, 0xda, 0xf1, 0xf0, 0xee, 0xc2, 0x58, 0x4f, 0x56, 0xb4, 0x63, 0x23, 0xdd, 0x8d, 0x7a,
	0x18, 0x09, 0x97, 0x1e, 0x67, 0x0b, 0x83, 0x1d, 0x41, 0xdb, 0xd5, 0x80, 0x2c, 0xa8, 0x45, 0x0b,
	0x43, 0xb2, 0x6a, 0xd2, 0x2a, 0x7a, 0x6d, 0xe0, 0x91, 0x53, 0x68, 0xbb, 0xe4, 0x96, 0xd0, 0xa2,
	0x42, 0x4b, 0x68, 0x99, 0x7f, 0x44, 0xec, 0xf5, 0x3d, 0x02, 0x78, 0x7d, 0xc6, 0x0c, 0xd6, 0xed,
	0x14, 0xda, 0x67, 0x2c, 0xff, 0xff, 0x41, 0xf6, 0x6d, 0x90, 0x94, 0xe5, 0x85, 0xce, 0x4e, 0x76,
	0x7f, 0x78, 0xd2, 0x27, 0x9b, 0x50, 0xfd, 0x7b, 0x1d, 0x36, 0x1e, 0x0e, 0xdf, 0x8b, 0x9a, 0xf8,
	0x1f, 0xb3, 0xef, 0x79, 0x93, 0xb6, 0x1d, 0xd3, 0xa3, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xd4,
	0x8e, 0x2f, 0x4b, 0xea, 0x06, 0x00, 0x00,
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
	GetMsg(ctx context.Context, in *GetMsgReq, opts ...grpc.CallOption) (*GetMsgResp, error)
	DelMsg(ctx context.Context, in *GetMsgReq, opts ...grpc.CallOption) (*GetMsgResp, error)
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

func (c *sampleClient) GetMsg(ctx context.Context, in *GetMsgReq, opts ...grpc.CallOption) (*GetMsgResp, error) {
	out := new(GetMsgResp)
	err := c.cc.Invoke(ctx, "/sample.Sample/GetMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sampleClient) DelMsg(ctx context.Context, in *GetMsgReq, opts ...grpc.CallOption) (*GetMsgResp, error) {
	out := new(GetMsgResp)
	err := c.cc.Invoke(ctx, "/sample.Sample/DelMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SampleServer is the server API for Sample service.
type SampleServer interface {
	SayHello(context.Context, *Request) (*Reply, error)
	PingPong(context.Context, *PingRequest) (*PongReply, error)
	Upload(Sample_UploadServer) error
	GetMsg(context.Context, *GetMsgReq) (*GetMsgResp, error)
	DelMsg(context.Context, *GetMsgReq) (*GetMsgResp, error)
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
func (*UnimplementedSampleServer) GetMsg(ctx context.Context, req *GetMsgReq) (*GetMsgResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMsg not implemented")
}
func (*UnimplementedSampleServer) DelMsg(ctx context.Context, req *GetMsgReq) (*GetMsgResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelMsg not implemented")
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

func _Sample_GetMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SampleServer).GetMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sample.Sample/GetMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SampleServer).GetMsg(ctx, req.(*GetMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sample_DelMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SampleServer).DelMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sample.Sample/DelMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SampleServer).DelMsg(ctx, req.(*GetMsgReq))
	}
	return interceptor(ctx, in, info, handler)
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
		{
			MethodName: "GetMsg",
			Handler:    _Sample_GetMsg_Handler,
		},
		{
			MethodName: "DelMsg",
			Handler:    _Sample_DelMsg_Handler,
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
