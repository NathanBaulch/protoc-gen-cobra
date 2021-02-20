// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// Proto2Client is the client API for Proto2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Proto2Client interface {
	Echo(ctx context.Context, in *Sound2, opts ...grpc.CallOption) (*Sound2, error)
}

type proto2Client struct {
	cc grpc.ClientConnInterface
}

func NewProto2Client(cc grpc.ClientConnInterface) Proto2Client {
	return &proto2Client{cc}
}

func (c *proto2Client) Echo(ctx context.Context, in *Sound2, opts ...grpc.CallOption) (*Sound2, error) {
	out := new(Sound2)
	err := c.cc.Invoke(ctx, "/example.Proto2/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Proto2Server is the server API for Proto2 service.
// All implementations must embed UnimplementedProto2Server
// for forward compatibility
type Proto2Server interface {
	Echo(context.Context, *Sound2) (*Sound2, error)
	mustEmbedUnimplementedProto2Server()
}

// UnimplementedProto2Server must be embedded to have forward compatible implementations.
type UnimplementedProto2Server struct {
}

func (UnimplementedProto2Server) Echo(context.Context, *Sound2) (*Sound2, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (UnimplementedProto2Server) mustEmbedUnimplementedProto2Server() {}

// UnsafeProto2Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Proto2Server will
// result in compilation errors.
type UnsafeProto2Server interface {
	mustEmbedUnimplementedProto2Server()
}

func RegisterProto2Server(s grpc.ServiceRegistrar, srv Proto2Server) {
	s.RegisterService(&Proto2_ServiceDesc, srv)
}

func _Proto2_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Sound2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Proto2Server).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.Proto2/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Proto2Server).Echo(ctx, req.(*Sound2))
	}
	return interceptor(ctx, in, info, handler)
}

// Proto2_ServiceDesc is the grpc.ServiceDesc for Proto2 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Proto2_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "example.Proto2",
	HandlerType: (*Proto2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _Proto2_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto2.proto",
}
