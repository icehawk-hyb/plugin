// Code generated by protoc-gen-go. DO NOT EDIT.
// source: js.proto

package jsproto

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// create action
type Create struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Create) Reset()         { *m = Create{} }
func (m *Create) String() string { return proto.CompactTextString(m) }
func (*Create) ProtoMessage()    {}
func (*Create) Descriptor() ([]byte, []int) {
	return fileDescriptor_d11539bc790542aa, []int{0}
}

func (m *Create) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Create.Unmarshal(m, b)
}
func (m *Create) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Create.Marshal(b, m, deterministic)
}
func (m *Create) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Create.Merge(m, src)
}
func (m *Create) XXX_Size() int {
	return xxx_messageInfo_Create.Size(m)
}
func (m *Create) XXX_DiscardUnknown() {
	xxx_messageInfo_Create.DiscardUnknown(m)
}

var xxx_messageInfo_Create proto.InternalMessageInfo

func (m *Create) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Create) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// call action
type Call struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Funcname             string   `protobuf:"bytes,2,opt,name=funcname,proto3" json:"funcname,omitempty"`
	Args                 string   `protobuf:"bytes,3,opt,name=args,proto3" json:"args,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Call) Reset()         { *m = Call{} }
func (m *Call) String() string { return proto.CompactTextString(m) }
func (*Call) ProtoMessage()    {}
func (*Call) Descriptor() ([]byte, []int) {
	return fileDescriptor_d11539bc790542aa, []int{1}
}

func (m *Call) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Call.Unmarshal(m, b)
}
func (m *Call) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Call.Marshal(b, m, deterministic)
}
func (m *Call) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Call.Merge(m, src)
}
func (m *Call) XXX_Size() int {
	return xxx_messageInfo_Call.Size(m)
}
func (m *Call) XXX_DiscardUnknown() {
	xxx_messageInfo_Call.DiscardUnknown(m)
}

var xxx_messageInfo_Call proto.InternalMessageInfo

func (m *Call) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Call) GetFuncname() string {
	if m != nil {
		return m.Funcname
	}
	return ""
}

func (m *Call) GetArgs() string {
	if m != nil {
		return m.Args
	}
	return ""
}

type JsAction struct {
	// Types that are valid to be assigned to Value:
	//	*JsAction_Create
	//	*JsAction_Call
	Value                isJsAction_Value `protobuf_oneof:"value"`
	Ty                   int32            `protobuf:"varint,3,opt,name=ty,proto3" json:"ty,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *JsAction) Reset()         { *m = JsAction{} }
func (m *JsAction) String() string { return proto.CompactTextString(m) }
func (*JsAction) ProtoMessage()    {}
func (*JsAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_d11539bc790542aa, []int{2}
}

func (m *JsAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JsAction.Unmarshal(m, b)
}
func (m *JsAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JsAction.Marshal(b, m, deterministic)
}
func (m *JsAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JsAction.Merge(m, src)
}
func (m *JsAction) XXX_Size() int {
	return xxx_messageInfo_JsAction.Size(m)
}
func (m *JsAction) XXX_DiscardUnknown() {
	xxx_messageInfo_JsAction.DiscardUnknown(m)
}

var xxx_messageInfo_JsAction proto.InternalMessageInfo

type isJsAction_Value interface {
	isJsAction_Value()
}

