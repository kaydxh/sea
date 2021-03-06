// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/openapi-spec/v1.0/date/date.proto

package date

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/protoc-gen-go/descriptor"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "github.com/mwitkow/go-proto-validators"
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

type DateRequest struct {
	RequestId            string   `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DateRequest) Reset()         { *m = DateRequest{} }
func (m *DateRequest) String() string { return proto.CompactTextString(m) }
func (*DateRequest) ProtoMessage()    {}
func (*DateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac80638d97e7358d, []int{0}
}

func (m *DateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DateRequest.Unmarshal(m, b)
}
func (m *DateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DateRequest.Marshal(b, m, deterministic)
}
func (m *DateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DateRequest.Merge(m, src)
}
func (m *DateRequest) XXX_Size() int {
	return xxx_messageInfo_DateRequest.Size(m)
}
func (m *DateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DateRequest proto.InternalMessageInfo

func (m *DateRequest) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

type DateResponse struct {
	RequestId            string   `protobuf:"bytes,1,opt,name=request_id,json=RequestId,proto3" json:"request_id,omitempty"`
	Date                 string   `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DateResponse) Reset()         { *m = DateResponse{} }
func (m *DateResponse) String() string { return proto.CompactTextString(m) }
func (*DateResponse) ProtoMessage()    {}
func (*DateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac80638d97e7358d, []int{1}
}

func (m *DateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DateResponse.Unmarshal(m, b)
}
func (m *DateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DateResponse.Marshal(b, m, deterministic)
}
func (m *DateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DateResponse.Merge(m, src)
}
func (m *DateResponse) XXX_Size() int {
	return xxx_messageInfo_DateResponse.Size(m)
}
func (m *DateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DateResponse proto.InternalMessageInfo

func (m *DateResponse) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *DateResponse) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func init() {
	proto.RegisterType((*DateRequest)(nil), "sea.api.v1.date.DateRequest")
	proto.RegisterType((*DateResponse)(nil), "sea.api.v1.date.DateResponse")
}

func init() {
	proto.RegisterFile("api/openapi-spec/v1.0/date/date.proto", fileDescriptor_ac80638d97e7358d)
}

