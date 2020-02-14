// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/micro/services/ping/proto/ping.proto

package ping

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_c39c4501d20b0113, []int{0}
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

type Response struct {
	Ping                 string   `protobuf:"bytes,1,opt,name=ping,proto3" json:"ping,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_c39c4501d20b0113, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetPing() string {
	if m != nil {
		return m.Ping
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "ping.Request")
	proto.RegisterType((*Response)(nil), "ping.Response")
}

func init() {
	proto.RegisterFile("github.com/micro/services/ping/proto/ping.proto", fileDescriptor_c39c4501d20b0113)
}

var fileDescriptor_c39c4501d20b0113 = []byte{
	// 136 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xd2, 0x4f, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0xcd, 0x4c, 0x2e, 0xca, 0xd7, 0x2f, 0x4e, 0x2d,
	0x2a, 0xcb, 0x4c, 0x4e, 0x2d, 0xd6, 0x2f, 0xc8, 0xcc, 0x4b, 0xd7, 0x2f, 0x28, 0xca, 0x2f, 0xc9,
	0x07, 0x33, 0xf5, 0xc0, 0x4c, 0x21, 0x16, 0x10, 0x5b, 0x89, 0x93, 0x8b, 0x3d, 0x28, 0xb5, 0xb0,
	0x34, 0xb5, 0xb8, 0x44, 0x49, 0x8e, 0x8b, 0x23, 0x28, 0xb5, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55,
	0x48, 0x88, 0x0b, 0x2c, 0x2d, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0x66, 0x1b, 0xe9, 0x73,
	0xb1, 0x04, 0x00, 0x69, 0x21, 0x75, 0x28, 0xcd, 0xab, 0x07, 0x36, 0x0d, 0xaa, 0x5d, 0x8a, 0x0f,
	0xc6, 0x85, 0x18, 0xa1, 0xc4, 0x90, 0xc4, 0x06, 0xb6, 0xc8, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff,
	0xa3, 0xf5, 0xb7, 0x1e, 0x9b, 0x00, 0x00, 0x00,
}