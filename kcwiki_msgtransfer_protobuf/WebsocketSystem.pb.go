// Code generated by protoc-gen-go. DO NOT EDIT.
// source: WebsocketSystem.proto

package kcwiki_msgtransfer_protobuf

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

type WebsocketSystem_SystemMessageType int32

const (
	WebsocketSystem_SYSTEM_INFO            WebsocketSystem_SystemMessageType = 0
	WebsocketSystem_DEBUG_MSG              WebsocketSystem_SystemMessageType = 1
	WebsocketSystem_PAYLOAD_ERROR          WebsocketSystem_SystemMessageType = 2
	WebsocketSystem_CLIENT_REBOOT          WebsocketSystem_SystemMessageType = 3
	WebsocketSystem_CLIENT_SHUTDOWN        WebsocketSystem_SystemMessageType = 4
	WebsocketSystem_SERVER_REBOOT          WebsocketSystem_SystemMessageType = 5
	WebsocketSystem_SERVER_SHUTDOWN        WebsocketSystem_SystemMessageType = 6
	WebsocketSystem_AUTHORIZATION_REQUIRED WebsocketSystem_SystemMessageType = 7
	WebsocketSystem_AUTHORIZATION_FAIL     WebsocketSystem_SystemMessageType = 8
	WebsocketSystem_AUTHORIZATION_SUCCESS  WebsocketSystem_SystemMessageType = 9
	WebsocketSystem_PING                   WebsocketSystem_SystemMessageType = 10
)

var WebsocketSystem_SystemMessageType_name = map[int32]string{
	0:  "SYSTEM_INFO",
	1:  "DEBUG_MSG",
	2:  "PAYLOAD_ERROR",
	3:  "CLIENT_REBOOT",
	4:  "CLIENT_SHUTDOWN",
	5:  "SERVER_REBOOT",
	6:  "SERVER_SHUTDOWN",
	7:  "AUTHORIZATION_REQUIRED",
	8:  "AUTHORIZATION_FAIL",
	9:  "AUTHORIZATION_SUCCESS",
	10: "PING",
}
var WebsocketSystem_SystemMessageType_value = map[string]int32{
	"SYSTEM_INFO":            0,
	"DEBUG_MSG":              1,
	"PAYLOAD_ERROR":          2,
	"CLIENT_REBOOT":          3,
	"CLIENT_SHUTDOWN":        4,
	"SERVER_REBOOT":          5,
	"SERVER_SHUTDOWN":        6,
	"AUTHORIZATION_REQUIRED": 7,
	"AUTHORIZATION_FAIL":     8,
	"AUTHORIZATION_SUCCESS":  9,
	"PING":                   10,
}

func (x WebsocketSystem_SystemMessageType) String() string {
	return proto.EnumName(WebsocketSystem_SystemMessageType_name, int32(x))
}
func (WebsocketSystem_SystemMessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_WebsocketSystem_03fa01f3f9f94f01, []int{0, 0}
}

