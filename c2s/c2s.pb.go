// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/protofiles/c2s.proto

package c2s

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

// client  <-> server  10000 - 19999
type MessageCmd int32

const (
	MessageCmd_NULL       MessageCmd = 0
	MessageCmd_JOIN_ROOM  MessageCmd = 1
	MessageCmd_FRAME      MessageCmd = 2
	MessageCmd_FRAMES     MessageCmd = 3
	MessageCmd_LEAVE_ROOM MessageCmd = 4
	MessageCmd_HEARTBEAT  MessageCmd = 5
	MessageCmd_START_GAME MessageCmd = 6
)

var MessageCmd_name = map[int32]string{
	0: "NULL",
	1: "JOIN_ROOM",
	2: "FRAME",
	3: "FRAMES",
	4: "LEAVE_ROOM",
	5: "HEARTBEAT",
	6: "START_GAME",
}

var MessageCmd_value = map[string]int32{
	"NULL":       0,
	"JOIN_ROOM":  1,
	"FRAME":      2,
	"FRAMES":     3,
	"LEAVE_ROOM": 4,
	"HEARTBEAT":  5,
	"START_GAME": 6,
}

func (x MessageCmd) String() string {
	return proto.EnumName(MessageCmd_name, int32(x))
}

func (MessageCmd) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1657ad7e68cde744, []int{0}
}

type ErrorCode int32

const (
	ErrorCode_OK ErrorCode = 0
)

var ErrorCode_name = map[int32]string{
	0: "OK",
}

var ErrorCode_value = map[string]int32{
	"OK": 0,
}

func (x ErrorCode) String() string {
	return proto.EnumName(ErrorCode_name, int32(x))
}

func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1657ad7e68cde744, []int{1}
}

type Result struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_1657ad7e68cde744, []int{0}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Result) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type JoinRoomReq struct {
	Uid                  uint64   `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	RoomID               int32    `protobuf:"varint,2,opt,name=roomID,proto3" json:"roomID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinRoomReq) Reset()         { *m = JoinRoomReq{} }
func (m *JoinRoomReq) String() string { return proto.CompactTextString(m) }
func (*JoinRoomReq) ProtoMessage()    {}
func (*JoinRoomReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_1657ad7e68cde744, []int{1}
}

func (m *JoinRoomReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinRoomReq.Unmarshal(m, b)
}
func (m *JoinRoomReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinRoomReq.Marshal(b, m, deterministic)
}
func (m *JoinRoomReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinRoomReq.Merge(m, src)
}
func (m *JoinRoomReq) XXX_Size() int {
	return xxx_messageInfo_JoinRoomReq.Size(m)
}
func (m *JoinRoomReq) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinRoomReq.DiscardUnknown(m)
}

var xxx_messageInfo_JoinRoomReq proto.InternalMessageInfo

func (m *JoinRoomReq) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *JoinRoomReq) GetRoomID() int32 {
	if m != nil {
		return m.RoomID
	}
	return 0
}

