// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: proto/consignment/consignment.proto

package consignment

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Response_Things int32

const (
	Response_FIRST  Response_Things = 0
	Response_SECOND Response_Things = 1
	Response_THIRD  Response_Things = 2
	Response_FOURTH Response_Things = 4
)

// Enum value maps for Response_Things.
var (
	Response_Things_name = map[int32]string{
		0: "FIRST",
		1: "SECOND",
		2: "THIRD",
		4: "FOURTH",
	}
	Response_Things_value = map[string]int32{
		"FIRST":  0,
		"SECOND": 1,
		"THIRD":  2,
		"FOURTH": 4,
	}
)

func (x Response_Things) Enum() *Response_Things {
	p := new(Response_Things)
	*p = x
	return p
}

func (x Response_Things) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Response_Things) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_consignment_consignment_proto_enumTypes[0].Descriptor()
}

func (Response_Things) Type() protoreflect.EnumType {
	return &file_proto_consignment_consignment_proto_enumTypes[0]
}

func (x Response_Things) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Response_Things.Descriptor instead.
func (Response_Things) EnumDescriptor() ([]byte, []int) {
	return file_proto_consignment_consignment_proto_rawDescGZIP(), []int{1, 0}
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body   string `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`      //тип имя = номер поля
	Age    int32  `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`       //тип целый 32  go  -> int32
	Salary int64  `protobuf:"varint,3,opt,name=salary,proto3" json:"salary,omitempty"` //тип целый 64  go -> int64
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_consignment_consignment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_proto_consignment_consignment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_proto_consignment_consignment_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *Request) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Request) GetSalary() int64 {
	if x != nil {
		return x.Salary
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MyDouble  float64         `protobuf:"fixed64,1,opt,name=my_double,json=myDouble,proto3" json:"my_double,omitempty"`    // go -> float64
	MyFloat   float32         `protobuf:"fixed32,2,opt,name=my_float,json=myFloat,proto3" json:"my_float,omitempty"`       // go ->float32
	MyUint32  uint32          `protobuf:"varint,3,opt,name=my_uint32,json=myUint32,proto3" json:"my_uint32,omitempty"`     // go -> uint32
	MyUint64  uint64          `protobuf:"varint,4,opt,name=my_uint64,json=myUint64,proto3" json:"my_uint64,omitempty"`     // go -> uint64
	MyFixed32 uint32          `protobuf:"fixed32,5,opt,name=my_fixed32,json=myFixed32,proto3" json:"my_fixed32,omitempty"` // go -> uint32
	MyBool    bool            `protobuf:"varint,6,opt,name=my_bool,json=myBool,proto3" json:"my_bool,omitempty"`
	MyString  string          `protobuf:"bytes,7,opt,name=my_string,json=myString,proto3" json:"my_string,omitempty"`
	MyBytes   []byte          `protobuf:"bytes,8,opt,name=my_bytes,json=myBytes,proto3" json:"my_bytes,omitempty"` //go -> []byte
	Things    Response_Things `protobuf:"varint,9,opt,name=things,proto3,enum=consignment.Response_Things" json:"things,omitempty"`
	Req       *Request        `protobuf:"bytes,10,opt,name=req,proto3" json:"req,omitempty"`
	Snippet   []*Request      `protobuf:"bytes,11,rep,name=snippet,proto3" json:"snippet,omitempty"` // go-> []Request
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_consignment_consignment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_consignment_consignment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proto_consignment_consignment_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetMyDouble() float64 {
	if x != nil {
		return x.MyDouble
	}
	return 0
}

func (x *Response) GetMyFloat() float32 {
	if x != nil {
		return x.MyFloat
	}
	return 0
}

func (x *Response) GetMyUint32() uint32 {
	if x != nil {
		return x.MyUint32
	}
	return 0
}

func (x *Response) GetMyUint64() uint64 {
	if x != nil {
		return x.MyUint64
	}
	return 0
}

func (x *Response) GetMyFixed32() uint32 {
	if x != nil {
		return x.MyFixed32
	}
	return 0
}

func (x *Response) GetMyBool() bool {
	if x != nil {
		return x.MyBool
	}
	return false
}

func (x *Response) GetMyString() string {
	if x != nil {
		return x.MyString
	}
	return ""
}

func (x *Response) GetMyBytes() []byte {
	if x != nil {
		return x.MyBytes
	}
	return nil
}

func (x *Response) GetThings() Response_Things {
	if x != nil {
		return x.Things
	}
	return Response_FIRST
}

func (x *Response) GetReq() *Request {
	if x != nil {
		return x.Req
	}
	return nil
}

func (x *Response) GetSnippet() []*Request {
	if x != nil {
		return x.Snippet
	}
	return nil
}

type TimeDate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Year  int32 `protobuf:"varint,1,opt,name=year,proto3" json:"year,omitempty"`
	Month int32 `protobuf:"varint,2,opt,name=month,proto3" json:"month,omitempty"`
	Day   int32 `protobuf:"varint,3,opt,name=day,proto3" json:"day,omitempty"`
	Hour  int32 `protobuf:"varint,4,opt,name=hour,proto3" json:"hour,omitempty"`
}

func (x *TimeDate) Reset() {
	*x = TimeDate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_consignment_consignment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeDate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeDate) ProtoMessage() {}

func (x *TimeDate) ProtoReflect() protoreflect.Message {
	mi := &file_proto_consignment_consignment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeDate.ProtoReflect.Descriptor instead.
func (*TimeDate) Descriptor() ([]byte, []int) {
	return file_proto_consignment_consignment_proto_rawDescGZIP(), []int{2}
}

func (x *TimeDate) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *TimeDate) GetMonth() int32 {
	if x != nil {
		return x.Month
	}
	return 0
}

func (x *TimeDate) GetDay() int32 {
	if x != nil {
		return x.Day
	}
	return 0
}

func (x *TimeDate) GetHour() int32 {
	if x != nil {
		return x.Hour
	}
	return 0
}

type OuterSearcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []*OuterSearcher_InnerSearcher `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
}

func (x *OuterSearcher) Reset() {
	*x = OuterSearcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_consignment_consignment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OuterSearcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OuterSearcher) ProtoMessage() {}

func (x *OuterSearcher) ProtoReflect() protoreflect.Message {
	mi := &file_proto_consignment_consignment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OuterSearcher.ProtoReflect.Descriptor instead.
func (*OuterSearcher) Descriptor() ([]byte, []int) {
	return file_proto_consignment_consignment_proto_rawDescGZIP(), []int{3}
}

func (x *OuterSearcher) GetResult() []*OuterSearcher_InnerSearcher {
	if x != nil {
		return x.Result
	}
	return nil
}

type AnotherSearcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *OuterSearcher_InnerSearcher `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *AnotherSearcher) Reset() {
	*x = AnotherSearcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_consignment_consignment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnotherSearcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnotherSearcher) ProtoMessage() {}

