// Code generated by protoc-gen-go. DO NOT EDIT.
// source: exchange.proto

package experimental

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

type Capabilities int32

const (
	Capabilities_SKIP             Capabilities = 0
	Capabilities_ZLIB_COMPRESSION Capabilities = 1
	Capabilities_LZ4_COMPRESSION  Capabilities = 2
)

var Capabilities_name = map[int32]string{
	0: "SKIP",
	1: "ZLIB_COMPRESSION",
	2: "LZ4_COMPRESSION",
}

var Capabilities_value = map[string]int32{
	"SKIP":             0,
	"ZLIB_COMPRESSION": 1,
	"LZ4_COMPRESSION":  2,
}

func (x Capabilities) String() string {
	return proto.EnumName(Capabilities_name, int32(x))
}

func (Capabilities) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e0328a4f16f87ea1, []int{0}
}

type CompressionMethod int32

const (
	CompressionMethod_NONE CompressionMethod = 0
	CompressionMethod_LZ4  CompressionMethod = 1
	CompressionMethod_ZLIB CompressionMethod = 2
)

var CompressionMethod_name = map[int32]string{
	0: "NONE",
	1: "LZ4",
	2: "ZLIB",
}

var CompressionMethod_value = map[string]int32{
	"NONE": 0,
	"LZ4":  1,
	"ZLIB": 2,
}

func (x CompressionMethod) String() string {
	return proto.EnumName(CompressionMethod_name, int32(x))
}

func (CompressionMethod) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e0328a4f16f87ea1, []int{1}
}

type RequestType int32

const (
	RequestType__           RequestType = 0
	RequestType_TraceExport RequestType = 1
)

var RequestType_name = map[int32]string{
	0: "_",
	1: "TraceExport",
}

var RequestType_value = map[string]int32{
	"_":           0,
	"TraceExport": 1,
}

func (x RequestType) String() string {
	return proto.EnumName(RequestType_name, int32(x))
}

func (RequestType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e0328a4f16f87ea1, []int{2}
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

var ExportResponse_ResultCode_name = map[int32]string{
	0: "Success",
	1: "FailedNoneRetryable",
	2: "FailedRetryable",
}

var ExportResponse_ResultCode_value = map[string]int32{
	"Success":             0,
	"FailedNoneRetryable": 1,
	"FailedRetryable":     2,
}

func (x ExportResponse_ResultCode) String() string {
	return proto.EnumName(ExportResponse_ResultCode_name, int32(x))
}

func (ExportResponse_ResultCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e0328a4f16f87ea1, []int{3, 0}
}

// A request from client to server containing trace data to export.
type TraceExportRequest struct {
	// Unique sequential ID generated by the client.
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Telemetry data. An array of ResourceSpans.
	ResourceSpans        []*ResourceSpans `protobuf:"bytes,2,rep,name=resourceSpans,proto3" json:"resourceSpans,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TraceExportRequest) Reset()         { *m = TraceExportRequest{} }
func (m *TraceExportRequest) String() string { return proto.CompactTextString(m) }
func (*TraceExportRequest) ProtoMessage()    {}
func (*TraceExportRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0328a4f16f87ea1, []int{0}
}

func (m *TraceExportRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TraceExportRequest.Unmarshal(m, b)
}
func (m *TraceExportRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TraceExportRequest.Marshal(b, m, deterministic)
}
func (m *TraceExportRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TraceExportRequest.Merge(m, src)
}
func (m *TraceExportRequest) XXX_Size() int {
	return xxx_messageInfo_TraceExportRequest.Size(m)
}
func (m *TraceExportRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TraceExportRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TraceExportRequest proto.InternalMessageInfo

func (m *TraceExportRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TraceExportRequest) GetResourceSpans() []*ResourceSpans {
	if m != nil {
		return m.ResourceSpans
	}
	return nil
}

// A request from client to server containing trace data to export.
type TraceExportRequestPrepared struct {
	// Unique sequential ID generated by the client.
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Telemetry data. An array of ResourceSpans.
	ResourceSpans        []*ResourceSpansPrepared `protobuf:"bytes,2,rep,name=resourceSpans,proto3" json:"resourceSpans,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *TraceExportRequestPrepared) Reset()         { *m = TraceExportRequestPrepared{} }
func (m *TraceExportRequestPrepared) String() string { return proto.CompactTextString(m) }
func (*TraceExportRequestPrepared) ProtoMessage()    {}
func (*TraceExportRequestPrepared) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0328a4f16f87ea1, []int{1}
}

func (m *TraceExportRequestPrepared) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TraceExportRequestPrepared.Unmarshal(m, b)
}
func (m *TraceExportRequestPrepared) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TraceExportRequestPrepared.Marshal(b, m, deterministic)
}
func (m *TraceExportRequestPrepared) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TraceExportRequestPrepared.Merge(m, src)
}
func (m *TraceExportRequestPrepared) XXX_Size() int {
	return xxx_messageInfo_TraceExportRequestPrepared.Size(m)
}
func (m *TraceExportRequestPrepared) XXX_DiscardUnknown() {
	xxx_messageInfo_TraceExportRequestPrepared.DiscardUnknown(m)
}

