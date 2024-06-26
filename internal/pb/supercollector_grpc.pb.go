// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: supercollector.proto

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

const (
	Supercollector_User_FullMethodName = "/supercollector.Supercollector/User"
)

// SupercollectorClient is the client API for Supercollector service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SupercollectorClient interface {
	User(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
}

type supercollectorClient struct {
	cc grpc.ClientConnInterface
}

func NewSupercollectorClient(cc grpc.ClientConnInterface) SupercollectorClient {
	return &supercollectorClient{cc}
}

func (c *supercollectorClient) User(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, Supercollector_User_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SupercollectorServer is the server API for Supercollector service.
// All implementations must embed UnimplementedSupercollectorServer
// for forward compatibility
type SupercollectorServer interface {
	User(context.Context, *UserRequest) (*UserResponse, error)
	mustEmbedUnimplementedSupercollectorServer()
}

// UnimplementedSupercollectorServer must be embedded to have forward compatible implementations.
type UnimplementedSupercollectorServer struct {
}

func (UnimplementedSupercollectorServer) User(context.Context, *UserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method User not implemented")
}
func (UnimplementedSupercollectorServer) mustEmbedUnimplementedSupercollectorServer() {}

// UnsafeSupercollectorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SupercollectorServer will
// result in compilation errors.
type UnsafeSupercollectorServer interface {
	mustEmbedUnimplementedSupercollectorServer()
}

func RegisterSupercollectorServer(s grpc.ServiceRegistrar, srv SupercollectorServer) {
	s.RegisterService(&Supercollector_ServiceDesc, srv)
}

func _Supercollector_User_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupercollectorServer).User(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Supercollector_User_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupercollectorServer).User(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Supercollector_ServiceDesc is the grpc.ServiceDesc for Supercollector service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Supercollector_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "supercollector.Supercollector",
	HandlerType: (*SupercollectorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "User",
			Handler:    _Supercollector_User_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "supercollector.proto",
}