func (x *AnotherSearcher) ProtoReflect() protoreflect.Message {
	mi := &file_proto_consignment_consignment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnotherSearcher.ProtoReflect.Descriptor instead.
func (*AnotherSearcher) Descriptor() ([]byte, []int) {
	return file_proto_consignment_consignment_proto_rawDescGZIP(), []int{4}
}

func (x *AnotherSearcher) GetResult() *OuterSearcher_InnerSearcher {
	if x != nil {
		return x.Result
	}
	return nil
}

type OuterSearcher_InnerSearcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Url  string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *OuterSearcher_InnerSearcher) Reset() {
	*x = OuterSearcher_InnerSearcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_consignment_consignment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OuterSearcher_InnerSearcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OuterSearcher_InnerSearcher) ProtoMessage() {}

func (x *OuterSearcher_InnerSearcher) ProtoReflect() protoreflect.Message {
	mi := &file_proto_consignment_consignment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OuterSearcher_InnerSearcher.ProtoReflect.Descriptor instead.
func (*OuterSearcher_InnerSearcher) Descriptor() ([]byte, []int) {
	return file_proto_consignment_consignment_proto_rawDescGZIP(), []int{3, 0}
}

func (x *OuterSearcher_InnerSearcher) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *OuterSearcher_InnerSearcher) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_proto_consignment_consignment_proto protoreflect.FileDescriptor