type JsAction_Create struct {
	Create *Create `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type JsAction_Call struct {
	Call *Call `protobuf:"bytes,2,opt,name=call,proto3,oneof"`
}

func (*JsAction_Create) isJsAction_Value() {}

func (*JsAction_Call) isJsAction_Value() {}

func (m *JsAction) GetValue() isJsAction_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *JsAction) GetCreate() *Create {
	if x, ok := m.GetValue().(*JsAction_Create); ok {
		return x.Create
	}
	return nil
}

func (m *JsAction) GetCall() *Call {
	if x, ok := m.GetValue().(*JsAction_Call); ok {
		return x.Call
	}
	return nil
}

func (m *JsAction) GetTy() int32 {
	if m != nil {
		return m.Ty
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*JsAction) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*JsAction_Create)(nil),
		(*JsAction_Call)(nil),
	}
}

type JsLog struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JsLog) Reset()         { *m = JsLog{} }
func (m *JsLog) String() string { return proto.CompactTextString(m) }
func (*JsLog) ProtoMessage()    {}
func (*JsLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_d11539bc790542aa, []int{3}
}

func (m *JsLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JsLog.Unmarshal(m, b)
}
func (m *JsLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JsLog.Marshal(b, m, deterministic)
}
func (m *JsLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JsLog.Merge(m, src)
}
func (m *JsLog) XXX_Size() int {
	return xxx_messageInfo_JsLog.Size(m)
}
func (m *JsLog) XXX_DiscardUnknown() {
	xxx_messageInfo_JsLog.DiscardUnknown(m)
}

var xxx_messageInfo_JsLog proto.InternalMessageInfo

func (m *JsLog) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type QueryResult struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryResult) Reset()         { *m = QueryResult{} }
func (m *QueryResult) String() string { return proto.CompactTextString(m) }
func (*QueryResult) ProtoMessage()    {}
func (*QueryResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_d11539bc790542aa, []int{4}
}

func (m *QueryResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryResult.Unmarshal(m, b)
}
func (m *QueryResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryResult.Marshal(b, m, deterministic)
}
func (m *QueryResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryResult.Merge(m, src)
}
func (m *QueryResult) XXX_Size() int {
	return xxx_messageInfo_QueryResult.Size(m)
}
func (m *QueryResult) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryResult.DiscardUnknown(m)
}

var xxx_messageInfo_QueryResult proto.InternalMessageInfo

func (m *QueryResult) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*Create)(nil), "jsproto.Create")
	proto.RegisterType((*Call)(nil), "jsproto.Call")
	proto.RegisterType((*JsAction)(nil), "jsproto.JsAction")
	proto.RegisterType((*JsLog)(nil), "jsproto.JsLog")
	proto.RegisterType((*QueryResult)(nil), "jsproto.QueryResult")
}

func init() {
	proto.RegisterFile("js.proto", fileDescriptor_d11539bc790542aa)
}

var fileDescriptor_d11539bc790542aa = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0xc6, 0x9b, 0x90, 0xa4, 0xe1, 0x22, 0x40, 0xf2, 0x14, 0xc1, 0x02, 0x66, 0x81, 0x25, 0x42,
	0xe5, 0x09, 0xa0, 0x4b, 0x15, 0xb1, 0xe0, 0x37, 0x38, 0x5c, 0x53, 0x51, 0x1d, 0x31, 0xf2, 0x1f,
	0xa4, 0xbc, 0x3d, 0xf2, 0xa5, 0x34, 0x0c, 0x6c, 0x3f, 0xdf, 0xfd, 0xec, 0xef, 0x33, 0xd4, 0x7b,
	0xdf, 0x7d, 0x39, 0x1b, 0xac, 0x58, 0xee, 0x3d, 0x83, 0x7c, 0x80, 0x6a, 0xed, 0x0c, 0x06, 0x23,
	0x04, 0x14, 0xda, 0x6e, 0x4d, 0x9b, 0x5d, 0x67, 0x77, 0xa7, 0x8a, 0x39, 0xcd, 0x06, 0xfc, 0x34,
	0x6d, 0x3e, 0xcd, 0x12, 0xcb, 0x1e, 0x8a, 0x35, 0x12, 0x1d, 0x77, 0xd9, 0xbc, 0x13, 0x97, 0x50,
	0xbf, 0xc7, 0x41, 0xff, 0xb9, 0x73, 0x3c, 0x27, 0x1f, 0xdd, 0xce, 0xb7, 0x27, 0x93, 0x9f, 0x58,
	0x7a, 0xa8, 0x7b, 0xff, 0xa4, 0xc3, 0x87, 0x1d, 0xc4, 0x3d, 0x54, 0x9a, 0x9b, 0xf0, 0x8b, 0xcd,
	0xea, 0xa2, 0x3b, 0x74, 0xec, 0xa6, 0x82, 0x9b, 0x85, 0x3a, 0x08, 0xe2, 0x16, 0x0a, 0x8d, 0x44,
	0x1c, 0xd1, 0xac, 0xce, 0x66, 0x11, 0x89, 0x36, 0x0b, 0xc5, 0x4b, 0x71, 0x0e, 0x79, 0x18, 0x39,
	0xad, 0x54, 0x79, 0x18, 0x9f, 0x97, 0x50, 0x7e, 0x23, 0x45, 0x23, 0xaf, 0xa0, 0xec, 0xfd, 0x8b,
	0xdd, 0xa5, 0x46, 0x5b, 0x0c, 0xf8, 0xfb, 0x83, 0xc4, 0xf2, 0x06, 0x9a, 0xd7, 0x68, 0xdc, 0xa8,
	0x8c, 0x8f, 0x14, 0xfe, 0x53, 0xde, 0x2a, 0x0e, 0x7b, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0xdf,
	0xa4, 0x90, 0x55, 0x4e, 0x01, 0x00, 0x00,
}
