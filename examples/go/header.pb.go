// Code generated by protoc-gen-go. DO NOT EDIT.
// source: header.proto

package example

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/william-lg/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Tmp struct {
	Major                *int32   `protobuf:"varint,1,opt,name=major" json:"major,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tmp) Reset()         { *m = Tmp{} }
func (m *Tmp) String() string { return proto.CompactTextString(m) }
func (*Tmp) ProtoMessage()    {}
func (*Tmp) Descriptor() ([]byte, []int) {
	return fileDescriptor_header_9b3597c00b52c06a, []int{0}
}
func (m *Tmp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tmp.Unmarshal(m, b)
}
func (m *Tmp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tmp.Marshal(b, m, deterministic)
}
func (dst *Tmp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tmp.Merge(dst, src)
}
func (m *Tmp) XXX_Size() int {
	return xxx_messageInfo_Tmp.Size(m)
}
func (m *Tmp) XXX_DiscardUnknown() {
	xxx_messageInfo_Tmp.DiscardUnknown(m)
}

var xxx_messageInfo_Tmp proto.InternalMessageInfo

func (m *Tmp) GetMajor() int32 {
	if m != nil && m.Major != nil {
		return *m.Major
	}
	return 0
}

type Version struct {
	Major                *int32   `protobuf:"varint,1,opt,name=major" json:"major,omitempty"`
	Mior                 *int32   `protobuf:"varint,2,opt,name=mior" json:"mior,omitempty"`
	Revision             *int32   `protobuf:"varint,3,opt,name=revision" json:"revision,omitempty"`
	Tmp                  *Tmp     `protobuf:"bytes,4,opt,name=tmp" json:"tmp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Version) Reset()         { *m = Version{} }
func (m *Version) String() string { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()    {}
func (*Version) Descriptor() ([]byte, []int) {
	return fileDescriptor_header_9b3597c00b52c06a, []int{1}
}
func (m *Version) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Version.Unmarshal(m, b)
}
func (m *Version) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Version.Marshal(b, m, deterministic)
}
func (dst *Version) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Version.Merge(dst, src)
}
func (m *Version) XXX_Size() int {
	return xxx_messageInfo_Version.Size(m)
}
func (m *Version) XXX_DiscardUnknown() {
	xxx_messageInfo_Version.DiscardUnknown(m)
}

var xxx_messageInfo_Version proto.InternalMessageInfo

func (m *Version) GetMajor() int32 {
	if m != nil && m.Major != nil {
		return *m.Major
	}
	return 0
}

func (m *Version) GetMior() int32 {
	if m != nil && m.Mior != nil {
		return *m.Mior
	}
	return 0
}

func (m *Version) GetRevision() int32 {
	if m != nil && m.Revision != nil {
		return *m.Revision
	}
	return 0
}

func (m *Version) GetTmp() *Tmp {
	if m != nil {
		return m.Tmp
	}
	return nil
}

type ReqHeader struct {
	SeqId                *int64   `protobuf:"varint,1,opt,name=seq_id,json=seqId" json:"seq_id,omitempty"`
	Version              *Version `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	CallerId             *int64   `protobuf:"varint,3,opt,name=caller_id,json=callerId" json:"caller_id,omitempty"`
	CallerName           *string  `protobuf:"bytes,4,opt,name=caller_name,json=callerName" json:"caller_name,omitempty"`
	CallerIp             *string  `protobuf:"bytes,5,opt,name=caller_ip,json=callerIp" json:"caller_ip,omitempty"`
	Timestamp            *int64   `protobuf:"varint,6,opt,name=timestamp" json:"timestamp,omitempty"`
	Token                *string  `protobuf:"bytes,7,opt,name=token" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqHeader) Reset()         { *m = ReqHeader{} }
func (m *ReqHeader) String() string { return proto.CompactTextString(m) }
func (*ReqHeader) ProtoMessage()    {}
func (*ReqHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_header_9b3597c00b52c06a, []int{2}
}
func (m *ReqHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqHeader.Unmarshal(m, b)
}
func (m *ReqHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqHeader.Marshal(b, m, deterministic)
}
func (dst *ReqHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqHeader.Merge(dst, src)
}
func (m *ReqHeader) XXX_Size() int {
	return xxx_messageInfo_ReqHeader.Size(m)
}
func (m *ReqHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqHeader.DiscardUnknown(m)
}

var xxx_messageInfo_ReqHeader proto.InternalMessageInfo

func (m *ReqHeader) GetSeqId() int64 {
	if m != nil && m.SeqId != nil {
		return *m.SeqId
	}
	return 0
}

func (m *ReqHeader) GetVersion() *Version {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *ReqHeader) GetCallerId() int64 {
	if m != nil && m.CallerId != nil {
		return *m.CallerId
	}
	return 0
}

func (m *ReqHeader) GetCallerName() string {
	if m != nil && m.CallerName != nil {
		return *m.CallerName
	}
	return ""
}