var file_proto_consignment_consignment_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x22, 0x47, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x22, 0xb2, 0x03, 0x0a, 0x08,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x79, 0x5f, 0x64,
	0x6f, 0x75, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6d, 0x79, 0x44,
	0x6f, 0x75, 0x62, 0x6c, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x79, 0x5f, 0x66, 0x6c, 0x6f, 0x61,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x6d, 0x79, 0x46, 0x6c, 0x6f, 0x61, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x79, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x6d, 0x79, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x1b, 0x0a,
	0x09, 0x6d, 0x79, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x08, 0x6d, 0x79, 0x55, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x79,
	0x5f, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x18, 0x05, 0x20, 0x01, 0x28, 0x07, 0x52, 0x09,
	0x6d, 0x79, 0x46, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x79, 0x5f,
	0x62, 0x6f, 0x6f, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x6d, 0x79, 0x42, 0x6f,
	0x6f, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x79, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x79, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12,
	0x19, 0x0a, 0x08, 0x6d, 0x79, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x07, 0x6d, 0x79, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x34, 0x0a, 0x06, 0x74, 0x68,
	0x69, 0x6e, 0x67, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x6e,
	0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x06, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73,
	0x12, 0x26, 0x0a, 0x03, 0x72, 0x65, 0x71, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x63, 0x6f, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x03, 0x72, 0x65, 0x71, 0x12, 0x2e, 0x0a, 0x07, 0x73, 0x6e, 0x69, 0x70,
	0x70, 0x65, 0x74, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6e, 0x73,
	0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x07, 0x73, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x22, 0x36, 0x0a, 0x06, 0x54, 0x68, 0x69, 0x6e,
	0x67, 0x73, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x49, 0x52, 0x53, 0x54, 0x10, 0x00, 0x12, 0x0a, 0x0a,
	0x06, 0x53, 0x45, 0x43, 0x4f, 0x4e, 0x44, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x48, 0x49,
	0x52, 0x44, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x4f, 0x55, 0x52, 0x54, 0x48, 0x10, 0x04,
	0x22, 0x5a, 0x0a, 0x08, 0x54, 0x69, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x79, 0x65, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72,
	0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x64, 0x61, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x75, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x68, 0x6f, 0x75, 0x72, 0x22, 0x88, 0x01, 0x0a,
	0x0d, 0x4f, 0x75, 0x74, 0x65, 0x72, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x12, 0x40,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28,
	0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4f, 0x75, 0x74,
	0x65, 0x72, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x6e, 0x65, 0x72,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x1a, 0x35, 0x0a, 0x0d, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x53, 0x0a, 0x0f, 0x41, 0x6e, 0x6f, 0x74, 0x68,
	0x65, 0x72, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x12, 0x40, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63, 0x6f, 0x6e,
	0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4f, 0x75, 0x74, 0x65, 0x72, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x65, 0x72, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0xc1, 0x01, 0x0a,
	0x0b, 0x54, 0x65, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x07,
	0x48, 0x74, 0x74, 0x70, 0x47, 0x45, 0x54, 0x12, 0x14, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x69, 0x67,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e,
	0x63, 0x6f, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x08, 0x48, 0x74, 0x74, 0x70, 0x50, 0x4f,
	0x53, 0x54, 0x12, 0x14, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x69,
	0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x3d, 0x0a, 0x0b, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x15, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x65, 0x1a, 0x15, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x69, 0x67,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x65, 0x22, 0x00,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_consignment_consignment_proto_rawDescOnce sync.Once
	file_proto_consignment_consignment_proto_rawDescData = file_proto_consignment_consignment_proto_rawDesc
)

func file_proto_consignment_consignment_proto_rawDescGZIP() []byte {
	file_proto_consignment_consignment_proto_rawDescOnce.Do(func() {
		file_proto_consignment_consignment_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_consignment_consignment_proto_rawDescData)
	})
	return file_proto_consignment_consignment_proto_rawDescData
}

