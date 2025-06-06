// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: examples/internal/proto/examplepb/excess_body.proto

// Excess Body Service
// Used to test server context cancellation with Unary and ServerStream methods
// when client sends more data than expected.

package examplepb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ExcessBodyService_NoBodyRpc_FullMethodName            = "/grpc.gateway.examples.internal.proto.examplepb.ExcessBodyService/NoBodyRpc"
	ExcessBodyService_NoBodyServerStream_FullMethodName   = "/grpc.gateway.examples.internal.proto.examplepb.ExcessBodyService/NoBodyServerStream"
	ExcessBodyService_WithBodyRpc_FullMethodName          = "/grpc.gateway.examples.internal.proto.examplepb.ExcessBodyService/WithBodyRpc"
	ExcessBodyService_WithBodyServerStream_FullMethodName = "/grpc.gateway.examples.internal.proto.examplepb.ExcessBodyService/WithBodyServerStream"
)

// ExcessBodyServiceClient is the client API for ExcessBodyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExcessBodyServiceClient interface {
	NoBodyRpc(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	NoBodyServerStream(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (grpc.ServerStreamingClient[emptypb.Empty], error)
	WithBodyRpc(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	WithBodyServerStream(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (grpc.ServerStreamingClient[emptypb.Empty], error)
}

type excessBodyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExcessBodyServiceClient(cc grpc.ClientConnInterface) ExcessBodyServiceClient {
	return &excessBodyServiceClient{cc}
}

func (c *excessBodyServiceClient) NoBodyRpc(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ExcessBodyService_NoBodyRpc_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *excessBodyServiceClient) NoBodyServerStream(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (grpc.ServerStreamingClient[emptypb.Empty], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ExcessBodyService_ServiceDesc.Streams[0], ExcessBodyService_NoBodyServerStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[emptypb.Empty, emptypb.Empty]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ExcessBodyService_NoBodyServerStreamClient = grpc.ServerStreamingClient[emptypb.Empty]

func (c *excessBodyServiceClient) WithBodyRpc(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ExcessBodyService_WithBodyRpc_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *excessBodyServiceClient) WithBodyServerStream(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (grpc.ServerStreamingClient[emptypb.Empty], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ExcessBodyService_ServiceDesc.Streams[1], ExcessBodyService_WithBodyServerStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[emptypb.Empty, emptypb.Empty]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ExcessBodyService_WithBodyServerStreamClient = grpc.ServerStreamingClient[emptypb.Empty]

// ExcessBodyServiceServer is the server API for ExcessBodyService service.
// All implementations should embed UnimplementedExcessBodyServiceServer
// for forward compatibility.
type ExcessBodyServiceServer interface {
	NoBodyRpc(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	NoBodyServerStream(*emptypb.Empty, grpc.ServerStreamingServer[emptypb.Empty]) error
	WithBodyRpc(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	WithBodyServerStream(*emptypb.Empty, grpc.ServerStreamingServer[emptypb.Empty]) error
}

// UnimplementedExcessBodyServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedExcessBodyServiceServer struct{}

func (UnimplementedExcessBodyServiceServer) NoBodyRpc(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NoBodyRpc not implemented")
}
func (UnimplementedExcessBodyServiceServer) NoBodyServerStream(*emptypb.Empty, grpc.ServerStreamingServer[emptypb.Empty]) error {
	return status.Errorf(codes.Unimplemented, "method NoBodyServerStream not implemented")
}
func (UnimplementedExcessBodyServiceServer) WithBodyRpc(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WithBodyRpc not implemented")
}
func (UnimplementedExcessBodyServiceServer) WithBodyServerStream(*emptypb.Empty, grpc.ServerStreamingServer[emptypb.Empty]) error {
	return status.Errorf(codes.Unimplemented, "method WithBodyServerStream not implemented")
}
func (UnimplementedExcessBodyServiceServer) testEmbeddedByValue() {}

// UnsafeExcessBodyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExcessBodyServiceServer will
// result in compilation errors.
type UnsafeExcessBodyServiceServer interface {
	mustEmbedUnimplementedExcessBodyServiceServer()
}

func RegisterExcessBodyServiceServer(s grpc.ServiceRegistrar, srv ExcessBodyServiceServer) {
	// If the following call pancis, it indicates UnimplementedExcessBodyServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ExcessBodyService_ServiceDesc, srv)
}

func _ExcessBodyService_NoBodyRpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExcessBodyServiceServer).NoBodyRpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExcessBodyService_NoBodyRpc_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExcessBodyServiceServer).NoBodyRpc(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExcessBodyService_NoBodyServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExcessBodyServiceServer).NoBodyServerStream(m, &grpc.GenericServerStream[emptypb.Empty, emptypb.Empty]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ExcessBodyService_NoBodyServerStreamServer = grpc.ServerStreamingServer[emptypb.Empty]

func _ExcessBodyService_WithBodyRpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExcessBodyServiceServer).WithBodyRpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExcessBodyService_WithBodyRpc_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExcessBodyServiceServer).WithBodyRpc(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExcessBodyService_WithBodyServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExcessBodyServiceServer).WithBodyServerStream(m, &grpc.GenericServerStream[emptypb.Empty, emptypb.Empty]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ExcessBodyService_WithBodyServerStreamServer = grpc.ServerStreamingServer[emptypb.Empty]

// ExcessBodyService_ServiceDesc is the grpc.ServiceDesc for ExcessBodyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExcessBodyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.gateway.examples.internal.proto.examplepb.ExcessBodyService",
	HandlerType: (*ExcessBodyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NoBodyRpc",
			Handler:    _ExcessBodyService_NoBodyRpc_Handler,
		},
		{
			MethodName: "WithBodyRpc",
			Handler:    _ExcessBodyService_WithBodyRpc_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "NoBodyServerStream",
			Handler:       _ExcessBodyService_NoBodyServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "WithBodyServerStream",
			Handler:       _ExcessBodyService_WithBodyServerStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "examples/internal/proto/examplepb/excess_body.proto",
}
