// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.9.1
// source: exchange.proto

package traceprotobuf

import (
	proto "github.com/golang/protobuf/proto"
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

type Capabilities int32

const (
	Capabilities_SKIP             Capabilities = 0
	Capabilities_ZLIB_COMPRESSION Capabilities = 1 // bit 0 of capabilities
	Capabilities_LZ4_COMPRESSION  Capabilities = 2 // bit 1 of capabilities
)

// Enum value maps for Capabilities.
var (
	Capabilities_name = map[int32]string{
		0: "SKIP",
		1: "ZLIB_COMPRESSION",
		2: "LZ4_COMPRESSION",
	}
	Capabilities_value = map[string]int32{
		"SKIP":             0,
		"ZLIB_COMPRESSION": 1,
		"LZ4_COMPRESSION":  2,
	}
)

func (x Capabilities) Enum() *Capabilities {
	p := new(Capabilities)
	*p = x
	return p
}

func (x Capabilities) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Capabilities) Descriptor() protoreflect.EnumDescriptor {
	return file_exchange_proto_enumTypes[0].Descriptor()
}

func (Capabilities) Type() protoreflect.EnumType {
	return &file_exchange_proto_enumTypes[0]
}

func (x Capabilities) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Capabilities.Descriptor instead.
func (Capabilities) EnumDescriptor() ([]byte, []int) {
	return file_exchange_proto_rawDescGZIP(), []int{0}
}

type CompressionMethod int32

const (
	CompressionMethod_NONE CompressionMethod = 0
	CompressionMethod_LZ4  CompressionMethod = 1
	CompressionMethod_ZLIB CompressionMethod = 2
)

// Enum value maps for CompressionMethod.
var (
	CompressionMethod_name = map[int32]string{
		0: "NONE",
		1: "LZ4",
		2: "ZLIB",
	}
	CompressionMethod_value = map[string]int32{
		"NONE": 0,
		"LZ4":  1,
		"ZLIB": 2,
	}
)

func (x CompressionMethod) Enum() *CompressionMethod {
	p := new(CompressionMethod)
	*p = x
	return p
}

func (x CompressionMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CompressionMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_exchange_proto_enumTypes[1].Descriptor()
}

func (CompressionMethod) Type() protoreflect.EnumType {
	return &file_exchange_proto_enumTypes[1]
}

func (x CompressionMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CompressionMethod.Descriptor instead.
func (CompressionMethod) EnumDescriptor() ([]byte, []int) {
	return file_exchange_proto_rawDescGZIP(), []int{1}
}

type ExportResponse_ResultCode int32

const (
	// Telemetry data is successfully processed by the server.
	ExportResponse_Success ExportResponse_ResultCode = 0
	// processing of telemetry data failed. The client MUST NOT retry
	// sending the same telemetry data. The telemetry data MUST be dropped.
	// This for example can happen when the request contains bad data and
	// cannot be deserialized or otherwise processed by the server.
	ExportResponse_FailedNoneRetryable ExportResponse_ResultCode = 1
	// Processing of telemetry data failed. The client SHOULD record the
	// error and MAY retry exporting the same data after some time. This
	// for example can happen when the server is overloaded.
	ExportResponse_FailedRetryable ExportResponse_ResultCode = 2
)

// Enum value maps for ExportResponse_ResultCode.
var (
	ExportResponse_ResultCode_name = map[int32]string{
		0: "Success",
		1: "FailedNoneRetryable",
		2: "FailedRetryable",
	}
	ExportResponse_ResultCode_value = map[string]int32{
		"Success":             0,
		"FailedNoneRetryable": 1,
		"FailedRetryable":     2,
	}
)

func (x ExportResponse_ResultCode) Enum() *ExportResponse_ResultCode {
	p := new(ExportResponse_ResultCode)
	*p = x
	return p
}

func (x ExportResponse_ResultCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExportResponse_ResultCode) Descriptor() protoreflect.EnumDescriptor {
	return file_exchange_proto_enumTypes[2].Descriptor()
}

func (ExportResponse_ResultCode) Type() protoreflect.EnumType {
	return &file_exchange_proto_enumTypes[2]
}

func (x ExportResponse_ResultCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExportResponse_ResultCode.Descriptor instead.
func (ExportResponse_ResultCode) EnumDescriptor() ([]byte, []int) {
	return file_exchange_proto_rawDescGZIP(), []int{3, 0}
}

// Hello is the first request from client to server.
type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exchange_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_exchange_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_exchange_proto_rawDescGZIP(), []int{0}
}

// Response to Hello request.
type HelloResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HelloResponse) Reset() {
	*x = HelloResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exchange_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloResponse) ProtoMessage() {}