var file_proto_consignment_consignment_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_consignment_consignment_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_consignment_consignment_proto_goTypes = []interface{}{
	(Response_Things)(0),                // 0: consignment.Response.Things
	(*Request)(nil),                     // 1: consignment.Request
	(*Response)(nil),                    // 2: consignment.Response
	(*TimeDate)(nil),                    // 3: consignment.TimeDate
	(*OuterSearcher)(nil),               // 4: consignment.OuterSearcher
	(*AnotherSearcher)(nil),             // 5: consignment.AnotherSearcher
	(*OuterSearcher_InnerSearcher)(nil), // 6: consignment.OuterSearcher.InnerSearcher
}
var file_proto_consignment_consignment_proto_depIdxs = []int32{
	0, // 0: consignment.Response.things:type_name -> consignment.Response.Things
	1, // 1: consignment.Response.req:type_name -> consignment.Request
	1, // 2: consignment.Response.snippet:type_name -> consignment.Request
	6, // 3: consignment.OuterSearcher.result:type_name -> consignment.OuterSearcher.InnerSearcher
	6, // 4: consignment.AnotherSearcher.result:type_name -> consignment.OuterSearcher.InnerSearcher
	1, // 5: consignment.TestService.HttpGET:input_type -> consignment.Request
	1, // 6: consignment.TestService.HttpPOST:input_type -> consignment.Request
	3, // 7: consignment.TestService.CurrentTime:input_type -> consignment.TimeDate
	2, // 8: consignment.TestService.HttpGET:output_type -> consignment.Response
	2, // 9: consignment.TestService.HttpPOST:output_type -> consignment.Response
	3, // 10: consignment.TestService.CurrentTime:output_type -> consignment.TimeDate
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_consignment_consignment_proto_init() }
func file_proto_consignment_consignment_proto_init() {
	if File_proto_consignment_consignment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_consignment_consignment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_consignment_consignment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_consignment_consignment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeDate); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_consignment_consignment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OuterSearcher); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_consignment_consignment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnotherSearcher); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_consignment_consignment_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OuterSearcher_InnerSearcher); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_consignment_consignment_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_consignment_consignment_proto_goTypes,
		DependencyIndexes: file_proto_consignment_consignment_proto_depIdxs,
		EnumInfos:         file_proto_consignment_consignment_proto_enumTypes,
		MessageInfos:      file_proto_consignment_consignment_proto_msgTypes,
	}.Build()
	File_proto_consignment_consignment_proto = out.File
	file_proto_consignment_consignment_proto_rawDesc = nil
	file_proto_consignment_consignment_proto_goTypes = nil
	file_proto_consignment_consignment_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TestServiceClient is the client API for TestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestServiceClient interface {
	HttpGET(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	HttpPOST(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	CurrentTime(ctx context.Context, in *TimeDate, opts ...grpc.CallOption) (*TimeDate, error)
}

type testServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTestServiceClient(cc grpc.ClientConnInterface) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) HttpGET(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/consignment.TestService/HttpGET", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) HttpPOST(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/consignment.TestService/HttpPOST", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) CurrentTime(ctx context.Context, in *TimeDate, opts ...grpc.CallOption) (*TimeDate, error) {
	out := new(TimeDate)
	err := c.cc.Invoke(ctx, "/consignment.TestService/CurrentTime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestServiceServer is the server API for TestService service.
type TestServiceServer interface {
	HttpGET(context.Context, *Request) (*Response, error)
	HttpPOST(context.Context, *Request) (*Response, error)
	CurrentTime(context.Context, *TimeDate) (*TimeDate, error)
}

// UnimplementedTestServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTestServiceServer struct {
}

func (*UnimplementedTestServiceServer) HttpGET(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HttpGET not implemented")
}
func (*UnimplementedTestServiceServer) HttpPOST(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HttpPOST not implemented")
}
func (*UnimplementedTestServiceServer) CurrentTime(context.Context, *TimeDate) (*TimeDate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CurrentTime not implemented")
}

func RegisterTestServiceServer(s *grpc.Server, srv TestServiceServer) {
	s.RegisterService(&_TestService_serviceDesc, srv)
}

func _TestService_HttpGET_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).HttpGET(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consignment.TestService/HttpGET",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).HttpGET(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_HttpPOST_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).HttpPOST(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consignment.TestService/HttpPOST",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).HttpPOST(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_CurrentTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TimeDate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).CurrentTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consignment.TestService/CurrentTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).CurrentTime(ctx, req.(*TimeDate))
	}
	return interceptor(ctx, in, info, handler)
}

var _TestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "consignment.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HttpGET",
			Handler:    _TestService_HttpGET_Handler,
		},
		{
			MethodName: "HttpPOST",
			Handler:    _TestService_HttpPOST_Handler,
		},
		{
			MethodName: "CurrentTime",
			Handler:    _TestService_CurrentTime_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/consignment/consignment.proto",
}
