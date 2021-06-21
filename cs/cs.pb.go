// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/protofiles/cs.proto

package cs

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	anypb "google.golang.org/protobuf/types/known/anypb"
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

type MessageType int32

const (
	MessageType_None      MessageType = 0
	MessageType_Request   MessageType = 1
	MessageType_Response  MessageType = 2
	MessageType_Notify    MessageType = 3
	MessageType_Broadcast MessageType = 4
)

var MessageType_name = map[int32]string{
	0: "None",
	1: "Request",
	2: "Response",
	3: "Notify",
	4: "Broadcast",
}

var MessageType_value = map[string]int32{
	"None":      0,
	"Request":   1,
	"Response":  2,
	"Notify":    3,
	"Broadcast": 4,
}

func (x MessageType) String() string {
	return proto.EnumName(MessageType_name, int32(x))
}

func (MessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8d43b15b16c0125a, []int{0}
}

type Message struct {
	ID                   uint32            `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	UID                  uint64            `protobuf:"varint,2,opt,name=UID,proto3" json:"UID,omitempty"`
	MsgID                uint32            `protobuf:"varint,3,opt,name=msgID,proto3" json:"msgID,omitempty"`
	MsgType              uint32            `protobuf:"varint,4,opt,name=msgType,proto3" json:"msgType,omitempty"`
	Seq                  uint32            `protobuf:"varint,5,opt,name=seq,proto3" json:"seq,omitempty"`
	Options              map[string][]byte `protobuf:"bytes,6,rep,name=options,proto3" json:"options,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body                 *anypb.Any        `protobuf:"bytes,11,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d43b15b16c0125a, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetID() uint32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Message) GetUID() uint64 {
	if m != nil {
		return m.UID
	}
	return 0
}

func (m *Message) GetMsgID() uint32 {
	if m != nil {
		return m.MsgID
	}
	return 0
}

func (m *Message) GetMsgType() uint32 {
	if m != nil {
		return m.MsgType
	}
	return 0
}

func (m *Message) GetSeq() uint32 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *Message) GetOptions() map[string][]byte {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *Message) GetBody() *anypb.Any {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterEnum("cs.MessageType", MessageType_name, MessageType_value)
	proto.RegisterType((*Message)(nil), "cs.Message")
	proto.RegisterMapType((map[string][]byte)(nil), "cs.Message.OptionsEntry")
}

func init() { proto.RegisterFile("protobuf/protofiles/cs.proto", fileDescriptor_8d43b15b16c0125a) }

var fileDescriptor_8d43b15b16c0125a = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0x41, 0x6b, 0xfa, 0x40,
	0x10, 0xc5, 0xff, 0xd9, 0x44, 0xa3, 0x13, 0xfd, 0x13, 0x16, 0x0f, 0x5b, 0xe9, 0x21, 0xf4, 0x14,
	0x7a, 0xd8, 0x80, 0xbd, 0x14, 0x7b, 0xaa, 0xa4, 0x87, 0x1c, 0xb4, 0xb0, 0xb4, 0x97, 0xde, 0x62,
	0x1c, 0x83, 0x54, 0xb3, 0xd1, 0x89, 0x85, 0xfd, 0x04, 0xfd, 0xda, 0x25, 0x1b, 0x2d, 0xbd, 0xbd,
	0x37, 0xf3, 0x98, 0x79, 0x3f, 0xb8, 0xad, 0x4f, 0xba, 0xd1, 0xeb, 0xf3, 0x36, 0xb1, 0x62, 0xbb,
	0xdb, 0x23, 0x25, 0x05, 0x49, 0xeb, 0x38, 0x2b, 0x68, 0x7a, 0x53, 0x6a, 0x5d, 0xee, 0x31, 0xf9,
	0x0d, 0xe6, 0x95, 0xe9, 0xd6, 0x77, 0xdf, 0x0c, 0xfc, 0x25, 0x12, 0xe5, 0x25, 0xf2, 0xff, 0xc0,
	0xb2, 0x54, 0x38, 0x91, 0x13, 0x8f, 0x15, 0xcb, 0x52, 0x1e, 0x82, 0xfb, 0x9e, 0xa5, 0x82, 0x45,
	0x4e, 0xec, 0xa9, 0x56, 0xf2, 0x09, 0xf4, 0x0e, 0x54, 0x66, 0xa9, 0x70, 0x6d, 0xa8, 0x33, 0x5c,
	0x80, 0x7f, 0xa0, 0xf2, 0xcd, 0xd4, 0x28, 0x3c, 0x3b, 0xbf, 0xda, 0xf6, 0x02, 0xe1, 0x51, 0xf4,
	0xec, 0xb4, 0x95, 0x7c, 0x06, 0xbe, 0xae, 0x9b, 0x9d, 0xae, 0x48, 0xf4, 0x23, 0x37, 0x0e, 0x66,
	0x42, 0x16, 0x24, 0x2f, 0x0d, 0xe4, 0x6b, 0xb7, 0x7a, 0xa9, 0x9a, 0x93, 0x51, 0xd7, 0x20, 0x8f,
	0xc1, 0x5b, 0xeb, 0x8d, 0x11, 0x41, 0xe4, 0xc4, 0xc1, 0x6c, 0x22, 0x3b, 0x1a, 0x79, 0xa5, 0x91,
	0xcf, 0x95, 0x51, 0x36, 0x31, 0x9d, 0xc3, 0xe8, 0xef, 0x89, 0xf6, 0xff, 0x27, 0x1a, 0x8b, 0x34,
	0x54, 0xad, 0x6c, 0x09, 0xbe, 0xf2, 0xfd, 0x19, 0x2d, 0xd5, 0x48, 0x75, 0x66, 0xce, 0x1e, 0x9d,
	0xfb, 0x25, 0x04, 0x97, 0x1a, 0xb6, 0xfa, 0x00, 0xbc, 0x95, 0xae, 0x30, 0xfc, 0xc7, 0x03, 0xf0,
	0x15, 0x1e, 0xcf, 0x48, 0x4d, 0xe8, 0xf0, 0x11, 0x0c, 0x14, 0x52, 0xad, 0x2b, 0xc2, 0x90, 0x71,
	0x80, 0xfe, 0x4a, 0x37, 0xbb, 0xad, 0x09, 0x5d, 0x3e, 0x86, 0xe1, 0xe2, 0xa4, 0xf3, 0x4d, 0x91,
	0x53, 0x13, 0x7a, 0x0b, 0xff, 0xa3, 0x27, 0x93, 0xa7, 0x82, 0xd6, 0x7d, 0xdb, 0xf3, 0xe1, 0x27,
	0x00, 0x00, 0xff, 0xff, 0xb5, 0x80, 0xa4, 0x8b, 0xa7, 0x01, 0x00, 0x00,
}