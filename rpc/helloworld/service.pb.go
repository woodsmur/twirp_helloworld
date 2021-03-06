// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/helloworld/service.proto

package helloworld

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

type GoodBye struct {
	UserID               string   `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GoodBye) Reset()         { *m = GoodBye{} }
func (m *GoodBye) String() string { return proto.CompactTextString(m) }
func (*GoodBye) ProtoMessage()    {}
func (*GoodBye) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4a9b41a6e929104, []int{0}
}

func (m *GoodBye) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoodBye.Unmarshal(m, b)
}
func (m *GoodBye) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoodBye.Marshal(b, m, deterministic)
}
func (m *GoodBye) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoodBye.Merge(m, src)
}
func (m *GoodBye) XXX_Size() int {
	return xxx_messageInfo_GoodBye.Size(m)
}
func (m *GoodBye) XXX_DiscardUnknown() {
	xxx_messageInfo_GoodBye.DiscardUnknown(m)
}

var xxx_messageInfo_GoodBye proto.InternalMessageInfo

func (m *GoodBye) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *GoodBye) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Hello struct {
	UserID               string   `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Hello) Reset()         { *m = Hello{} }
func (m *Hello) String() string { return proto.CompactTextString(m) }
func (*Hello) ProtoMessage()    {}
func (*Hello) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4a9b41a6e929104, []int{1}
}

func (m *Hello) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Hello.Unmarshal(m, b)
}
func (m *Hello) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Hello.Marshal(b, m, deterministic)
}
func (m *Hello) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hello.Merge(m, src)
}
func (m *Hello) XXX_Size() int {
	return xxx_messageInfo_Hello.Size(m)
}
func (m *Hello) XXX_DiscardUnknown() {
	xxx_messageInfo_Hello.DiscardUnknown(m)
}

var xxx_messageInfo_Hello proto.InternalMessageInfo

func (m *Hello) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *Hello) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ReqHello struct {
	Hello                *Hello   `protobuf:"bytes,1,opt,name=Hello,proto3" json:"Hello,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqHello) Reset()         { *m = ReqHello{} }
func (m *ReqHello) String() string { return proto.CompactTextString(m) }
func (*ReqHello) ProtoMessage()    {}
func (*ReqHello) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4a9b41a6e929104, []int{2}
}

func (m *ReqHello) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqHello.Unmarshal(m, b)
}
func (m *ReqHello) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqHello.Marshal(b, m, deterministic)
}
func (m *ReqHello) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqHello.Merge(m, src)
}
func (m *ReqHello) XXX_Size() int {
	return xxx_messageInfo_ReqHello.Size(m)
}
func (m *ReqHello) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqHello.DiscardUnknown(m)
}

var xxx_messageInfo_ReqHello proto.InternalMessageInfo

func (m *ReqHello) GetHello() *Hello {
	if m != nil {
		return m.Hello
	}
	return nil
}

type ResHello struct {
	GoodBye              *GoodBye `protobuf:"bytes,1,opt,name=GoodBye,proto3" json:"GoodBye,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResHello) Reset()         { *m = ResHello{} }
func (m *ResHello) String() string { return proto.CompactTextString(m) }
func (*ResHello) ProtoMessage()    {}
func (*ResHello) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4a9b41a6e929104, []int{3}
}

func (m *ResHello) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResHello.Unmarshal(m, b)
}
func (m *ResHello) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResHello.Marshal(b, m, deterministic)
}
func (m *ResHello) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResHello.Merge(m, src)
}
func (m *ResHello) XXX_Size() int {
	return xxx_messageInfo_ResHello.Size(m)
}
func (m *ResHello) XXX_DiscardUnknown() {
	xxx_messageInfo_ResHello.DiscardUnknown(m)
}

var xxx_messageInfo_ResHello proto.InternalMessageInfo

func (m *ResHello) GetGoodBye() *GoodBye {
	if m != nil {
		return m.GoodBye
	}
	return nil
}

func init() {
	proto.RegisterType((*GoodBye)(nil), "twirp.example.helloworld.GoodBye")
	proto.RegisterType((*Hello)(nil), "twirp.example.helloworld.Hello")
	proto.RegisterType((*ReqHello)(nil), "twirp.example.helloworld.ReqHello")
	proto.RegisterType((*ResHello)(nil), "twirp.example.helloworld.ResHello")
}

func init() { proto.RegisterFile("rpc/helloworld/service.proto", fileDescriptor_f4a9b41a6e929104) }

var fileDescriptor_f4a9b41a6e929104 = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x29, 0x2a, 0x48, 0xd6,
	0xcf, 0x48, 0xcd, 0xc9, 0xc9, 0x2f, 0xcf, 0x2f, 0xca, 0x49, 0xd1, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb,
	0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x28, 0x29, 0xcf, 0x2c, 0x2a, 0xd0,
	0x4b, 0xad, 0x48, 0xcc, 0x2d, 0xc8, 0x49, 0xd5, 0x43, 0xa8, 0x53, 0xb2, 0xe6, 0x62, 0x77, 0xcf,
	0xcf, 0x4f, 0x71, 0xaa, 0x4c, 0x15, 0x12, 0xe3, 0x62, 0x0b, 0x2d, 0x4e, 0x2d, 0xf2, 0x74, 0x91,
	0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf2, 0x84, 0x24, 0xb8, 0xd8, 0x7d, 0x53, 0x8b, 0x8b,
	0x13, 0xd3, 0x53, 0x25, 0x98, 0xc0, 0x12, 0x30, 0xae, 0x92, 0x25, 0x17, 0xab, 0x07, 0xc8, 0x28,
	0x32, 0xb4, 0x3a, 0x72, 0x71, 0x04, 0xa5, 0x16, 0x42, 0x74, 0x9b, 0x42, 0x8d, 0x01, 0x6b, 0xe6,
	0x36, 0x92, 0xd7, 0xc3, 0xe5, 0x5a, 0x3d, 0xb0, 0xb2, 0x20, 0x88, 0x6a, 0x25, 0x77, 0x90, 0x11,
	0xc5, 0x10, 0x23, 0x10, 0xde, 0x80, 0x1a, 0xa2, 0x88, 0xdb, 0x10, 0xa8, 0xc2, 0x20, 0x98, 0x0e,
	0xa3, 0xa3, 0x8c, 0x5c, 0xc2, 0x8e, 0xe5, 0xa9, 0xc5, 0xf9, 0xb9, 0xa9, 0x21, 0x20, 0x4d, 0xc1,
	0x90, 0xb0, 0x13, 0xf2, 0xe3, 0xe2, 0x08, 0x4e, 0xac, 0x84, 0x58, 0x40, 0xc8, 0x51, 0x52, 0x84,
	0x2d, 0x14, 0x8a, 0xe3, 0x12, 0x81, 0x99, 0x17, 0x52, 0x59, 0x90, 0x5a, 0xec, 0x58, 0xec, 0x96,
	0x99, 0x9a, 0x93, 0x22, 0xa4, 0x84, 0x5b, 0x2b, 0x2c, 0x8c, 0xa4, 0xf0, 0xaa, 0x81, 0x04, 0x82,
	0x13, 0x4f, 0x14, 0x17, 0x42, 0x38, 0x89, 0x0d, 0x1c, 0xf5, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x32, 0x23, 0xbc, 0x46, 0x1a, 0x02, 0x00, 0x00,
}
