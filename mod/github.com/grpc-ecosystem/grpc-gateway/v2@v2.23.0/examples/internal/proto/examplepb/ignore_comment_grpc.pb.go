// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: examples/internal/proto/examplepb/ignore_comment.proto

package examplepb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	FooService_Foo_FullMethodName = "/grpc.gateway.examples.internal.proto.examplepb.FooService/Foo"
)

// FooServiceClient is the client API for FooService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// This comment should be excluded from OpenAPI output
type FooServiceClient interface {
	// This comment should be excluded from OpenAPI output
	Foo(ctx context.Context, in *FooRequest, opts ...grpc.CallOption) (*FooReply, error)
}

type fooServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFooServiceClient(cc grpc.ClientConnInterface) FooServiceClient {
	return &fooServiceClient{cc}
}

func (c *fooServiceClient) Foo(ctx context.Context, in *FooRequest, opts ...grpc.CallOption) (*FooReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FooReply)
	err := c.cc.Invoke(ctx, FooService_Foo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FooServiceServer is the server API for FooService service.
// All implementations should embed UnimplementedFooServiceServer
// for forward compatibility.
//
// This comment should be excluded from OpenAPI output
type FooServiceServer interface {
	// This comment should be excluded from OpenAPI output
	Foo(context.Context, *FooRequest) (*FooReply, error)
}

// UnimplementedFooServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedFooServiceServer struct{}

func (UnimplementedFooServiceServer) Foo(context.Context, *FooRequest) (*FooReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Foo not implemented")
}
func (UnimplementedFooServiceServer) testEmbeddedByValue() {}

// UnsafeFooServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FooServiceServer will
// result in compilation errors.
type UnsafeFooServiceServer interface {
	mustEmbedUnimplementedFooServiceServer()
}

func RegisterFooServiceServer(s grpc.ServiceRegistrar, srv FooServiceServer) {
	// If the following call pancis, it indicates UnimplementedFooServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&FooService_ServiceDesc, srv)
}

func _FooService_Foo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FooRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FooServiceServer).Foo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FooService_Foo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FooServiceServer).Foo(ctx, req.(*FooRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FooService_ServiceDesc is the grpc.ServiceDesc for FooService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FooService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.gateway.examples.internal.proto.examplepb.FooService",
	HandlerType: (*FooServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Foo",
			Handler:    _FooService_Foo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "examples/internal/proto/examplepb/ignore_comment.proto",
}