type JoinRoomResp struct {
	Result               *Result  `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Uid                  uint64   `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Index                int32    `protobuf:"varint,3,opt,name=index,proto3" json:"index,omitempty"`
	Seed                 int32    `protobuf:"varint,4,opt,name=seed,proto3" json:"seed,omitempty"`
	FrameIndex           int32    `protobuf:"varint,5,opt,name=frameIndex,proto3" json:"frameIndex,omitempty"`
	RoomID               int32    `protobuf:"varint,6,opt,name=roomID,proto3" json:"roomID,omitempty"`
	Frame                []*Frame `protobuf:"bytes,7,rep,name=frame,proto3" json:"frame,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinRoomResp) Reset()         { *m = JoinRoomResp{} }
func (m *JoinRoomResp) String() string { return proto.CompactTextString(m) }
func (*JoinRoomResp) ProtoMessage()    {}
func (*JoinRoomResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_1657ad7e68cde744, []int{2}
}

func (m *JoinRoomResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinRoomResp.Unmarshal(m, b)
}
func (m *JoinRoomResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinRoomResp.Marshal(b, m, deterministic)
}
func (m *JoinRoomResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinRoomResp.Merge(m, src)
}
func (m *JoinRoomResp) XXX_Size() int {
	return xxx_messageInfo_JoinRoomResp.Size(m)
}
func (m *JoinRoomResp) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinRoomResp.DiscardUnknown(m)
}

var xxx_messageInfo_JoinRoomResp proto.InternalMessageInfo

func (m *JoinRoomResp) GetResult() *Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *JoinRoomResp) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *JoinRoomResp) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *JoinRoomResp) GetSeed() int32 {
	if m != nil {
		return m.Seed
	}
	return 0
}

func (m *JoinRoomResp) GetFrameIndex() int32 {
	if m != nil {
		return m.FrameIndex
	}
	return 0
}

func (m *JoinRoomResp) GetRoomID() int32 {
	if m != nil {
		return m.RoomID
	}
	return 0
}

func (m *JoinRoomResp) GetFrame() []*Frame {
	if m != nil {
		return m.Frame
	}
	return nil
}

type LeaveRoomReq struct {
	Uid                  uint64   `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	RoomID               int32    `protobuf:"varint,2,opt,name=roomID,proto3" json:"roomID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeaveRoomReq) Reset()         { *m = LeaveRoomReq{} }
func (m *LeaveRoomReq) String() string { return proto.CompactTextString(m) }
func (*LeaveRoomReq) ProtoMessage()    {}
func (*LeaveRoomReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_1657ad7e68cde744, []int{3}
}

func (m *LeaveRoomReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LeaveRoomReq.Unmarshal(m, b)
}
func (m *LeaveRoomReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LeaveRoomReq.Marshal(b, m, deterministic)
}
func (m *LeaveRoomReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaveRoomReq.Merge(m, src)
}
func (m *LeaveRoomReq) XXX_Size() int {
	return xxx_messageInfo_LeaveRoomReq.Size(m)
}
func (m *LeaveRoomReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaveRoomReq.DiscardUnknown(m)
}

var xxx_messageInfo_LeaveRoomReq proto.InternalMessageInfo

func (m *LeaveRoomReq) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *LeaveRoomReq) GetRoomID() int32 {
	if m != nil {
		return m.RoomID
	}
	return 0
}

type LeaveRoomResp struct {
	Result               *Result  `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeaveRoomResp) Reset()         { *m = LeaveRoomResp{} }
func (m *LeaveRoomResp) String() string { return proto.CompactTextString(m) }
func (*LeaveRoomResp) ProtoMessage()    {}
func (*LeaveRoomResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_1657ad7e68cde744, []int{4}
}

func (m *LeaveRoomResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LeaveRoomResp.Unmarshal(m, b)
}
func (m *LeaveRoomResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LeaveRoomResp.Marshal(b, m, deterministic)
}
func (m *LeaveRoomResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaveRoomResp.Merge(m, src)
}
func (m *LeaveRoomResp) XXX_Size() int {
	return xxx_messageInfo_LeaveRoomResp.Size(m)
}
func (m *LeaveRoomResp) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaveRoomResp.DiscardUnknown(m)
}

var xxx_messageInfo_LeaveRoomResp proto.InternalMessageInfo

func (m *LeaveRoomResp) GetResult() *Result {
	if m != nil {
		return m.Result
	}
	return nil
}

type Cmd struct {
	Cmd                  int32    `protobuf:"varint,1,opt,name=cmd,proto3" json:"cmd,omitempty"`
	Opt                  int32    `protobuf:"varint,2,opt,name=opt,proto3" json:"opt,omitempty"`
	Data                 string   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cmd) Reset()         { *m = Cmd{} }
func (m *Cmd) String() string { return proto.CompactTextString(m) }
func (*Cmd) ProtoMessage()    {}
func (*Cmd) Descriptor() ([]byte, []int) {
	return fileDescriptor_1657ad7e68cde744, []int{5}
}

func (m *Cmd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cmd.Unmarshal(m, b)
}
func (m *Cmd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cmd.Marshal(b, m, deterministic)
}
func (m *Cmd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cmd.Merge(m, src)
}
func (m *Cmd) XXX_Size() int {
	return xxx_messageInfo_Cmd.Size(m)
}
func (m *Cmd) XXX_DiscardUnknown() {
	xxx_messageInfo_Cmd.DiscardUnknown(m)
}

var xxx_messageInfo_Cmd proto.InternalMessageInfo

func (m *Cmd) GetCmd() int32 {
	if m != nil {
		return m.Cmd
	}
	return 0
}

func (m *Cmd) GetOpt() int32 {
	if m != nil {
		return m.Opt
	}
	return 0
}

