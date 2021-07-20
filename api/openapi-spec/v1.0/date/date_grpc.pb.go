// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package date

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DateServiceClient is the client API for DateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
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
// All implementations must embed UnimplementedDateServiceServer
// for forward compatibility
type DateServiceServer interface {
	Now(context.Context, *DateRequest) (*DateResponse, error)
	mustEmbedUnimplementedDateServiceServer()
}

// UnimplementedDateServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDateServiceServer struct {
}

func (UnimplementedDateServiceServer) Now(context.Context, *DateRequest) (*DateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Now not implemented")
}
func (UnimplementedDateServiceServer) mustEmbedUnimplementedDateServiceServer() {}

// UnsafeDateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DateServiceServer will
// result in compilation errors.
type UnsafeDateServiceServer interface {
	mustEmbedUnimplementedDateServiceServer()
}

func RegisterDateServiceServer(s grpc.ServiceRegistrar, srv DateServiceServer) {
	s.RegisterService(&DateService_ServiceDesc, srv)
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

// DateService_ServiceDesc is the grpc.ServiceDesc for DateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DateService_ServiceDesc = grpc.ServiceDesc{
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
