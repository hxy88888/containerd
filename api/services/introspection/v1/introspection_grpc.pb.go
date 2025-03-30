//go:build !no_grpc

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: github.com/containerd/containerd/api/services/introspection/v1/introspection.proto

package introspection

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// IntrospectionClient is the client API for Introspection service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IntrospectionClient interface {
	// Plugins returns a list of plugins in containerd.
	//
	// Clients can use this to detect features and capabilities when using
	// containerd.
	Plugins(ctx context.Context, in *PluginsRequest, opts ...grpc.CallOption) (*PluginsResponse, error)
	// Server returns information about the containerd server
	Server(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ServerResponse, error)
	// PluginInfo returns information directly from a plugin if the plugin supports it
	PluginInfo(ctx context.Context, in *PluginInfoRequest, opts ...grpc.CallOption) (*PluginInfoResponse, error)
}

type introspectionClient struct {
	cc grpc.ClientConnInterface
}

func NewIntrospectionClient(cc grpc.ClientConnInterface) IntrospectionClient {
	return &introspectionClient{cc}
}

func (c *introspectionClient) Plugins(ctx context.Context, in *PluginsRequest, opts ...grpc.CallOption) (*PluginsResponse, error) {
	out := new(PluginsResponse)
	err := c.cc.Invoke(ctx, "/containerd.services.introspection.v1.Introspection/Plugins", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *introspectionClient) Server(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ServerResponse, error) {
	out := new(ServerResponse)
	err := c.cc.Invoke(ctx, "/containerd.services.introspection.v1.Introspection/Server", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *introspectionClient) PluginInfo(ctx context.Context, in *PluginInfoRequest, opts ...grpc.CallOption) (*PluginInfoResponse, error) {
	out := new(PluginInfoResponse)
	err := c.cc.Invoke(ctx, "/containerd.services.introspection.v1.Introspection/PluginInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IntrospectionServer is the server API for Introspection service.
// All implementations must embed UnimplementedIntrospectionServer
// for forward compatibility
type IntrospectionServer interface {
	// Plugins returns a list of plugins in containerd.
	//
	// Clients can use this to detect features and capabilities when using
	// containerd.
	Plugins(context.Context, *PluginsRequest) (*PluginsResponse, error)
	// Server returns information about the containerd server
	Server(context.Context, *emptypb.Empty) (*ServerResponse, error)
	// PluginInfo returns information directly from a plugin if the plugin supports it
	PluginInfo(context.Context, *PluginInfoRequest) (*PluginInfoResponse, error)
	mustEmbedUnimplementedIntrospectionServer()
}

// UnimplementedIntrospectionServer must be embedded to have forward compatible implementations.
type UnimplementedIntrospectionServer struct {
}

func (UnimplementedIntrospectionServer) Plugins(context.Context, *PluginsRequest) (*PluginsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Plugins not implemented")
}
func (UnimplementedIntrospectionServer) Server(context.Context, *emptypb.Empty) (*ServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Server not implemented")
}
func (UnimplementedIntrospectionServer) PluginInfo(context.Context, *PluginInfoRequest) (*PluginInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PluginInfo not implemented")
}
func (UnimplementedIntrospectionServer) mustEmbedUnimplementedIntrospectionServer() {}

// UnsafeIntrospectionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IntrospectionServer will
// result in compilation errors.
type UnsafeIntrospectionServer interface {
	mustEmbedUnimplementedIntrospectionServer()
}

func RegisterIntrospectionServer(s grpc.ServiceRegistrar, srv IntrospectionServer) {
	s.RegisterService(&Introspection_ServiceDesc, srv)
}

func _Introspection_Plugins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntrospectionServer).Plugins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.services.introspection.v1.Introspection/Plugins",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntrospectionServer).Plugins(ctx, req.(*PluginsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Introspection_Server_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntrospectionServer).Server(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.services.introspection.v1.Introspection/Server",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntrospectionServer).Server(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Introspection_PluginInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntrospectionServer).PluginInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.services.introspection.v1.Introspection/PluginInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntrospectionServer).PluginInfo(ctx, req.(*PluginInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Introspection_ServiceDesc is the grpc.ServiceDesc for Introspection service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Introspection_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "containerd.services.introspection.v1.Introspection",
	HandlerType: (*IntrospectionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Plugins",
			Handler:    _Introspection_Plugins_Handler,
		},
		{
			MethodName: "Server",
			Handler:    _Introspection_Server_Handler,
		},
		{
			MethodName: "PluginInfo",
			Handler:    _Introspection_PluginInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/containerd/containerd/api/services/introspection/v1/introspection.proto",
}