func (x *HelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_exchange_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloResponse.ProtoReflect.Descriptor instead.
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return file_exchange_proto_rawDescGZIP(), []int{1}
}

// A request from client to server containing telemetry data to export.
type ExportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique sequential ID generated by the client.
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Telemetry data.
	NodeSpans   []*NodeSpans   `protobuf:"bytes,2,rep,name=nodeSpans,proto3" json:"nodeSpans,omitempty"`
	NodeMetrics []*NodeMetrics `protobuf:"bytes,3,rep,name=nodeMetrics,proto3" json:"nodeMetrics,omitempty"`
}

func (x *ExportRequest) Reset() {
	*x = ExportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exchange_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportRequest) ProtoMessage() {}

func (x *ExportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_exchange_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportRequest.ProtoReflect.Descriptor instead.
func (*ExportRequest) Descriptor() ([]byte, []int) {
	return file_exchange_proto_rawDescGZIP(), []int{2}
}

func (x *ExportRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ExportRequest) GetNodeSpans() []*NodeSpans {
	if x != nil {
		return x.NodeSpans
	}
	return nil
}

func (x *ExportRequest) GetNodeMetrics() []*NodeMetrics {
	if x != nil {
		return x.NodeMetrics
	}
	return nil
}

// A response to ExportRequest.
type ExportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of a response that the server acknowledges.
	Id         uint64                    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ResultCode ExportResponse_ResultCode `protobuf:"varint,2,opt,name=result_code,json=resultCode,proto3,enum=traceprotobuf.ExportResponse_ResultCode" json:"result_code,omitempty"`
}

func (x *ExportResponse) Reset() {
	*x = ExportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exchange_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportResponse) ProtoMessage() {}

func (x *ExportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_exchange_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportResponse.ProtoReflect.Descriptor instead.
func (*ExportResponse) Descriptor() ([]byte, []int) {
	return file_exchange_proto_rawDescGZIP(), []int{3}
}

func (x *ExportResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ExportResponse) GetResultCode() ExportResponse_ResultCode {
	if x != nil {
		return x.ResultCode
	}
	return ExportResponse_Success
}

// RequestHeader is used by transports that unlike gRPC don't have built-in request
// compression such as WebSocket. Request body typically follows the header.
type RequestHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Compression method used for body.
	Compression CompressionMethod `protobuf:"varint,1,opt,name=compression,proto3,enum=traceprotobuf.CompressionMethod" json:"compression,omitempty"`
	// Compression level as defined by the compression method.
	CompressionLevel int32 `protobuf:"varint,2,opt,name=compression_level,json=compressionLevel,proto3" json:"compression_level,omitempty"`
}

func (x *RequestHeader) Reset() {
	*x = RequestHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exchange_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestHeader) ProtoMessage() {}

func (x *RequestHeader) ProtoReflect() protoreflect.Message {
	mi := &file_exchange_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestHeader.ProtoReflect.Descriptor instead.
func (*RequestHeader) Descriptor() ([]byte, []int) {
	return file_exchange_proto_rawDescGZIP(), []int{4}
}

func (x *RequestHeader) GetCompression() CompressionMethod {
	if x != nil {
		return x.Compression
	}
	return CompressionMethod_NONE
}

func (x *RequestHeader) GetCompressionLevel() int32 {
	if x != nil {
		return x.CompressionLevel
	}
	return 0
}

// RequestBody is used by transports that unlike gRPC don't have built-in message type
// multiplexing such as WebSocket.
type RequestBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Body:
	//	*RequestBody_Hello
	//	*RequestBody_Export
	Body isRequestBody_Body `protobuf_oneof:"body"`
}

func (x *RequestBody) Reset() {
	*x = RequestBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exchange_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestBody) ProtoMessage() {}

