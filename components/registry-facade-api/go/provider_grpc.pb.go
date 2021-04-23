// Copyright (c) 2021 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

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

// SpecProviderClient is the client API for SpecProvider service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SpecProviderClient interface {
	// GetImageSpec provides the image spec for a particular ID. What the ID referes to is specific to
	// the spec provider. For example, in case of ws-manager providing the spec, the ID is a
	// workspace instance ID.
	GetImageSpec(ctx context.Context, in *GetImageSpecRequest, opts ...grpc.CallOption) (*GetImageSpecResponse, error)
	GetOfflineImageSpec(ctx context.Context, in *GetOfflineImageSpecRequest, opts ...grpc.CallOption) (*GetImageSpecResponse, error)
}

type specProviderClient struct {
	cc grpc.ClientConnInterface
}

func NewSpecProviderClient(cc grpc.ClientConnInterface) SpecProviderClient {
	return &specProviderClient{cc}
}

func (c *specProviderClient) GetImageSpec(ctx context.Context, in *GetImageSpecRequest, opts ...grpc.CallOption) (*GetImageSpecResponse, error) {
	out := new(GetImageSpecResponse)
	err := c.cc.Invoke(ctx, "/registryfacade.SpecProvider/GetImageSpec", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *specProviderClient) GetOfflineImageSpec(ctx context.Context, in *GetOfflineImageSpecRequest, opts ...grpc.CallOption) (*GetImageSpecResponse, error) {
	out := new(GetImageSpecResponse)
	err := c.cc.Invoke(ctx, "/registryfacade.SpecProvider/GetOfflineImageSpec", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SpecProviderServer is the server API for SpecProvider service.
// All implementations must embed UnimplementedSpecProviderServer
// for forward compatibility
type SpecProviderServer interface {
	// GetImageSpec provides the image spec for a particular ID. What the ID referes to is specific to
	// the spec provider. For example, in case of ws-manager providing the spec, the ID is a
	// workspace instance ID.
	GetImageSpec(context.Context, *GetImageSpecRequest) (*GetImageSpecResponse, error)
	GetOfflineImageSpec(context.Context, *GetOfflineImageSpecRequest) (*GetImageSpecResponse, error)
	mustEmbedUnimplementedSpecProviderServer()
}

// UnimplementedSpecProviderServer must be embedded to have forward compatible implementations.
type UnimplementedSpecProviderServer struct {
}

func (UnimplementedSpecProviderServer) GetImageSpec(context.Context, *GetImageSpecRequest) (*GetImageSpecResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetImageSpec not implemented")
}
func (UnimplementedSpecProviderServer) GetOfflineImageSpec(context.Context, *GetOfflineImageSpecRequest) (*GetImageSpecResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOfflineImageSpec not implemented")
}
func (UnimplementedSpecProviderServer) mustEmbedUnimplementedSpecProviderServer() {}

// UnsafeSpecProviderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SpecProviderServer will
// result in compilation errors.
type UnsafeSpecProviderServer interface {
	mustEmbedUnimplementedSpecProviderServer()
}

func RegisterSpecProviderServer(s grpc.ServiceRegistrar, srv SpecProviderServer) {
	s.RegisterService(&SpecProvider_ServiceDesc, srv)
}

func _SpecProvider_GetImageSpec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetImageSpecRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpecProviderServer).GetImageSpec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registryfacade.SpecProvider/GetImageSpec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpecProviderServer).GetImageSpec(ctx, req.(*GetImageSpecRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpecProvider_GetOfflineImageSpec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOfflineImageSpecRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpecProviderServer).GetOfflineImageSpec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registryfacade.SpecProvider/GetOfflineImageSpec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpecProviderServer).GetOfflineImageSpec(ctx, req.(*GetOfflineImageSpecRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SpecProvider_ServiceDesc is the grpc.ServiceDesc for SpecProvider service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SpecProvider_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "registryfacade.SpecProvider",
	HandlerType: (*SpecProviderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetImageSpec",
			Handler:    _SpecProvider_GetImageSpec_Handler,
		},
		{
			MethodName: "GetOfflineImageSpec",
			Handler:    _SpecProvider_GetOfflineImageSpec_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "provider.proto",
}

// ContextProviderClient is the client API for ContextProvider service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContextProviderClient interface {
	GetWorkspaceContext(ctx context.Context, in *GetWorkspaceContextRequest, opts ...grpc.CallOption) (*GetWorkspaceContextResponse, error)
}

type contextProviderClient struct {
	cc grpc.ClientConnInterface
}

func NewContextProviderClient(cc grpc.ClientConnInterface) ContextProviderClient {
	return &contextProviderClient{cc}
}

func (c *contextProviderClient) GetWorkspaceContext(ctx context.Context, in *GetWorkspaceContextRequest, opts ...grpc.CallOption) (*GetWorkspaceContextResponse, error) {
	out := new(GetWorkspaceContextResponse)
	err := c.cc.Invoke(ctx, "/registryfacade.ContextProvider/GetWorkspaceContext", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContextProviderServer is the server API for ContextProvider service.
// All implementations must embed UnimplementedContextProviderServer
// for forward compatibility
type ContextProviderServer interface {
	GetWorkspaceContext(context.Context, *GetWorkspaceContextRequest) (*GetWorkspaceContextResponse, error)
	mustEmbedUnimplementedContextProviderServer()
}

// UnimplementedContextProviderServer must be embedded to have forward compatible implementations.
type UnimplementedContextProviderServer struct {
}

func (UnimplementedContextProviderServer) GetWorkspaceContext(context.Context, *GetWorkspaceContextRequest) (*GetWorkspaceContextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkspaceContext not implemented")
}
func (UnimplementedContextProviderServer) mustEmbedUnimplementedContextProviderServer() {}

// UnsafeContextProviderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContextProviderServer will
// result in compilation errors.
type UnsafeContextProviderServer interface {
	mustEmbedUnimplementedContextProviderServer()
}

func RegisterContextProviderServer(s grpc.ServiceRegistrar, srv ContextProviderServer) {
	s.RegisterService(&ContextProvider_ServiceDesc, srv)
}

func _ContextProvider_GetWorkspaceContext_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkspaceContextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContextProviderServer).GetWorkspaceContext(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registryfacade.ContextProvider/GetWorkspaceContext",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContextProviderServer).GetWorkspaceContext(ctx, req.(*GetWorkspaceContextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ContextProvider_ServiceDesc is the grpc.ServiceDesc for ContextProvider service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ContextProvider_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "registryfacade.ContextProvider",
	HandlerType: (*ContextProviderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWorkspaceContext",
			Handler:    _ContextProvider_GetWorkspaceContext_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "provider.proto",
}