func (m *Cmd) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type Frame struct {
	FrameIndex           int32    `protobuf:"varint,1,opt,name=frameIndex,proto3" json:"frameIndex,omitempty"`
	Uid                  uint64   `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Index                int32    `protobuf:"varint,3,opt,name=index,proto3" json:"index,omitempty"`
	Cmds                 []*Cmd   `protobuf:"bytes,4,rep,name=cmds,proto3" json:"cmds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Frame) Reset()         { *m = Frame{} }
func (m *Frame) String() string { return proto.CompactTextString(m) }
func (*Frame) ProtoMessage()    {}
func (*Frame) Descriptor() ([]byte, []int) {
	return fileDescriptor_1657ad7e68cde744, []int{6}
}

func (m *Frame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Frame.Unmarshal(m, b)
}
func (m *Frame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Frame.Marshal(b, m, deterministic)
}
func (m *Frame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Frame.Merge(m, src)
}
func (m *Frame) XXX_Size() int {
	return xxx_messageInfo_Frame.Size(m)
}
func (m *Frame) XXX_DiscardUnknown() {
	xxx_messageInfo_Frame.DiscardUnknown(m)
}

var xxx_messageInfo_Frame proto.InternalMessageInfo

func (m *Frame) GetFrameIndex() int32 {
	if m != nil {
		return m.FrameIndex
	}
	return 0
}

func (m *Frame) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *Frame) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Frame) GetCmds() []*Cmd {
	if m != nil {
		return m.Cmds
	}
	return nil
}

type FrameReq struct {
	RoomID               int32    `protobuf:"varint,1,opt,name=roomID,proto3" json:"roomID,omitempty"`
	Frame                *Frame   `protobuf:"bytes,2,opt,name=frame,proto3" json:"frame,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FrameReq) Reset()         { *m = FrameReq{} }
func (m *FrameReq) String() string { return proto.CompactTextString(m) }
func (*FrameReq) ProtoMessage()    {}
func (*FrameReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_1657ad7e68cde744, []int{7}
}

func (m *FrameReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FrameReq.Unmarshal(m, b)
}
func (m *FrameReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FrameReq.Marshal(b, m, deterministic)
}
func (m *FrameReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FrameReq.Merge(m, src)
}
func (m *FrameReq) XXX_Size() int {
	return xxx_messageInfo_FrameReq.Size(m)
}
func (m *FrameReq) XXX_DiscardUnknown() {
	xxx_messageInfo_FrameReq.DiscardUnknown(m)
}

var xxx_messageInfo_FrameReq proto.InternalMessageInfo

func (m *FrameReq) GetRoomID() int32 {
	if m != nil {
		return m.RoomID
	}
	return 0
}

func (m *FrameReq) GetFrame() *Frame {
	if m != nil {
		return m.Frame
	}
	return nil
}

type FrameResp struct {
	Result               *Result  `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Uid                  uint64   `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	FrameIndex           int32    `protobuf:"varint,3,opt,name=frameIndex,proto3" json:"frameIndex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FrameResp) Reset()         { *m = FrameResp{} }
func (m *FrameResp) String() string { return proto.CompactTextString(m) }
func (*FrameResp) ProtoMessage()    {}
func (*FrameResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_1657ad7e68cde744, []int{8}
}

func (m *FrameResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FrameResp.Unmarshal(m, b)
}
func (m *FrameResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FrameResp.Marshal(b, m, deterministic)
}
func (m *FrameResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FrameResp.Merge(m, src)
}
func (m *FrameResp) XXX_Size() int {
	return xxx_messageInfo_FrameResp.Size(m)
}
func (m *FrameResp) XXX_DiscardUnknown() {
	xxx_messageInfo_FrameResp.DiscardUnknown(m)
}

var xxx_messageInfo_FrameResp proto.InternalMessageInfo

func (m *FrameResp) GetResult() *Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *FrameResp) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *FrameResp) GetFrameIndex() int32 {
	if m != nil {
		return m.FrameIndex
	}
	return 0
}

