// Code generated by protoc-gen-go. DO NOT EDIT.
// source: marabunta.proto

package marabunta

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type StreamRequest struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamRequest) Reset()         { *m = StreamRequest{} }
func (m *StreamRequest) String() string { return proto.CompactTextString(m) }
func (*StreamRequest) ProtoMessage()    {}
func (*StreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f992b7cdbde158a, []int{0}
}

func (m *StreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamRequest.Unmarshal(m, b)
}
func (m *StreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamRequest.Marshal(b, m, deterministic)
}
func (m *StreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamRequest.Merge(m, src)
}
func (m *StreamRequest) XXX_Size() int {
	return xxx_messageInfo_StreamRequest.Size(m)
}
func (m *StreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamRequest proto.InternalMessageInfo

func (m *StreamRequest) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type StreamResponse struct {
	// Types that are valid to be assigned to Event:
	//	*StreamResponse_EDo
	//	*StreamResponse_EPing
	//	*StreamResponse_EPulse
	//	*StreamResponse_EStatus
	Event                isStreamResponse_Event `protobuf_oneof:"event"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *StreamResponse) Reset()         { *m = StreamResponse{} }
func (m *StreamResponse) String() string { return proto.CompactTextString(m) }
func (*StreamResponse) ProtoMessage()    {}
func (*StreamResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f992b7cdbde158a, []int{1}
}

func (m *StreamResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamResponse.Unmarshal(m, b)
}
func (m *StreamResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamResponse.Marshal(b, m, deterministic)
}
func (m *StreamResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamResponse.Merge(m, src)
}
func (m *StreamResponse) XXX_Size() int {
	return xxx_messageInfo_StreamResponse.Size(m)
}
func (m *StreamResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamResponse proto.InternalMessageInfo

type isStreamResponse_Event interface {
	isStreamResponse_Event()
}

type StreamResponse_EDo struct {
	EDo *StreamResponse_Do `protobuf:"bytes,1,opt,name=e_do,json=eDo,proto3,oneof"`
}

type StreamResponse_EPing struct {
	EPing *StreamResponse_Ping `protobuf:"bytes,2,opt,name=e_ping,json=ePing,proto3,oneof"`
}

type StreamResponse_EPulse struct {
	EPulse *StreamResponse_Pulse `protobuf:"bytes,3,opt,name=e_pulse,json=ePulse,proto3,oneof"`
}

type StreamResponse_EStatus struct {
	EStatus *StreamResponse_Status `protobuf:"bytes,4,opt,name=e_status,json=eStatus,proto3,oneof"`
}

func (*StreamResponse_EDo) isStreamResponse_Event() {}

func (*StreamResponse_EPing) isStreamResponse_Event() {}

func (*StreamResponse_EPulse) isStreamResponse_Event() {}

func (*StreamResponse_EStatus) isStreamResponse_Event() {}

func (m *StreamResponse) GetEvent() isStreamResponse_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *StreamResponse) GetEDo() *StreamResponse_Do {
	if x, ok := m.GetEvent().(*StreamResponse_EDo); ok {
		return x.EDo
	}
	return nil
}

func (m *StreamResponse) GetEPing() *StreamResponse_Ping {
	if x, ok := m.GetEvent().(*StreamResponse_EPing); ok {
		return x.EPing
	}
	return nil
}

func (m *StreamResponse) GetEPulse() *StreamResponse_Pulse {
	if x, ok := m.GetEvent().(*StreamResponse_EPulse); ok {
		return x.EPulse
	}
	return nil
}

func (m *StreamResponse) GetEStatus() *StreamResponse_Status {
	if x, ok := m.GetEvent().(*StreamResponse_EStatus); ok {
		return x.EStatus
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*StreamResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*StreamResponse_EDo)(nil),
		(*StreamResponse_EPing)(nil),
		(*StreamResponse_EPulse)(nil),
		(*StreamResponse_EStatus)(nil),
	}
}

type StreamResponse_Do struct {
	Uuid                 []byte   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamResponse_Do) Reset()         { *m = StreamResponse_Do{} }
func (m *StreamResponse_Do) String() string { return proto.CompactTextString(m) }
func (*StreamResponse_Do) ProtoMessage()    {}
func (*StreamResponse_Do) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f992b7cdbde158a, []int{1, 0}
}

func (m *StreamResponse_Do) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamResponse_Do.Unmarshal(m, b)
}
func (m *StreamResponse_Do) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamResponse_Do.Marshal(b, m, deterministic)
}
func (m *StreamResponse_Do) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamResponse_Do.Merge(m, src)
}
func (m *StreamResponse_Do) XXX_Size() int {
	return xxx_messageInfo_StreamResponse_Do.Size(m)
}
func (m *StreamResponse_Do) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamResponse_Do.DiscardUnknown(m)
}

var xxx_messageInfo_StreamResponse_Do proto.InternalMessageInfo

func (m *StreamResponse_Do) GetUuid() []byte {
	if m != nil {
		return m.Uuid
	}
	return nil
}

type StreamResponse_Ping struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamResponse_Ping) Reset()         { *m = StreamResponse_Ping{} }
func (m *StreamResponse_Ping) String() string { return proto.CompactTextString(m) }
func (*StreamResponse_Ping) ProtoMessage()    {}
func (*StreamResponse_Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f992b7cdbde158a, []int{1, 1}
}

func (m *StreamResponse_Ping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamResponse_Ping.Unmarshal(m, b)
}
func (m *StreamResponse_Ping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamResponse_Ping.Marshal(b, m, deterministic)
}
func (m *StreamResponse_Ping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamResponse_Ping.Merge(m, src)
}
func (m *StreamResponse_Ping) XXX_Size() int {
	return xxx_messageInfo_StreamResponse_Ping.Size(m)
}
func (m *StreamResponse_Ping) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamResponse_Ping.DiscardUnknown(m)
}

var xxx_messageInfo_StreamResponse_Ping proto.InternalMessageInfo

func (m *StreamResponse_Ping) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type StreamResponse_Pulse struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamResponse_Pulse) Reset()         { *m = StreamResponse_Pulse{} }
func (m *StreamResponse_Pulse) String() string { return proto.CompactTextString(m) }
func (*StreamResponse_Pulse) ProtoMessage()    {}
func (*StreamResponse_Pulse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f992b7cdbde158a, []int{1, 2}
}

func (m *StreamResponse_Pulse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamResponse_Pulse.Unmarshal(m, b)
}
func (m *StreamResponse_Pulse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamResponse_Pulse.Marshal(b, m, deterministic)
}
func (m *StreamResponse_Pulse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamResponse_Pulse.Merge(m, src)
}
func (m *StreamResponse_Pulse) XXX_Size() int {
	return xxx_messageInfo_StreamResponse_Pulse.Size(m)
}
func (m *StreamResponse_Pulse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamResponse_Pulse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamResponse_Pulse proto.InternalMessageInfo

func (m *StreamResponse_Pulse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type StreamResponse_Status struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamResponse_Status) Reset()         { *m = StreamResponse_Status{} }
func (m *StreamResponse_Status) String() string { return proto.CompactTextString(m) }
func (*StreamResponse_Status) ProtoMessage()    {}
func (*StreamResponse_Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f992b7cdbde158a, []int{1, 3}
}

func (m *StreamResponse_Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamResponse_Status.Unmarshal(m, b)
}
func (m *StreamResponse_Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamResponse_Status.Marshal(b, m, deterministic)
}
func (m *StreamResponse_Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamResponse_Status.Merge(m, src)
}
func (m *StreamResponse_Status) XXX_Size() int {
	return xxx_messageInfo_StreamResponse_Status.Size(m)
}
func (m *StreamResponse_Status) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamResponse_Status.DiscardUnknown(m)
}

var xxx_messageInfo_StreamResponse_Status proto.InternalMessageInfo

func (m *StreamResponse_Status) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type PayloadRequest struct {
	Uuid                 []byte   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PayloadRequest) Reset()         { *m = PayloadRequest{} }
func (m *PayloadRequest) String() string { return proto.CompactTextString(m) }
func (*PayloadRequest) ProtoMessage()    {}
func (*PayloadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f992b7cdbde158a, []int{2}
}

func (m *PayloadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PayloadRequest.Unmarshal(m, b)
}
func (m *PayloadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PayloadRequest.Marshal(b, m, deterministic)
}
func (m *PayloadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PayloadRequest.Merge(m, src)
}
func (m *PayloadRequest) XXX_Size() int {
	return xxx_messageInfo_PayloadRequest.Size(m)
}
func (m *PayloadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PayloadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PayloadRequest proto.InternalMessageInfo

func (m *PayloadRequest) GetUuid() []byte {
	if m != nil {
		return m.Uuid
	}
	return nil
}

type PayloadResponse struct {
	Uuid                 []byte        `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Task                 string        `protobuf:"bytes,2,opt,name=task,proto3" json:"task,omitempty"`
	Attachment           []*Attachment `protobuf:"bytes,3,rep,name=attachment,proto3" json:"attachment,omitempty"`
	Metadata             []*Metadata   `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *PayloadResponse) Reset()         { *m = PayloadResponse{} }
func (m *PayloadResponse) String() string { return proto.CompactTextString(m) }
func (*PayloadResponse) ProtoMessage()    {}
func (*PayloadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f992b7cdbde158a, []int{3}
}

func (m *PayloadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PayloadResponse.Unmarshal(m, b)
}
func (m *PayloadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PayloadResponse.Marshal(b, m, deterministic)
}
func (m *PayloadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PayloadResponse.Merge(m, src)
}
func (m *PayloadResponse) XXX_Size() int {
	return xxx_messageInfo_PayloadResponse.Size(m)
}
func (m *PayloadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PayloadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PayloadResponse proto.InternalMessageInfo

func (m *PayloadResponse) GetUuid() []byte {
	if m != nil {
		return m.Uuid
	}
	return nil
}

func (m *PayloadResponse) GetTask() string {
	if m != nil {
		return m.Task
	}
	return ""
}

func (m *PayloadResponse) GetAttachment() []*Attachment {
	if m != nil {
		return m.Attachment
	}
	return nil
}

func (m *PayloadResponse) GetMetadata() []*Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type Attachment struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	File                 []byte   `protobuf:"bytes,2,opt,name=file,proto3" json:"file,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Attachment) Reset()         { *m = Attachment{} }
func (m *Attachment) String() string { return proto.CompactTextString(m) }
func (*Attachment) ProtoMessage()    {}
func (*Attachment) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f992b7cdbde158a, []int{4}
}

func (m *Attachment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Attachment.Unmarshal(m, b)
}
func (m *Attachment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Attachment.Marshal(b, m, deterministic)
}
func (m *Attachment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Attachment.Merge(m, src)
}
func (m *Attachment) XXX_Size() int {
	return xxx_messageInfo_Attachment.Size(m)
}
func (m *Attachment) XXX_DiscardUnknown() {
	xxx_messageInfo_Attachment.DiscardUnknown(m)
}

var xxx_messageInfo_Attachment proto.InternalMessageInfo

func (m *Attachment) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Attachment) GetFile() []byte {
	if m != nil {
		return m.File
	}
	return nil
}

