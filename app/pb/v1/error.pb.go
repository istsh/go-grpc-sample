// Code generated by protoc-gen-go. DO NOT EDIT.
// source: error.proto

package v1

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

type ErrorCode struct {
	ErrorCode            string   `protobuf:"bytes,1,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ErrorCode) Reset()         { *m = ErrorCode{} }
func (m *ErrorCode) String() string { return proto.CompactTextString(m) }
func (*ErrorCode) ProtoMessage()    {}
func (*ErrorCode) Descriptor() ([]byte, []int) {
	return fileDescriptor_0579b252106fcf4a, []int{0}
}

func (m *ErrorCode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorCode.Unmarshal(m, b)
}
func (m *ErrorCode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorCode.Marshal(b, m, deterministic)
}
func (m *ErrorCode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorCode.Merge(m, src)
}
func (m *ErrorCode) XXX_Size() int {
	return xxx_messageInfo_ErrorCode.Size(m)
}
func (m *ErrorCode) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorCode.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorCode proto.InternalMessageInfo

func (m *ErrorCode) GetErrorCode() string {
	if m != nil {
		return m.ErrorCode
	}
	return ""
}

type Error struct {
	Error                *Error_ErrorDetail `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_0579b252106fcf4a, []int{1}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetError() *Error_ErrorDetail {
	if m != nil {
		return m.Error
	}
	return nil
}

type Error_ErrorDetail struct {
	ErrorCode            string   `protobuf:"bytes,1,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	Locale               string   `protobuf:"bytes,2,opt,name=locale,proto3" json:"locale,omitempty"`
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error_ErrorDetail) Reset()         { *m = Error_ErrorDetail{} }
func (m *Error_ErrorDetail) String() string { return proto.CompactTextString(m) }
func (*Error_ErrorDetail) ProtoMessage()    {}
func (*Error_ErrorDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_0579b252106fcf4a, []int{1, 0}
}

func (m *Error_ErrorDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error_ErrorDetail.Unmarshal(m, b)
}
func (m *Error_ErrorDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error_ErrorDetail.Marshal(b, m, deterministic)
}
func (m *Error_ErrorDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error_ErrorDetail.Merge(m, src)
}
func (m *Error_ErrorDetail) XXX_Size() int {
	return xxx_messageInfo_Error_ErrorDetail.Size(m)
}
func (m *Error_ErrorDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_Error_ErrorDetail.DiscardUnknown(m)
}

var xxx_messageInfo_Error_ErrorDetail proto.InternalMessageInfo

func (m *Error_ErrorDetail) GetErrorCode() string {
	if m != nil {
		return m.ErrorCode
	}
	return ""
}

func (m *Error_ErrorDetail) GetLocale() string {
	if m != nil {
		return m.Locale
	}
	return ""
}

func (m *Error_ErrorDetail) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*ErrorCode)(nil), "v1.ErrorCode")
	proto.RegisterType((*Error)(nil), "v1.Error")
	proto.RegisterType((*Error_ErrorDetail)(nil), "v1.Error.ErrorDetail")
}

func init() { proto.RegisterFile("error.proto", fileDescriptor_0579b252106fcf4a) }

var fileDescriptor_0579b252106fcf4a = []byte{
	// 199 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8f, 0xcd, 0x4a, 0xc5, 0x30,
	0x10, 0x46, 0xe9, 0x95, 0x7b, 0xa5, 0xd3, 0x5d, 0x40, 0x09, 0x82, 0x20, 0x5d, 0xa9, 0xa5, 0x09,
	0xd5, 0x37, 0xf0, 0xe7, 0x05, 0xba, 0x74, 0xa1, 0xa4, 0xe9, 0x90, 0x16, 0x52, 0x26, 0x24, 0xb1,
	0x4f, 0xe2, 0x03, 0x8b, 0x53, 0x0b, 0xee, 0xdc, 0x04, 0xbe, 0x73, 0x0e, 0x84, 0x81, 0x0a, 0x63,
	0xa4, 0xa8, 0x42, 0xa4, 0x4c, 0xe2, 0xb0, 0x76, 0xf5, 0x3d, 0x94, 0xaf, 0x3f, 0xe8, 0x99, 0x46,
	0x14, 0xd7, 0x00, 0xec, 0x3f, 0x2c, 0x8d, 0x28, 0x8b, 0x9b, 0xe2, 0xb6, 0xec, 0x4b, 0xdc, 0x75,
	0xfd, 0x55, 0xc0, 0x91, 0x63, 0xd1, 0xc0, 0x91, 0x31, 0x37, 0xd5, 0xc3, 0x85, 0x5a, 0x3b, 0xc5,
	0x66, 0x7b, 0x5f, 0x30, 0x9b, 0xd9, 0xf7, 0x5b, 0x73, 0xf5, 0x0e, 0xd5, 0x1f, 0xfa, 0xcf, 0x27,
	0xe2, 0x12, 0x4e, 0x9e, 0xac, 0xf1, 0x28, 0x0f, 0xac, 0x7e, 0x97, 0x90, 0x70, 0xbe, 0x60, 0x4a,
	0xc6, 0xa1, 0x3c, 0x63, 0xb1, 0xcf, 0xa7, 0xe6, 0xed, 0xce, 0xcd, 0x79, 0xfa, 0x1c, 0x94, 0xa5,
	0x45, 0xcf, 0x29, 0xa7, 0x49, 0x3b, 0x6a, 0x5d, 0x0c, 0xb6, 0x4d, 0x66, 0x09, 0x1e, 0xb5, 0x09,
	0x41, 0x87, 0x41, 0xaf, 0xdd, 0x70, 0xe2, 0xd3, 0x1f, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x27,
	0x67, 0xc3, 0x73, 0x09, 0x01, 0x00, 0x00,
}