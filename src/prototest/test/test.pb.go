// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

package test

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MyMsg struct {
	Id                   *int32   `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Str                  *string  `protobuf:"bytes,2,req,name=str" json:"str,omitempty"`
	Opt                  *int32   `protobuf:"varint,3,opt,name=opt" json:"opt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MyMsg) Reset()         { *m = MyMsg{} }
func (m *MyMsg) String() string { return proto.CompactTextString(m) }
func (*MyMsg) ProtoMessage()    {}
func (*MyMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_d74065ee386b892c, []int{0}
}
func (m *MyMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MyMsg.Unmarshal(m, b)
}
func (m *MyMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MyMsg.Marshal(b, m, deterministic)
}
func (dst *MyMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MyMsg.Merge(dst, src)
}
func (m *MyMsg) XXX_Size() int {
	return xxx_messageInfo_MyMsg.Size(m)
}
func (m *MyMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_MyMsg.DiscardUnknown(m)
}

var xxx_messageInfo_MyMsg proto.InternalMessageInfo

func (m *MyMsg) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *MyMsg) GetStr() string {
	if m != nil && m.Str != nil {
		return *m.Str
	}
	return ""
}

func (m *MyMsg) GetOpt() int32 {
	if m != nil && m.Opt != nil {
		return *m.Opt
	}
	return 0
}

func init() {
	proto.RegisterType((*MyMsg)(nil), "test.myMsg")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_test_d74065ee386b892c) }

var fileDescriptor_test_d74065ee386b892c = []byte{
	// 94 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0xac, 0xb9, 0x58, 0x73, 0x2b,
	0x7d, 0x8b, 0xd3, 0x85, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x18, 0x15, 0x98, 0x34, 0x58, 0x83,
	0x98, 0x32, 0x53, 0x84, 0x04, 0xb8, 0x98, 0x8b, 0x4b, 0x8a, 0x24, 0x98, 0x14, 0x98, 0x34, 0x38,
	0x83, 0x40, 0x4c, 0x90, 0x48, 0x7e, 0x41, 0x89, 0x04, 0xb3, 0x02, 0xa3, 0x06, 0x6b, 0x10, 0x88,
	0x09, 0x08, 0x00, 0x00, 0xff, 0xff, 0xb4, 0x7b, 0x2f, 0x12, 0x4f, 0x00, 0x00, 0x00,
}
