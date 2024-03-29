// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: api/serviceName/v1/serviceName.proto

package v1

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

// ServiceNameClient is the client API for ServiceName service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceNameClient interface {
	CreateServiceName(ctx context.Context, in *CreateServiceNameRequest, opts ...grpc.CallOption) (*ServiceNameInfoResponse, error)
	UpdateServiceName(ctx context.Context, in *UpdateServiceNameRequest, opts ...grpc.CallOption) (*ServiceNameInfoResponse, error)
	DeleteServiceName(ctx context.Context, in *DeleteServiceNameRequest, opts ...grpc.CallOption) (*ServiceNameCheckResponse, error)
	GetServiceName(ctx context.Context, in *GetServiceNameRequest, opts ...grpc.CallOption) (*ServiceNameInfoResponse, error)
	ListServiceName(ctx context.Context, in *ListServiceNameRequest, opts ...grpc.CallOption) (*ListServiceNameReply, error)
}

type serviceNameClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceNameClient(cc grpc.ClientConnInterface) ServiceNameClient {
	return &serviceNameClient{cc}
}

func (c *serviceNameClient) CreateServiceName(ctx context.Context, in *CreateServiceNameRequest, opts ...grpc.CallOption) (*ServiceNameInfoResponse, error) {
	out := new(ServiceNameInfoResponse)
	err := c.cc.Invoke(ctx, "/api.serviceName.v1.ServiceName/CreateServiceName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceNameClient) UpdateServiceName(ctx context.Context, in *UpdateServiceNameRequest, opts ...grpc.CallOption) (*ServiceNameInfoResponse, error) {
	out := new(ServiceNameInfoResponse)
	err := c.cc.Invoke(ctx, "/api.serviceName.v1.ServiceName/UpdateServiceName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceNameClient) DeleteServiceName(ctx context.Context, in *DeleteServiceNameRequest, opts ...grpc.CallOption) (*ServiceNameCheckResponse, error) {
	out := new(ServiceNameCheckResponse)
	err := c.cc.Invoke(ctx, "/api.serviceName.v1.ServiceName/DeleteServiceName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceNameClient) GetServiceName(ctx context.Context, in *GetServiceNameRequest, opts ...grpc.CallOption) (*ServiceNameInfoResponse, error) {
	out := new(ServiceNameInfoResponse)
	err := c.cc.Invoke(ctx, "/api.serviceName.v1.ServiceName/GetServiceName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceNameClient) ListServiceName(ctx context.Context, in *ListServiceNameRequest, opts ...grpc.CallOption) (*ListServiceNameReply, error) {
	out := new(ListServiceNameReply)
	err := c.cc.Invoke(ctx, "/api.serviceName.v1.ServiceName/ListServiceName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceNameServer is the server API for ServiceName service.
// All implementations must embed UnimplementedServiceNameServer
// for forward compatibility
type ServiceNameServer interface {
	CreateServiceName(context.Context, *CreateServiceNameRequest) (*ServiceNameInfoResponse, error)
	UpdateServiceName(context.Context, *UpdateServiceNameRequest) (*ServiceNameInfoResponse, error)
	DeleteServiceName(context.Context, *DeleteServiceNameRequest) (*ServiceNameCheckResponse, error)
	GetServiceName(context.Context, *GetServiceNameRequest) (*ServiceNameInfoResponse, error)
	ListServiceName(context.Context, *ListServiceNameRequest) (*ListServiceNameReply, error)
	mustEmbedUnimplementedServiceNameServer()
}

// UnimplementedServiceNameServer must be embedded to have forward compatible implementations.
type UnimplementedServiceNameServer struct {
}

func (UnimplementedServiceNameServer) CreateServiceName(context.Context, *CreateServiceNameRequest) (*ServiceNameInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateServiceName not implemented")
}
func (UnimplementedServiceNameServer) UpdateServiceName(context.Context, *UpdateServiceNameRequest) (*ServiceNameInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateServiceName not implemented")
}
func (UnimplementedServiceNameServer) DeleteServiceName(context.Context, *DeleteServiceNameRequest) (*ServiceNameCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteServiceName not implemented")
}
func (UnimplementedServiceNameServer) GetServiceName(context.Context, *GetServiceNameRequest) (*ServiceNameInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServiceName not implemented")
}
func (UnimplementedServiceNameServer) ListServiceName(context.Context, *ListServiceNameRequest) (*ListServiceNameReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListServiceName not implemented")
}
func (UnimplementedServiceNameServer) mustEmbedUnimplementedServiceNameServer() {}

// UnsafeServiceNameServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceNameServer will
// result in compilation errors.
type UnsafeServiceNameServer interface {
	mustEmbedUnimplementedServiceNameServer()
}

func RegisterServiceNameServer(s grpc.ServiceRegistrar, srv ServiceNameServer) {
	s.RegisterService(&ServiceName_ServiceDesc, srv)
}

func _ServiceName_CreateServiceName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateServiceNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceNameServer).CreateServiceName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.serviceName.v1.ServiceName/CreateServiceName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceNameServer).CreateServiceName(ctx, req.(*CreateServiceNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceName_UpdateServiceName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateServiceNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceNameServer).UpdateServiceName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.serviceName.v1.ServiceName/UpdateServiceName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceNameServer).UpdateServiceName(ctx, req.(*UpdateServiceNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceName_DeleteServiceName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteServiceNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceNameServer).DeleteServiceName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.serviceName.v1.ServiceName/DeleteServiceName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceNameServer).DeleteServiceName(ctx, req.(*DeleteServiceNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceName_GetServiceName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServiceNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceNameServer).GetServiceName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.serviceName.v1.ServiceName/GetServiceName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceNameServer).GetServiceName(ctx, req.(*GetServiceNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceName_ListServiceName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListServiceNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceNameServer).ListServiceName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.serviceName.v1.ServiceName/ListServiceName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceNameServer).ListServiceName(ctx, req.(*ListServiceNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ServiceName_ServiceDesc is the grpc.ServiceDesc for ServiceName service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServiceName_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.serviceName.v1.ServiceName",
	HandlerType: (*ServiceNameServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateServiceName",
			Handler:    _ServiceName_CreateServiceName_Handler,
		},
		{
			MethodName: "UpdateServiceName",
			Handler:    _ServiceName_UpdateServiceName_Handler,
		},
		{
			MethodName: "DeleteServiceName",
			Handler:    _ServiceName_DeleteServiceName_Handler,
		},
		{
			MethodName: "GetServiceName",
			Handler:    _ServiceName_GetServiceName_Handler,
		},
		{
			MethodName: "ListServiceName",
			Handler:    _ServiceName_ListServiceName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/serviceName/v1/serviceName.proto",
}