type FramesNotify struct {
	NextFrame            int32    `protobuf:"varint,1,opt,name=nextFrame,proto3" json:"nextFrame,omitempty"`
	Frame                []*Frame `protobuf:"bytes,2,rep,name=frame,proto3" json:"frame,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FramesNotify) Reset()         { *m = FramesNotify{} }
func (m *FramesNotify) String() string { return proto.CompactTextString(m) }
func (*FramesNotify) ProtoMessage()    {}
func (*FramesNotify) Descriptor() ([]byte, []int) {
	return fileDescriptor_1657ad7e68cde744, []int{9}
}

func (m *FramesNotify) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FramesNotify.Unmarshal(m, b)
}
func (m *FramesNotify) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FramesNotify.Marshal(b, m, deterministic)
}
func (m *FramesNotify) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FramesNotify.Merge(m, src)
}
func (m *FramesNotify) XXX_Size() int {
	return xxx_messageInfo_FramesNotify.Size(m)
}
func (m *FramesNotify) XXX_DiscardUnknown() {
	xxx_messageInfo_FramesNotify.DiscardUnknown(m)
}

var xxx_messageInfo_FramesNotify proto.InternalMessageInfo

func (m *FramesNotify) GetNextFrame() int32 {
	if m != nil {
		return m.NextFrame
	}
	return 0
}

func (m *FramesNotify) GetFrame() []*Frame {
	if m != nil {
		return m.Frame
	}
	return nil
}

func init() {
	proto.RegisterEnum("c2s.MessageCmd", MessageCmd_name, MessageCmd_value)
	proto.RegisterEnum("c2s.ErrorCode", ErrorCode_name, ErrorCode_value)
	proto.RegisterType((*Result)(nil), "c2s.Result")
	proto.RegisterType((*JoinRoomReq)(nil), "c2s.JoinRoomReq")
	proto.RegisterType((*JoinRoomResp)(nil), "c2s.JoinRoomResp")
	proto.RegisterType((*LeaveRoomReq)(nil), "c2s.LeaveRoomReq")
	proto.RegisterType((*LeaveRoomResp)(nil), "c2s.LeaveRoomResp")
	proto.RegisterType((*Cmd)(nil), "c2s.Cmd")
	proto.RegisterType((*Frame)(nil), "c2s.Frame")
	proto.RegisterType((*FrameReq)(nil), "c2s.FrameReq")
	proto.RegisterType((*FrameResp)(nil), "c2s.FrameResp")
	proto.RegisterType((*FramesNotify)(nil), "c2s.FramesNotify")
}

func init() { proto.RegisterFile("protobuf/protofiles/c2s.proto", fileDescriptor_1657ad7e68cde744) }

var fileDescriptor_1657ad7e68cde744 = []byte{
	// 489 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xad, 0xbf, 0x96, 0x78, 0xd2, 0x22, 0x6b, 0x41, 0xc8, 0x87, 0x82, 0x22, 0x73, 0x89, 0x7a,
	0x48, 0x24, 0x83, 0x04, 0x12, 0xe2, 0xe0, 0xa6, 0x2e, 0xa4, 0xe4, 0x43, 0xda, 0x06, 0x0e, 0x5c,
	0x2a, 0xc7, 0xbb, 0xa9, 0x2c, 0xba, 0xd9, 0xd4, 0xeb, 0xa0, 0xf2, 0x13, 0xf9, 0x57, 0x68, 0xc7,
	0x46, 0x31, 0x25, 0x87, 0xc2, 0xed, 0xed, 0xcc, 0xbc, 0x99, 0x37, 0x6f, 0x77, 0xe1, 0xf9, 0xa6,
	0x54, 0x95, 0x5a, 0x6e, 0x57, 0x43, 0x04, 0xab, 0xe2, 0x46, 0xe8, 0x61, 0x1e, 0xeb, 0x01, 0x1e,
	0xa9, 0x93, 0xc7, 0x3a, 0x1a, 0x00, 0x61, 0x42, 0x6f, 0x6f, 0x2a, 0x4a, 0xc1, 0xcd, 0x15, 0x17,
	0xa1, 0xd5, 0xb3, 0xfa, 0x1e, 0x43, 0x4c, 0x03, 0x70, 0xa4, 0xbe, 0x0e, 0xed, 0x9e, 0xd5, 0xf7,
	0x99, 0x81, 0xd1, 0x1b, 0xe8, 0x5e, 0xa8, 0x62, 0xcd, 0x94, 0x92, 0x4c, 0xdc, 0x9a, 0x82, 0x6d,
	0xc1, 0x91, 0xe3, 0x32, 0x03, 0xe9, 0x33, 0x20, 0xa5, 0x52, 0x72, 0x7c, 0x86, 0x2c, 0x8f, 0x35,
	0xa7, 0xe8, 0xa7, 0x05, 0x87, 0x3b, 0xa6, 0xde, 0xd0, 0x97, 0x40, 0x4a, 0x9c, 0x8c, 0xec, 0x6e,
	0xdc, 0x1d, 0x18, 0x69, 0xb5, 0x18, 0xd6, 0xa4, 0x7e, 0xf7, 0xb7, 0x77, 0xfd, 0x9f, 0x82, 0x57,
	0xac, 0xb9, 0xb8, 0x0b, 0x1d, 0x6c, 0x5f, 0x1f, 0x8c, 0x78, 0x2d, 0x04, 0x0f, 0xdd, 0x5a, 0xbc,
	0xc1, 0xf4, 0x05, 0xc0, 0xaa, 0xcc, 0xa4, 0x18, 0x63, 0xb9, 0x87, 0x99, 0x56, 0xa4, 0xa5, 0x94,
	0xb4, 0x95, 0xd2, 0x1e, 0x78, 0x58, 0x15, 0x3e, 0xea, 0x39, 0xfd, 0x6e, 0x0c, 0xa8, 0xeb, 0xdc,
	0x44, 0x58, 0x9d, 0x88, 0xde, 0xc2, 0xe1, 0x44, 0x64, 0xdf, 0xc5, 0xbf, 0xbb, 0xf0, 0x1a, 0x8e,
	0x5a, 0xcc, 0x07, 0xba, 0x10, 0xbd, 0x07, 0x67, 0x24, 0xb9, 0x19, 0x93, 0x4b, 0xde, 0x5c, 0x90,
	0x81, 0x26, 0xa2, 0x36, 0x55, 0x33, 0xc3, 0x40, 0x63, 0x04, 0xcf, 0xaa, 0x0c, 0xdd, 0xf1, 0x19,
	0xe2, 0x48, 0x82, 0x87, 0xf2, 0xef, 0x39, 0x62, 0xfd, 0xe5, 0xc8, 0x43, 0xdd, 0x3e, 0x06, 0x37,
	0x97, 0x5c, 0x87, 0x2e, 0x1a, 0xd4, 0x41, 0xc9, 0x23, 0xc9, 0x19, 0x46, 0xa3, 0x33, 0xe8, 0xd4,
	0x6e, 0x89, 0xdb, 0x96, 0x0f, 0xd6, 0x7e, 0x8f, 0x6d, 0xdc, 0x7a, 0x8f, 0xc7, 0x4b, 0xf0, 0x9b,
	0x2e, 0xff, 0xff, 0x56, 0xfe, 0xdc, 0xd7, 0xb9, 0xbf, 0x6f, 0x34, 0x83, 0x43, 0x9c, 0xa1, 0x67,
	0xaa, 0x2a, 0x56, 0x3f, 0xe8, 0x31, 0xf8, 0x6b, 0x71, 0x57, 0x61, 0xac, 0x11, 0xbc, 0x0b, 0xb4,
	0x35, 0xef, 0x7f, 0x17, 0x27, 0xdf, 0x00, 0xa6, 0x42, 0xeb, 0xec, 0x5a, 0x98, 0xeb, 0xea, 0x80,
	0x3b, 0xfb, 0x3c, 0x99, 0x04, 0x07, 0xf4, 0x08, 0xfc, 0x8b, 0xf9, 0x78, 0x76, 0xc5, 0xe6, 0xf3,
	0x69, 0x60, 0x51, 0x1f, 0xbc, 0x73, 0x96, 0x4c, 0xd3, 0xc0, 0xa6, 0x00, 0x04, 0xe1, 0x65, 0xe0,
	0xd0, 0xc7, 0x00, 0x93, 0x34, 0xf9, 0x92, 0xd6, 0x65, 0xae, 0x61, 0x7d, 0x4c, 0x13, 0xb6, 0x38,
	0x4d, 0x93, 0x45, 0xe0, 0x99, 0xf4, 0xe5, 0x22, 0x61, 0x8b, 0xab, 0x0f, 0x86, 0x4a, 0x4e, 0x9e,
	0x80, 0x9f, 0x96, 0xa5, 0x2a, 0x47, 0xe6, 0xa3, 0x12, 0xb0, 0xe7, 0x9f, 0x82, 0x83, 0xd3, 0xce,
	0x57, 0x32, 0x18, 0xbe, 0xcb, 0x63, 0xbd, 0x24, 0xf8, 0xc9, 0x5f, 0xfd, 0x0a, 0x00, 0x00, 0xff,
	0xff, 0x4a, 0xd6, 0x31, 0x15, 0x05, 0x04, 0x00, 0x00,
}