type Metadata struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Val                  string   `protobuf:"bytes,2,opt,name=val,proto3" json:"val,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f992b7cdbde158a, []int{5}
}

func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metadata.Unmarshal(m, b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
}
func (m *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(m, src)
}
func (m *Metadata) XXX_Size() int {
	return xxx_messageInfo_Metadata.Size(m)
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Metadata) GetVal() string {
	if m != nil {
		return m.Val
	}
	return ""
}

type UpdateRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f992b7cdbde158a, []int{6}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type UpdateResponse struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f992b7cdbde158a, []int{7}
}

func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (m *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(m, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

func (m *UpdateResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func init() {
	proto.RegisterType((*StreamRequest)(nil), "marabunta.StreamRequest")
	proto.RegisterType((*StreamResponse)(nil), "marabunta.StreamResponse")
	proto.RegisterType((*StreamResponse_Do)(nil), "marabunta.StreamResponse.Do")
	proto.RegisterType((*StreamResponse_Ping)(nil), "marabunta.StreamResponse.Ping")
	proto.RegisterType((*StreamResponse_Pulse)(nil), "marabunta.StreamResponse.Pulse")
	proto.RegisterType((*StreamResponse_Status)(nil), "marabunta.StreamResponse.Status")
	proto.RegisterType((*PayloadRequest)(nil), "marabunta.PayloadRequest")
	proto.RegisterType((*PayloadResponse)(nil), "marabunta.PayloadResponse")
	proto.RegisterType((*Attachment)(nil), "marabunta.Attachment")
	proto.RegisterType((*Metadata)(nil), "marabunta.Metadata")
	proto.RegisterType((*UpdateRequest)(nil), "marabunta.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "marabunta.UpdateResponse")
}

func init() { proto.RegisterFile("marabunta.proto", fileDescriptor_3f992b7cdbde158a) }

var fileDescriptor_3f992b7cdbde158a = []byte{
	// 475 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0xdd, 0x8a, 0xd3, 0x40,
	0x14, 0x6e, 0x7e, 0x9a, 0xb6, 0x67, 0x77, 0xbb, 0x32, 0x22, 0x4c, 0x83, 0x68, 0x8d, 0x5e, 0xf4,
	0xaa, 0x6a, 0x55, 0x04, 0x61, 0xc1, 0x95, 0x5e, 0xf4, 0x66, 0x61, 0x99, 0xc5, 0xeb, 0x32, 0x6b,
	0x8e, 0xdd, 0x92, 0x26, 0x13, 0x9b, 0xc9, 0x42, 0x1f, 0xc7, 0x27, 0xf2, 0x41, 0x7c, 0x09, 0x99,
	0x93, 0xa4, 0x4d, 0x68, 0xea, 0xdd, 0xc7, 0x39, 0xdf, 0x77, 0x7e, 0xbe, 0x33, 0x03, 0x97, 0xb1,
	0xdc, 0xca, 0xfb, 0x3c, 0xd1, 0x72, 0x9a, 0x6e, 0x95, 0x56, 0x6c, 0xb0, 0x0f, 0x04, 0xaf, 0xe0,
	0xe2, 0x4e, 0x6f, 0x51, 0xc6, 0x02, 0x7f, 0xe5, 0x98, 0x69, 0xf6, 0x04, 0x9c, 0x38, 0x5b, 0x71,
	0x6b, 0x6c, 0x4d, 0x06, 0xc2, 0xc0, 0xe0, 0xaf, 0x0d, 0xc3, 0x8a, 0x93, 0xa5, 0x2a, 0xc9, 0x90,
	0xbd, 0x07, 0x17, 0x97, 0xa1, 0x22, 0xd6, 0xd9, 0xec, 0xf9, 0xf4, 0xd0, 0xa0, 0x49, 0x9c, 0xce,
	0xd5, 0xa2, 0x23, 0x1c, 0x9c, 0x2b, 0xf6, 0x19, 0x3c, 0x5c, 0xa6, 0xeb, 0x64, 0xc5, 0x6d, 0x12,
	0xbd, 0x38, 0x2d, 0xba, 0x5d, 0x27, 0xab, 0x45, 0x47, 0x74, 0xd1, 0x00, 0xf6, 0x05, 0x7a, 0xb8,
	0x4c, 0xf3, 0x4d, 0x86, 0xdc, 0x21, 0xe5, 0xcb, 0xff, 0x28, 0x0d, 0x6d, 0xd1, 0x11, 0x1e, 0x12,
	0x62, 0x57, 0xd0, 0xc7, 0x65, 0xa6, 0xa5, 0xce, 0x33, 0xee, 0x92, 0x78, 0x7c, 0x5a, 0x7c, 0x47,
	0xbc, 0x45, 0x47, 0xf4, 0xb0, 0x80, 0x3e, 0x07, 0x7b, 0xae, 0x18, 0x03, 0x37, 0xcf, 0xd7, 0x21,
	0x2d, 0x7b, 0x2e, 0x08, 0xfb, 0x1c, 0x5c, 0x1a, 0xee, 0xc8, 0x2d, 0x7f, 0x04, 0xdd, 0xa2, 0xf7,
	0x71, 0xca, 0x07, 0xaf, 0x28, 0x7c, 0x9c, 0xfb, 0xd6, 0x83, 0x2e, 0x3e, 0x62, 0xa2, 0x83, 0x37,
	0x30, 0xbc, 0x95, 0xbb, 0x8d, 0x92, 0x61, 0x75, 0x91, 0x96, 0xfe, 0xc1, 0x6f, 0x0b, 0x2e, 0xf7,
	0xb4, 0xf2, 0x28, 0x2d, 0x3c, 0x13, 0xd3, 0x32, 0x8b, 0xc8, 0xf3, 0x81, 0x20, 0xcc, 0x3e, 0x01,
	0x48, 0xad, 0xe5, 0x8f, 0x87, 0x18, 0x13, 0xcd, 0x9d, 0xb1, 0x33, 0x39, 0x9b, 0x3d, 0xab, 0xd9,
	0x72, 0xbd, 0x4f, 0x8a, 0x1a, 0x91, 0xbd, 0x85, 0x7e, 0x8c, 0x5a, 0x86, 0x52, 0x4b, 0xee, 0x92,
	0xe8, 0x69, 0x4d, 0x74, 0x53, 0xa6, 0xc4, 0x9e, 0x14, 0x7c, 0x04, 0x38, 0x94, 0x32, 0x93, 0xa4,
	0x52, 0x3f, 0x94, 0x3b, 0x13, 0x36, 0xb1, 0x9f, 0xeb, 0x0d, 0xd2, 0x74, 0xe7, 0x82, 0x70, 0x30,
	0x85, 0x7e, 0x55, 0xcb, 0xd8, 0x14, 0xe1, 0xae, 0xb2, 0x29, 0xc2, 0x9d, 0x89, 0x3c, 0xca, 0x4d,
	0xb9, 0x8e, 0x81, 0xc1, 0x6b, 0xb8, 0xf8, 0x9e, 0x86, 0x52, 0x63, 0xcd, 0xae, 0x44, 0xc6, 0x58,
	0x35, 0x32, 0x38, 0x18, 0xc3, 0xb0, 0x22, 0x95, 0x66, 0x0d, 0xc1, 0x56, 0x11, 0x71, 0xfa, 0xc2,
	0x56, 0xd1, 0xec, 0x8f, 0x05, 0x83, 0x9b, 0x6a, 0x1b, 0xf6, 0x15, 0x7a, 0xa5, 0xbb, 0x6c, 0x54,
	0x5b, 0xb2, 0x79, 0x18, 0xdf, 0x6f, 0x4b, 0x95, 0xf5, 0xaf, 0xcd, 0xad, 0xcd, 0xf3, 0x62, 0xbc,
	0xe5, 0xc5, 0x15, 0xfa, 0xd1, 0xc9, 0xb7, 0x38, 0xb1, 0xde, 0x59, 0xec, 0x0a, 0xbc, 0x62, 0xe8,
	0x46, 0x89, 0xc6, 0xb2, 0x8d, 0x12, 0xcd, 0x0d, 0xef, 0x3d, 0xfa, 0xeb, 0x1f, 0xfe, 0x05, 0x00,
	0x00, 0xff, 0xff, 0x42, 0xf4, 0xa7, 0xa8, 0xfe, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MarabuntaClient is the client API for Marabunta service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MarabuntaClient interface {
	Payload(ctx context.Context, in *PayloadRequest, opts ...grpc.CallOption) (*PayloadResponse, error)
	Stream(ctx context.Context, opts ...grpc.CallOption) (Marabunta_StreamClient, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
}

type marabuntaClient struct {
	cc *grpc.ClientConn
}

func NewMarabuntaClient(cc *grpc.ClientConn) MarabuntaClient {
	return &marabuntaClient{cc}
}

func (c *marabuntaClient) Payload(ctx context.Context, in *PayloadRequest, opts ...grpc.CallOption) (*PayloadResponse, error) {
	out := new(PayloadResponse)
	err := c.cc.Invoke(ctx, "/marabunta.Marabunta/Payload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marabuntaClient) Stream(ctx context.Context, opts ...grpc.CallOption) (Marabunta_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Marabunta_serviceDesc.Streams[0], "/marabunta.Marabunta/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &marabuntaStreamClient{stream}
	return x, nil
}

type Marabunta_StreamClient interface {
	Send(*StreamRequest) error
	Recv() (*StreamResponse, error)
	grpc.ClientStream
}

type marabuntaStreamClient struct {
	grpc.ClientStream
}

func (x *marabuntaStreamClient) Send(m *StreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *marabuntaStreamClient) Recv() (*StreamResponse, error) {
	m := new(StreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *marabuntaClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/marabunta.Marabunta/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MarabuntaServer is the server API for Marabunta service.
type MarabuntaServer interface {
	Payload(context.Context, *PayloadRequest) (*PayloadResponse, error)
	Stream(Marabunta_StreamServer) error
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
}

// UnimplementedMarabuntaServer can be embedded to have forward compatible implementations.
type UnimplementedMarabuntaServer struct {
}

func (*UnimplementedMarabuntaServer) Payload(ctx context.Context, req *PayloadRequest) (*PayloadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Payload not implemented")
}
func (*UnimplementedMarabuntaServer) Stream(srv Marabunta_StreamServer) error {
	return status.Errorf(codes.Unimplemented, "method Stream not implemented")
}
func (*UnimplementedMarabuntaServer) Update(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}

func RegisterMarabuntaServer(s *grpc.Server, srv MarabuntaServer) {
	s.RegisterService(&_Marabunta_serviceDesc, srv)
}

func _Marabunta_Payload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayloadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarabuntaServer).Payload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/marabunta.Marabunta/Payload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarabuntaServer).Payload(ctx, req.(*PayloadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Marabunta_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MarabuntaServer).Stream(&marabuntaStreamServer{stream})
}

type Marabunta_StreamServer interface {
	Send(*StreamResponse) error
	Recv() (*StreamRequest, error)
	grpc.ServerStream
}

type marabuntaStreamServer struct {
	grpc.ServerStream
}

func (x *marabuntaStreamServer) Send(m *StreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *marabuntaStreamServer) Recv() (*StreamRequest, error) {
	m := new(StreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Marabunta_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarabuntaServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/marabunta.Marabunta/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarabuntaServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Marabunta_serviceDesc = grpc.ServiceDesc{
	ServiceName: "marabunta.Marabunta",
	HandlerType: (*MarabuntaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Payload",
			Handler:    _Marabunta_Payload_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Marabunta_Update_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _Marabunta_Stream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "marabunta.proto",
}