type WebsocketSystem struct {
	MsgType              WebsocketSystem_SystemMessageType `protobuf:"varint,1,opt,name=msg_type,json=msgType,proto3,enum=kcwiki.msgtransfer.protobuf.WebsocketSystem_SystemMessageType" json:"msg_type,omitempty"`
	Data                 string                            `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *WebsocketSystem) Reset()         { *m = WebsocketSystem{} }
func (m *WebsocketSystem) String() string { return proto.CompactTextString(m) }
func (*WebsocketSystem) ProtoMessage()    {}
func (*WebsocketSystem) Descriptor() ([]byte, []int) {
	return fileDescriptor_WebsocketSystem_03fa01f3f9f94f01, []int{0}
}
func (m *WebsocketSystem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WebsocketSystem.Unmarshal(m, b)
}
func (m *WebsocketSystem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WebsocketSystem.Marshal(b, m, deterministic)
}
func (dst *WebsocketSystem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WebsocketSystem.Merge(dst, src)
}
func (m *WebsocketSystem) XXX_Size() int {
	return xxx_messageInfo_WebsocketSystem.Size(m)
}
func (m *WebsocketSystem) XXX_DiscardUnknown() {
	xxx_messageInfo_WebsocketSystem.DiscardUnknown(m)
}

var xxx_messageInfo_WebsocketSystem proto.InternalMessageInfo

func (m *WebsocketSystem) GetMsgType() WebsocketSystem_SystemMessageType {
	if m != nil {
		return m.MsgType
	}
	return WebsocketSystem_SYSTEM_INFO
}

func (m *WebsocketSystem) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*WebsocketSystem)(nil), "kcwiki.msgtransfer.protobuf.WebsocketSystem")
	proto.RegisterEnum("kcwiki.msgtransfer.protobuf.WebsocketSystem_SystemMessageType", WebsocketSystem_SystemMessageType_name, WebsocketSystem_SystemMessageType_value)
}

func init() {
	proto.RegisterFile("WebsocketSystem.proto", fileDescriptor_WebsocketSystem_03fa01f3f9f94f01)
}

var fileDescriptor_WebsocketSystem_03fa01f3f9f94f01 = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x5f, 0x6b, 0xf2, 0x30,
	0x14, 0xc6, 0xdf, 0xfa, 0x3a, 0xff, 0x9c, 0xa1, 0xc6, 0x6c, 0x8a, 0xdb, 0x6e, 0xc4, 0x2b, 0xaf,
	0x7a, 0xb1, 0xdd, 0x0f, 0x5a, 0x1b, 0x35, 0xa0, 0x4d, 0x97, 0xa4, 0x13, 0x77, 0x53, 0xaa, 0x8b,
	0x45, 0xa4, 0x53, 0x4c, 0xc7, 0xf0, 0x53, 0x0f, 0xf6, 0x09, 0x46, 0xd5, 0x09, 0x76, 0xb0, 0xab,
	0x3c, 0xfc, 0x9e, 0xdf, 0xc9, 0x81, 0x03, 0x8d, 0x89, 0x9a, 0xe9, 0xf5, 0x7c, 0xa5, 0x12, 0xb1,
	0xd3, 0x89, 0x8a, 0xcd, 0xcd, 0x76, 0x9d, 0xac, 0xf1, 0xdd, 0x6a, 0xfe, 0xb1, 0x5c, 0x2d, 0xcd,
	0x58, 0x47, 0xc9, 0x36, 0x7c, 0xd3, 0x0b, 0xb5, 0x3d, 0x34, 0xb3, 0xf7, 0x45, 0xe7, 0x2b, 0x07,
	0xb5, 0xcc, 0x18, 0x9e, 0x42, 0x29, 0xd6, 0x51, 0x90, 0xec, 0x36, 0xaa, 0x65, 0xb4, 0x8d, 0x6e,
	0xf5, 0xfe, 0xd1, 0xfc, 0xe3, 0x0f, 0x33, 0xbb, 0xf6, 0xf0, 0x8c, 0x95, 0xd6, 0x61, 0xa4, 0xe4,
	0x6e, 0xa3, 0x78, 0x31, 0xd6, 0x51, 0x1a, 0x30, 0x86, 0xfc, 0x6b, 0x98, 0x84, 0xad, 0x5c, 0xdb,
	0xe8, 0x96, 0xf9, 0x3e, 0x77, 0x3e, 0x0d, 0xa8, 0xff, 0x1a, 0xc1, 0x35, 0xb8, 0x14, 0x53, 0x21,
	0xc9, 0x38, 0xa0, 0x6e, 0x9f, 0xa1, 0x7f, 0xb8, 0x02, 0x65, 0x87, 0xd8, 0xfe, 0x20, 0x18, 0x8b,
	0x01, 0x32, 0x70, 0x1d, 0x2a, 0x9e, 0x35, 0x1d, 0x31, 0xcb, 0x09, 0x08, 0xe7, 0x8c, 0xa3, 0x5c,
	0x8a, 0x7a, 0x23, 0x4a, 0x5c, 0x19, 0x70, 0x62, 0x33, 0x26, 0xd1, 0x7f, 0x7c, 0x05, 0xb5, 0x23,
	0x12, 0x43, 0x5f, 0x3a, 0x6c, 0xe2, 0xa2, 0x7c, 0xea, 0x09, 0xc2, 0x9f, 0x09, 0xff, 0xf1, 0x2e,
	0x52, 0xef, 0x88, 0x4e, 0x5e, 0x01, 0xdf, 0x42, 0xd3, 0xf2, 0xe5, 0x90, 0x71, 0xfa, 0x62, 0x49,
	0xca, 0xdc, 0x80, 0x93, 0x27, 0x9f, 0x72, 0xe2, 0xa0, 0x22, 0x6e, 0x02, 0x3e, 0xef, 0xfa, 0x16,
	0x1d, 0xa1, 0x12, 0xbe, 0x81, 0xc6, 0x39, 0x17, 0x7e, 0xaf, 0x47, 0x84, 0x40, 0x65, 0x5c, 0x82,
	0xbc, 0x47, 0xdd, 0x01, 0x02, 0xbb, 0x0b, 0xd5, 0xd3, 0xf1, 0xf6, 0xc1, 0xbe, 0xce, 0xdc, 0xd0,
	0x4b, 0xa9, 0x67, 0xcc, 0x0a, 0xfb, 0xfa, 0xe1, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xd6, 0xa7, 0x9a,
	0x7c, 0xdb, 0x01, 0x00, 0x00,
}