var xxx_messageInfo_TraceExportRequestPrepared proto.InternalMessageInfo

func (m *TraceExportRequestPrepared) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TraceExportRequestPrepared) GetResourceSpans() []*ResourceSpansPrepared {
	if m != nil {
		return m.ResourceSpans
	}
	return nil
}

// A request from client to server containing metric data to export.
type MetricExportRequest struct {
	// Unique sequential ID generated by the client.
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Telemetry data. An array of ResourceMetrics.
	ResourceMetrics      []*ResourceMetrics `protobuf:"bytes,2,rep,name=resourceMetrics,proto3" json:"resourceMetrics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *MetricExportRequest) Reset()         { *m = MetricExportRequest{} }
func (m *MetricExportRequest) String() string { return proto.CompactTextString(m) }
func (*MetricExportRequest) ProtoMessage()    {}
func (*MetricExportRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0328a4f16f87ea1, []int{2}
}

func (m *MetricExportRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricExportRequest.Unmarshal(m, b)
}
func (m *MetricExportRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricExportRequest.Marshal(b, m, deterministic)
}
func (m *MetricExportRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricExportRequest.Merge(m, src)
}
func (m *MetricExportRequest) XXX_Size() int {
	return xxx_messageInfo_MetricExportRequest.Size(m)
}
func (m *MetricExportRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricExportRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MetricExportRequest proto.InternalMessageInfo

func (m *MetricExportRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MetricExportRequest) GetResourceMetrics() []*ResourceMetrics {
	if m != nil {
		return m.ResourceMetrics
	}
	return nil
}

// A response to ExportRequest.
type ExportResponse struct {
	// ID of a response that the server acknowledges.
	Id                   uint64                    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ResultCode           ExportResponse_ResultCode `protobuf:"varint,2,opt,name=result_code,json=resultCode,proto3,enum=experimental.ExportResponse_ResultCode" json:"result_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ExportResponse) Reset()         { *m = ExportResponse{} }
func (m *ExportResponse) String() string { return proto.CompactTextString(m) }
func (*ExportResponse) ProtoMessage()    {}
func (*ExportResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0328a4f16f87ea1, []int{3}
}

func (m *ExportResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportResponse.Unmarshal(m, b)
}
func (m *ExportResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportResponse.Marshal(b, m, deterministic)
}
func (m *ExportResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportResponse.Merge(m, src)
}
func (m *ExportResponse) XXX_Size() int {
	return xxx_messageInfo_ExportResponse.Size(m)
}
func (m *ExportResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExportResponse proto.InternalMessageInfo

func (m *ExportResponse) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ExportResponse) GetResultCode() ExportResponse_ResultCode {
	if m != nil {
		return m.ResultCode
	}
	return ExportResponse_Success
}

// RequestHeader is used by transports that unlike gRPC don't have built-in request
// compression such as WebSocket. Request body typically follows the header.
type RequestHeader struct {
	// Compression method used for body.
	Compression CompressionMethod `protobuf:"varint,1,opt,name=compression,proto3,enum=experimental.CompressionMethod" json:"compression,omitempty"`
	// Compression level as defined by the compression method.
	CompressionLevel     int32    `protobuf:"varint,2,opt,name=compression_level,json=compressionLevel,proto3" json:"compression_level,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestHeader) Reset()         { *m = RequestHeader{} }
func (m *RequestHeader) String() string { return proto.CompactTextString(m) }
func (*RequestHeader) ProtoMessage()    {}
func (*RequestHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0328a4f16f87ea1, []int{4}
}

func (m *RequestHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestHeader.Unmarshal(m, b)
}
func (m *RequestHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestHeader.Marshal(b, m, deterministic)
}
func (m *RequestHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestHeader.Merge(m, src)
}
func (m *RequestHeader) XXX_Size() int {
	return xxx_messageInfo_RequestHeader.Size(m)
}
func (m *RequestHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestHeader.DiscardUnknown(m)
}

var xxx_messageInfo_RequestHeader proto.InternalMessageInfo

func (m *RequestHeader) GetCompression() CompressionMethod {
	if m != nil {
		return m.Compression
	}
	return CompressionMethod_NONE
}

func (m *RequestHeader) GetCompressionLevel() int32 {
	if m != nil {
		return m.CompressionLevel
	}
	return 0
}

// RequestBody is used by transports that unlike gRPC don't have built-in message type
// multiplexing such as WebSocket.
type RequestBody struct {
	RequestType          RequestType         `protobuf:"varint,1,opt,name=request_type,json=requestType,proto3,enum=experimental.RequestType" json:"request_type,omitempty"`
	Export               *TraceExportRequest `protobuf:"bytes,2,opt,name=export,proto3" json:"export,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *RequestBody) Reset()         { *m = RequestBody{} }
func (m *RequestBody) String() string { return proto.CompactTextString(m) }
func (*RequestBody) ProtoMessage()    {}
func (*RequestBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0328a4f16f87ea1, []int{5}
}

func (m *RequestBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestBody.Unmarshal(m, b)
}
func (m *RequestBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestBody.Marshal(b, m, deterministic)
}
func (m *RequestBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestBody.Merge(m, src)
}
func (m *RequestBody) XXX_Size() int {
	return xxx_messageInfo_RequestBody.Size(m)
}
func (m *RequestBody) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestBody.DiscardUnknown(m)
}

var xxx_messageInfo_RequestBody proto.InternalMessageInfo

func (m *RequestBody) GetRequestType() RequestType {
	if m != nil {
		return m.RequestType
	}
	return RequestType__
}

func (m *RequestBody) GetExport() *TraceExportRequest {
	if m != nil {
		return m.Export
	}
	return nil
}

// Response is used by transports that unlike gRPC don't have built-in message type
// multiplexing such as WebSocket.
type Response struct {
	ResponseType         RequestType     `protobuf:"varint,1,opt,name=response_type,json=responseType,proto3,enum=experimental.RequestType" json:"response_type,omitempty"`
	Export               *ExportResponse `protobuf:"bytes,2,opt,name=export,proto3" json:"export,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0328a4f16f87ea1, []int{6}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetResponseType() RequestType {
	if m != nil {
		return m.ResponseType
	}
	return RequestType__
}

func (m *Response) GetExport() *ExportResponse {
	if m != nil {
		return m.Export
	}
	return nil
}

func init() {
	proto.RegisterEnum("experimental.Capabilities", Capabilities_name, Capabilities_value)
	proto.RegisterEnum("experimental.CompressionMethod", CompressionMethod_name, CompressionMethod_value)
	proto.RegisterEnum("experimental.RequestType", RequestType_name, RequestType_value)
	proto.RegisterEnum("experimental.ExportResponse_ResultCode", ExportResponse_ResultCode_name, ExportResponse_ResultCode_value)
	proto.RegisterType((*TraceExportRequest)(nil), "experimental.TraceExportRequest")
	proto.RegisterType((*TraceExportRequestPrepared)(nil), "experimental.TraceExportRequestPrepared")
	proto.RegisterType((*MetricExportRequest)(nil), "experimental.MetricExportRequest")
	proto.RegisterType((*ExportResponse)(nil), "experimental.ExportResponse")
	proto.RegisterType((*RequestHeader)(nil), "experimental.RequestHeader")
	proto.RegisterType((*RequestBody)(nil), "experimental.RequestBody")
	proto.RegisterType((*Response)(nil), "experimental.Response")
}

func init() { proto.RegisterFile("exchange.proto", fileDescriptor_e0328a4f16f87ea1) }

var fileDescriptor_e0328a4f16f87ea1 = []byte{
	// 586 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0x5f, 0x4f, 0xdb, 0x3c,
	0x14, 0xc6, 0x9b, 0xf0, 0x57, 0x27, 0xa5, 0x04, 0x83, 0xf4, 0x02, 0x2f, 0xd3, 0xaa, 0x4c, 0xd3,
	0x10, 0x93, 0xa2, 0xd1, 0x71, 0xb1, 0x8b, 0x69, 0x12, 0xad, 0x18, 0x54, 0x2b, 0x6d, 0xe5, 0xa2,
	0x5d, 0x70, 0x53, 0x99, 0xe4, 0x08, 0x22, 0x85, 0xd8, 0xb3, 0x5d, 0x46, 0xaf, 0xb6, 0x9b, 0x7d,
	0xa2, 0x7d, 0xab, 0x7d, 0x8a, 0x29, 0x4e, 0x43, 0x93, 0x16, 0x55, 0xbb, 0x8b, 0x1f, 0x3f, 0xc7,
	0xbf, 0xc7, 0xe7, 0x58, 0x81, 0x1a, 0x3e, 0x06, 0x77, 0x2c, 0xb9, 0x45, 0x5f, 0x48, 0xae, 0x39,
	0xa9, 0xe2, 0xa3, 0x40, 0x19, 0xdd, 0x63, 0xa2, 0x59, 0xbc, 0xbf, 0xa3, 0x31, 0xc6, 0x7b, 0xd4,
	0x72, 0x3c, 0x0c, 0x99, 0x66, 0x99, 0x67, 0x7f, 0x2b, 0x55, 0xa2, 0xa0, 0x20, 0x79, 0xb7, 0x40,
	0xae, 0x24, 0x0b, 0xf0, 0xec, 0x51, 0x70, 0xa9, 0x29, 0x7e, 0x1b, 0xa1, 0xd2, 0xa4, 0x06, 0x76,
	0x14, 0xee, 0x5a, 0x75, 0xeb, 0x70, 0x99, 0xda, 0x51, 0x48, 0x4e, 0x61, 0x43, 0xa2, 0xe2, 0x23,
	0x19, 0xe0, 0x40, 0xb0, 0x44, 0xed, 0xda, 0xf5, 0xa5, 0x43, 0xa7, 0xf1, 0xbf, 0x5f, 0x84, 0xfa,
	0xb4, 0x68, 0xa1, 0xe5, 0x0a, 0xef, 0x3b, 0xec, 0xcf, 0x83, 0xfa, 0x12, 0x05, 0x93, 0x18, 0xce,
	0x01, 0xdb, 0xcf, 0x03, 0x5f, 0x2d, 0x00, 0xe6, 0x67, 0xcd, 0x82, 0x13, 0xd8, 0xbe, 0x34, 0xd7,
	0x5e, 0x7c, 0xc5, 0x73, 0xd8, 0xcc, 0xeb, 0x32, 0x7b, 0xce, 0x7c, 0xf1, 0x3c, 0x73, 0x62, 0xa2,
	0xb3, 0x55, 0xde, 0x6f, 0x0b, 0x6a, 0x39, 0x4a, 0x09, 0x9e, 0x28, 0x9c, 0x63, 0x5d, 0x80, 0x23,
	0x51, 0x8d, 0x62, 0x3d, 0x0c, 0x78, 0x88, 0xbb, 0x76, 0xdd, 0x3a, 0xac, 0x35, 0xde, 0x94, 0x39,
	0xe5, 0x23, 0x52, 0xec, 0x28, 0xd6, 0x2d, 0x1e, 0x22, 0x05, 0xf9, 0xf4, 0xed, 0x9d, 0x03, 0x4c,
	0x77, 0x88, 0x03, 0x6b, 0x83, 0x51, 0x10, 0xa0, 0x52, 0x6e, 0x85, 0xfc, 0x07, 0xdb, 0x9f, 0x59,
	0x14, 0x63, 0xd8, 0xe5, 0x09, 0xd2, 0xf4, 0x29, 0xb0, 0x9b, 0x18, 0x5d, 0x8b, 0x6c, 0xc3, 0x66,
	0xb6, 0x31, 0x15, 0x6d, 0xef, 0x07, 0x6c, 0x4c, 0x3a, 0x73, 0x81, 0x2c, 0x44, 0x49, 0x4e, 0xc1,
	0x09, 0xf8, 0xbd, 0x90, 0xa8, 0x54, 0xc4, 0x13, 0x13, 0xbe, 0xd6, 0x78, 0x59, 0xce, 0xd8, 0x9a,
	0x1a, 0x2e, 0x51, 0xdf, 0xf1, 0x90, 0x16, 0x6b, 0xc8, 0x5b, 0xd8, 0x2a, 0x2c, 0x87, 0x31, 0x3e,
	0x60, 0x6c, 0x2e, 0xbb, 0x42, 0xdd, 0xc2, 0x46, 0x27, 0xd5, 0xbd, 0x5f, 0x16, 0x38, 0x93, 0x04,
	0x4d, 0x1e, 0x8e, 0xc9, 0x47, 0xa8, 0xca, 0x6c, 0x39, 0xd4, 0x63, 0x81, 0x93, 0x00, 0x7b, 0xb3,
	0xc3, 0x30, 0x8e, 0xab, 0xb1, 0x40, 0xea, 0xc8, 0xe9, 0x82, 0x7c, 0x80, 0x55, 0x34, 0x0d, 0x34,
	0x3c, 0xa7, 0x51, 0x2f, 0xd7, 0xcd, 0xbf, 0x44, 0x3a, 0xf1, 0x7b, 0x3f, 0x2d, 0x58, 0x7f, 0x1a,
	0xdc, 0x27, 0xf3, 0x0c, 0xcd, 0xf7, 0x3f, 0xa6, 0xa8, 0xe6, 0x7e, 0x13, 0xe3, 0x64, 0x26, 0xc6,
	0xc1, 0xa2, 0x19, 0xe7, 0x11, 0x8e, 0x5a, 0x50, 0x6d, 0x31, 0xc1, 0x6e, 0xa2, 0x38, 0xd2, 0x11,
	0x2a, 0xb2, 0x0e, 0xcb, 0x83, 0x2f, 0xed, 0xbe, 0x5b, 0x21, 0x3b, 0xe0, 0x5e, 0x77, 0xda, 0xcd,
	0x61, 0xab, 0x77, 0xd9, 0xa7, 0x67, 0x83, 0x41, 0xbb, 0xd7, 0xcd, 0x06, 0xda, 0xb9, 0x3e, 0x29,
	0x89, 0xf6, 0xd1, 0x3b, 0xd8, 0x9a, 0x1b, 0x4f, 0x7a, 0x52, 0xb7, 0xd7, 0x3d, 0x73, 0x2b, 0x64,
	0x0d, 0x96, 0x3a, 0xd7, 0x27, 0xae, 0x95, 0x4a, 0xe9, 0x91, 0xae, 0x7d, 0xf4, 0xfa, 0x69, 0x00,
	0x26, 0xfb, 0x0a, 0x58, 0x43, 0xb7, 0x42, 0x36, 0xc1, 0x29, 0x74, 0xcb, 0xb5, 0x9a, 0x17, 0x70,
	0x10, 0x71, 0x9f, 0x0b, 0x4c, 0x02, 0x4c, 0xd4, 0x48, 0x65, 0x3f, 0x12, 0x5f, 0xa7, 0x26, 0xff,
	0xe1, 0xb8, 0x09, 0xc6, 0xde, 0x4f, 0xc5, 0xbe, 0xf5, 0xc7, 0xde, 0xeb, 0x09, 0x4c, 0x5a, 0x99,
	0xd3, 0x88, 0x59, 0xf3, 0xfd, 0xaf, 0xc7, 0x37, 0xab, 0xa6, 0xf2, 0xfd, 0xdf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x6d, 0x6c, 0xee, 0x25, 0xcb, 0x04, 0x00, 0x00,
}