var fileDescriptor_ac80638d97e7358d = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x4f, 0xeb, 0x13, 0x31,
	0x10, 0x75, 0xab, 0x14, 0x1a, 0x05, 0x65, 0x4f, 0xa5, 0x58, 0x5c, 0x0a, 0x05, 0x0f, 0x6e, 0x62,
	0x2b, 0x78, 0xa8, 0x78, 0xb0, 0x78, 0xe9, 0xc5, 0xc3, 0xf6, 0x22, 0x5e, 0x64, 0x9a, 0x9d, 0xa6,
	0xb1, 0x6d, 0x26, 0x4d, 0xb2, 0xbb, 0xee, 0xd1, 0x8f, 0xa0, 0x5f, 0x50, 0xf0, 0x93, 0x48, 0xb3,
	0x55, 0x4a, 0x7f, 0xfc, 0x2e, 0xc3, 0x9b, 0x3f, 0xef, 0x65, 0x78, 0x19, 0x36, 0x05, 0xab, 0x05,
	0x59, 0x34, 0x60, 0x75, 0xee, 0x2d, 0x4a, 0x51, 0xcf, 0xf8, 0x6b, 0x51, 0x42, 0xc0, 0x18, 0xb8,
	0x75, 0x14, 0x28, 0x7d, 0xea, 0x11, 0x38, 0x58, 0xcd, 0xeb, 0x19, 0x3f, 0x97, 0x47, 0x99, 0x22,
	0x52, 0x07, 0x14, 0xb1, 0xbd, 0xa9, 0xb6, 0xa2, 0x44, 0x2f, 0x9d, 0xb6, 0x81, 0x5c, 0x47, 0x19,
	0xad, 0x95, 0x0e, 0xbb, 0x6a, 0xc3, 0x25, 0x1d, 0x85, 0x72, 0x56, 0xe6, 0x28, 0xc9, 0xb7, 0x3e,
	0xe0, 0x25, 0x55, 0x10, 0xb0, 0x81, 0xb6, 0x53, 0x90, 0xb9, 0x42, 0x93, 0xfb, 0x06, 0x94, 0x42,
	0x27, 0xc8, 0x06, 0x4d, 0xc6, 0x0b, 0x30, 0x86, 0x02, 0x44, 0x7c, 0x11, 0x7d, 0x7b, 0x25, 0x7a,
	0x6c, 0x74, 0xd8, 0x53, 0x23, 0x14, 0xe5, 0xb1, 0x99, 0xd7, 0x70, 0xd0, 0x25, 0x04, 0x72, 0x5e,
	0xfc, 0x87, 0x1d, 0x6f, 0xf2, 0x8a, 0x3d, 0xfe, 0x08, 0x01, 0x0b, 0x3c, 0x55, 0xe8, 0x43, 0x3a,
	0x66, 0xcc, 0x75, 0xf0, 0xab, 0x2e, 0x87, 0x49, 0x96, 0xbc, 0x1c, 0x14, 0x83, 0x4b, 0x65, 0x55,
	0x4e, 0x56, 0xec, 0x49, 0x37, 0xed, 0x2d, 0x19, 0x8f, 0xe9, 0xf4, 0xee, 0xf8, 0xb2, 0xff, 0xe7,
	0xf7, 0x8b, 0xde, 0xe7, 0xa4, 0x18, 0x14, 0xff, 0x68, 0x69, 0xca, 0x1e, 0x9d, 0xbd, 0x19, 0xf6,
	0xa2, 0x5e, 0xc4, 0xf3, 0x9f, 0x49, 0xf7, 0xf2, 0x1a, 0x5d, 0xad, 0x25, 0xa6, 0x3f, 0x12, 0xf6,
	0xf0, 0x13, 0x35, 0xe9, 0x73, 0x7e, 0xe3, 0x28, 0xbf, 0xda, 0x6f, 0x34, 0xbe, 0xa7, 0xdb, 0xed,
	0x33, 0x79, 0xff, 0xeb, 0xc3, 0x82, 0xb1, 0x53, 0x85, 0xae, 0xcd, 0xe2, 0x6f, 0x0c, 0x1d, 0x86,
	0xca, 0x99, 0x98, 0x64, 0xb4, 0xcd, 0xc2, 0x0e, 0x33, 0x43, 0x25, 0xb2, 0xf9, 0x83, 0xc5, 0x33,
	0xb0, 0xf6, 0xa0, 0x65, 0xb4, 0x51, 0x7c, 0xf3, 0x64, 0x96, 0xb3, 0x2f, 0xe2, 0xca, 0xc6, 0x3d,
	0xb4, 0xe5, 0xf7, 0x9d, 0xf0, 0x08, 0xb7, 0x77, 0x10, 0x0f, 0xe0, 0xdd, 0x39, 0x6c, 0xfa, 0xd1,
	0xc6, 0x37, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x0f, 0x72, 0x31, 0x22, 0x2f, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DateServiceClient is the client API for DateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DateServiceClient interface {
	Now(ctx context.Context, in *DateRequest, opts ...grpc.CallOption) (*DateResponse, error)
}

type dateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDateServiceClient(cc grpc.ClientConnInterface) DateServiceClient {
	return &dateServiceClient{cc}
}

func (c *dateServiceClient) Now(ctx context.Context, in *DateRequest, opts ...grpc.CallOption) (*DateResponse, error) {
	out := new(DateResponse)
	err := c.cc.Invoke(ctx, "/sea.api.v1.date.DateService/Now", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DateServiceServer is the server API for DateService service.
type DateServiceServer interface {
	Now(context.Context, *DateRequest) (*DateResponse, error)
}

// UnimplementedDateServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDateServiceServer struct {
}

func (*UnimplementedDateServiceServer) Now(ctx context.Context, req *DateRequest) (*DateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Now not implemented")
}

func RegisterDateServiceServer(s *grpc.Server, srv DateServiceServer) {
	s.RegisterService(&_DateService_serviceDesc, srv)
}

func _DateService_Now_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DateServiceServer).Now(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sea.api.v1.date.DateService/Now",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DateServiceServer).Now(ctx, req.(*DateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DateService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sea.api.v1.date.DateService",
	HandlerType: (*DateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Now",
			Handler:    _DateService_Now_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/openapi-spec/v1.0/date/date.proto",
}
