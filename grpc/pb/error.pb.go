// Code generated by protoc-gen-go. DO NOT EDIT.
// source: error.proto

package goa_pb

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

// ErrorResponse message defines the error encoded in the gRPC response that
// correspond to the errors created by the generated code. This is mainly
// intended for the clients to decode error responses.
type ErrorResponse struct {
	// name is the name for that class of errors.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// id is the unique error instance identifier.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// msg describes the specific error occurrence.
	Msg string `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	// temporary indicates whether the error is temporary.
	Temporary bool `protobuf:"varint,4,opt,name=temporary,proto3" json:"temporary,omitempty"`
	// timeout indicates whether the error is a timeout.
	Timeout bool `protobuf:"varint,5,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// fault indicates whether the error is a server-side fault.
	Fault                bool     `protobuf:"varint,6,opt,name=fault,proto3" json:"fault,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ErrorResponse) Reset()         { *m = ErrorResponse{} }
func (m *ErrorResponse) String() string { return proto.CompactTextString(m) }
func (*ErrorResponse) ProtoMessage()    {}
func (*ErrorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_error_3dd3810be1549eb5, []int{0}
}
func (m *ErrorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorResponse.Unmarshal(m, b)
}
func (m *ErrorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorResponse.Marshal(b, m, deterministic)
}
func (dst *ErrorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorResponse.Merge(dst, src)
}
func (m *ErrorResponse) XXX_Size() int {
	return xxx_messageInfo_ErrorResponse.Size(m)
}
func (m *ErrorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorResponse proto.InternalMessageInfo

func (m *ErrorResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ErrorResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ErrorResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *ErrorResponse) GetTemporary() bool {
	if m != nil {
		return m.Temporary
	}
	return false
}

func (m *ErrorResponse) GetTimeout() bool {
	if m != nil {
		return m.Timeout
	}
	return false
}

func (m *ErrorResponse) GetFault() bool {
	if m != nil {
		return m.Fault
	}
	return false
}

func init() {
	proto.RegisterType((*ErrorResponse)(nil), "goa.pb.ErrorResponse")
}

func init() { proto.RegisterFile("error.proto", fileDescriptor_error_3dd3810be1549eb5) }

var fileDescriptor_error_3dd3810be1549eb5 = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x2d, 0x2a, 0xca,
	0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4b, 0xcf, 0x4f, 0xd4, 0x2b, 0x48, 0x52,
	0x9a, 0xcc, 0xc8, 0xc5, 0xeb, 0x0a, 0x12, 0x0f, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15,
	0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3,
	0x85, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x98, 0xc0, 0x22, 0x4c, 0x99, 0x29, 0x42, 0x02, 0x5c,
	0xcc, 0xb9, 0xc5, 0xe9, 0x12, 0xcc, 0x60, 0x01, 0x10, 0x53, 0x48, 0x86, 0x8b, 0xb3, 0x24, 0x35,
	0xb7, 0x20, 0xbf, 0x28, 0xb1, 0xa8, 0x52, 0x82, 0x45, 0x81, 0x51, 0x83, 0x23, 0x08, 0x21, 0x20,
	0x24, 0xc1, 0xc5, 0x5e, 0x92, 0x99, 0x9b, 0x9a, 0x5f, 0x5a, 0x22, 0xc1, 0x0a, 0x96, 0x83, 0x71,
	0x85, 0x44, 0xb8, 0x58, 0xd3, 0x12, 0x4b, 0x73, 0x4a, 0x24, 0xd8, 0xc0, 0xe2, 0x10, 0x4e, 0x12,
	0x1b, 0xd8, 0x91, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x33, 0x83, 0xe5, 0xba, 0xb3, 0x00,
	0x00, 0x00,
}