func (m *ReqHeader) GetCallerIp() string {
	if m != nil && m.CallerIp != nil {
		return *m.CallerIp
	}
	return ""
}

func (m *ReqHeader) GetTimestamp() int64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *ReqHeader) GetToken() string {
	if m != nil && m.Token != nil {
		return *m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*Tmp)(nil), "example.Tmp")
	proto.RegisterType((*Version)(nil), "example.Version")
	proto.RegisterType((*ReqHeader)(nil), "example.ReqHeader")
}

func init() { proto.RegisterFile("header.proto", fileDescriptor_header_9b3597c00b52c06a) }

var fileDescriptor_header_9b3597c00b52c06a = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0x86, 0x2d, 0x05, 0x4a, 0x07, 0x0e, 0x64, 0xe3, 0x61, 0x23, 0xc4, 0x36, 0xc4, 0x18, 0x2e,
	0x94, 0x84, 0x9b, 0x57, 0x4e, 0x72, 0xf1, 0xd0, 0x10, 0xaf, 0x66, 0xa5, 0x13, 0x58, 0xdd, 0x65,
	0x97, 0xed, 0x8a, 0xbe, 0x88, 0xcf, 0x47, 0xc2, 0x0b, 0xf8, 0x0a, 0xa6, 0xdb, 0x52, 0x91, 0x5b,
	0xe7, 0xff, 0xff, 0xf9, 0xf2, 0x4f, 0x17, 0x7a, 0x1b, 0x64, 0x19, 0x9a, 0x44, 0x1b, 0x65, 0x15,
	0x09, 0xf0, 0x8b, 0x49, 0x2d, 0xf0, 0xe6, 0x61, 0xcd, 0xed, 0xe6, 0xe3, 0x35, 0x59, 0x29, 0x39,
	0xfd, 0xe4, 0x42, 0x70, 0x26, 0x27, 0x62, 0x3d, 0x5d, 0xab, 0x89, 0x0b, 0x4e, 0xf6, 0x4c, 0xf0,
	0x8c, 0x59, 0x65, 0xf2, 0x69, 0xfd, 0x59, 0x32, 0x46, 0x03, 0xf0, 0x97, 0x52, 0x93, 0x6b, 0x68,
	0x49, 0xf6, 0xa6, 0x0c, 0xf5, 0x62, 0x6f, 0xdc, 0x4a, 0xcb, 0x61, 0xf4, 0xed, 0x41, 0xf0, 0x8c,
	0x26, 0xe7, 0x6a, 0x4b, 0x6e, 0xff, 0x25, 0xe6, 0x9d, 0xe3, 0x21, 0x6a, 0xf6, 0xaf, 0x62, 0xaf,
	0xca, 0x92, 0x21, 0x34, 0x25, 0x57, 0x86, 0x36, 0x2e, 0x6c, 0xa7, 0x92, 0x3b, 0xe8, 0x18, 0xdc,
	0xf3, 0x82, 0x44, 0xfd, 0x8b, 0x44, 0xed, 0x90, 0x7b, 0xf0, 0xad, 0xd4, 0xb4, 0x19, 0x7b, 0xe3,
	0xee, 0xac, 0x97, 0x54, 0xe7, 0x25, 0x4b, 0xa9, 0xe7, 0xed, 0xe3, 0x21, 0x6a, 0xc4, 0x5e, 0x5a,
	0x04, 0x46, 0x3f, 0x1e, 0x84, 0x29, 0xee, 0x1e, 0xdd, 0xcf, 0x20, 0x11, 0xb4, 0x73, 0xdc, 0xbd,
	0xf0, 0xcc, 0x55, 0xf3, 0xcf, 0xab, 0xe5, 0xb8, 0x5b, 0x64, 0x64, 0x06, 0xc1, 0xbe, 0xbc, 0xc2,
	0xb5, 0xeb, 0xce, 0xfa, 0x35, 0xba, 0xba, 0xae, 0xc6, 0x9f, 0x82, 0x64, 0x00, 0xe1, 0x8a, 0x09,
	0x81, 0xa6, 0xe0, 0x16, 0x8d, 0xfd, 0xb4, 0x53, 0x0a, 0x8b, 0x8c, 0x44, 0xd0, 0xad, 0xcc, 0x2d,
	0x93, 0xe8, 0xfa, 0x86, 0x29, 0x94, 0xd2, 0x13, 0x93, 0x78, 0xbe, 0xad, 0x69, 0xcb, 0xd9, 0xa7,
	0x6d, 0x4d, 0x86, 0x10, 0x5a, 0x2e, 0x31, 0xb7, 0x4c, 0x6a, 0xda, 0x76, 0xe8, 0x3f, 0xa1, 0x78,
	0x09, 0xab, 0xde, 0x71, 0x4b, 0x03, 0xb7, 0x56, 0x0e, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xea,
	0x84, 0xca, 0xfb, 0xf9, 0x01, 0x00, 0x00,
}