func (x *RequestBody) ProtoReflect() protoreflect.Message {
	mi := &file_exchange_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestBody.ProtoReflect.Descriptor instead.
func (*RequestBody) Descriptor() ([]byte, []int) {
	return file_exchange_proto_rawDescGZIP(), []int{5}
}

func (m *RequestBody) GetBody() isRequestBody_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (x *RequestBody) GetHello() *HelloRequest {
	if x, ok := x.GetBody().(*RequestBody_Hello); ok {
		return x.Hello
	}
	return nil
}

func (x *RequestBody) GetExport() *ExportRequest {
	if x, ok := x.GetBody().(*RequestBody_Export); ok {
		return x.Export
	}
	return nil
}

type isRequestBody_Body interface {
	isRequestBody_Body()
}

type RequestBody_Hello struct {
	Hello *HelloRequest `protobuf:"bytes,1,opt,name=hello,proto3,oneof"`
}

type RequestBody_Export struct {
	Export *ExportRequest `protobuf:"bytes,2,opt,name=export,proto3,oneof"`
}

func (*RequestBody_Hello) isRequestBody_Body() {}

func (*RequestBody_Export) isRequestBody_Body() {}

// Response is used by transports that unlike gRPC don't have built-in message type
// multiplexing such as WebSocket.
type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Body:
	//	*Response_Hello
	//	*Response_Export
	Body isResponse_Body `protobuf_oneof:"body"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exchange_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_exchange_proto_msgTypes[6]
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
	return file_exchange_proto_rawDescGZIP(), []int{6}
}

func (m *Response) GetBody() isResponse_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (x *Response) GetHello() *HelloResponse {
	if x, ok := x.GetBody().(*Response_Hello); ok {
		return x.Hello
	}
	return nil
}

func (x *Response) GetExport() *ExportResponse {
	if x, ok := x.GetBody().(*Response_Export); ok {
		return x.Export
	}
	return nil
}

type isResponse_Body interface {
	isResponse_Body()
}

type Response_Hello struct {
	Hello *HelloResponse `protobuf:"bytes,1,opt,name=hello,proto3,oneof"`
}

type Response_Export struct {
	Export *ExportResponse `protobuf:"bytes,2,opt,name=export,proto3,oneof"`
}

func (*Response_Hello) isResponse_Body() {}

func (*Response_Export) isResponse_Body() {}

var File_exchange_proto protoreflect.FileDescriptor

var file_exchange_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x74, 0x72, 0x61, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x1a,
	0x14, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0e, 0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x95, 0x01, 0x0a, 0x0d, 0x45, 0x78, 0x70, 0x6f, 0x72,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x36, 0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65,
	0x53, 0x70, 0x61, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x72,
	0x61, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4e, 0x6f, 0x64, 0x65,
	0x53, 0x70, 0x61, 0x6e, 0x73, 0x52, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x53, 0x70, 0x61, 0x6e, 0x73,
	0x12, 0x3c, 0x0a, 0x0b, 0x6e, 0x6f, 0x64, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x52, 0x0b, 0x6e, 0x6f, 0x64, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x22, 0xb4,
	0x01, 0x0a, 0x0e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x49, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x47, 0x0a, 0x0a,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x46, 0x61, 0x69, 0x6c, 0x65,
	0x64, 0x4e, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x74, 0x72, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x10, 0x01,
	0x12, 0x13, 0x0a, 0x0f, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x74, 0x72, 0x79, 0x61,
	0x62, 0x6c, 0x65, 0x10, 0x02, 0x22, 0x80, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x42, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x74,
	0x72, 0x61, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x43, 0x6f, 0x6d,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x0b,
	0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x11, 0x63,
	0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x82, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x33, 0x0a, 0x05, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x36, 0x0a,
	0x06, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x74, 0x72, 0x61, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x78,
	0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x06, 0x65,
	0x78, 0x70, 0x6f, 0x72, 0x74, 0x42, 0x06, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x81, 0x01,
	0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x72, 0x61, 0x63,
	0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x12, 0x37, 0x0a, 0x06, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48,
	0x00, 0x52, 0x06, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x42, 0x06, 0x0a, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x2a, 0x43, 0x0a, 0x0c, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x4b, 0x49, 0x50, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x5a,
	0x4c, 0x49, 0x42, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x52, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x10,
	0x01, 0x12, 0x13, 0x0a, 0x0f, 0x4c, 0x5a, 0x34, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x52, 0x45, 0x53,
	0x53, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x2a, 0x30, 0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e,
	0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x4c, 0x5a, 0x34, 0x10, 0x01, 0x12, 0x08,
	0x0a, 0x04, 0x5a, 0x4c, 0x49, 0x42, 0x10, 0x02, 0x42, 0x48, 0x0a, 0x1c, 0x69, 0x6f, 0x2e, 0x6f,
	0x70, 0x65, 0x6e, 0x63, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x54, 0x72, 0x61, 0x63, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0xea, 0x02, 0x19, 0x4f, 0x70, 0x65, 0x6e, 0x43, 0x65, 0x6e,
	0x73, 0x75, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x65, 0x2e,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_exchange_proto_rawDescOnce sync.Once
	file_exchange_proto_rawDescData = file_exchange_proto_rawDesc
)

func file_exchange_proto_rawDescGZIP() []byte {
	file_exchange_proto_rawDescOnce.Do(func() {
		file_exchange_proto_rawDescData = protoimpl.X.CompressGZIP(file_exchange_proto_rawDescData)
	})
	return file_exchange_proto_rawDescData
}

var file_exchange_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_exchange_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_exchange_proto_goTypes = []interface{}{
	(Capabilities)(0),              // 0: traceprotobuf.Capabilities
	(CompressionMethod)(0),         // 1: traceprotobuf.CompressionMethod
	(ExportResponse_ResultCode)(0), // 2: traceprotobuf.ExportResponse.ResultCode
	(*HelloRequest)(nil),           // 3: traceprotobuf.HelloRequest
	(*HelloResponse)(nil),          // 4: traceprotobuf.HelloResponse
	(*ExportRequest)(nil),          // 5: traceprotobuf.ExportRequest
	(*ExportResponse)(nil),         // 6: traceprotobuf.ExportResponse
	(*RequestHeader)(nil),          // 7: traceprotobuf.RequestHeader
	(*RequestBody)(nil),            // 8: traceprotobuf.RequestBody
	(*Response)(nil),               // 9: traceprotobuf.Response
	(*NodeSpans)(nil),              // 10: traceprotobuf.NodeSpans
	(*NodeMetrics)(nil),            // 11: traceprotobuf.NodeMetrics
}
var file_exchange_proto_depIdxs = []int32{
	10, // 0: traceprotobuf.ExportRequest.nodeSpans:type_name -> traceprotobuf.NodeSpans
	11, // 1: traceprotobuf.ExportRequest.nodeMetrics:type_name -> traceprotobuf.NodeMetrics
	2,  // 2: traceprotobuf.ExportResponse.result_code:type_name -> traceprotobuf.ExportResponse.ResultCode
	1,  // 3: traceprotobuf.RequestHeader.compression:type_name -> traceprotobuf.CompressionMethod
	3,  // 4: traceprotobuf.RequestBody.hello:type_name -> traceprotobuf.HelloRequest
	5,  // 5: traceprotobuf.RequestBody.export:type_name -> traceprotobuf.ExportRequest
	4,  // 6: traceprotobuf.Response.hello:type_name -> traceprotobuf.HelloResponse
	6,  // 7: traceprotobuf.Response.export:type_name -> traceprotobuf.ExportResponse
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_exchange_proto_init() }
func file_exchange_proto_init() {
	if File_exchange_proto != nil {
		return
	}
	file_telemetry_data_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_exchange_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_exchange_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloResponse); i {
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
		file_exchange_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExportRequest); i {
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
		file_exchange_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExportResponse); i {
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
		file_exchange_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestHeader); i {
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
		file_exchange_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestBody); i {
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
		file_exchange_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
	}
	file_exchange_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*RequestBody_Hello)(nil),
		(*RequestBody_Export)(nil),
	}
	file_exchange_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*Response_Hello)(nil),
		(*Response_Export)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_exchange_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_exchange_proto_goTypes,
		DependencyIndexes: file_exchange_proto_depIdxs,
		EnumInfos:         file_exchange_proto_enumTypes,
		MessageInfos:      file_exchange_proto_msgTypes,
	}.Build()
	File_exchange_proto = out.File
	file_exchange_proto_rawDesc = nil
	file_exchange_proto_goTypes = nil
	file_exchange_proto_depIdxs = nil
}
