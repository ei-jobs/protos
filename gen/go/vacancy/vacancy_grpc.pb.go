// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: vacancy/vacancy.proto

package vacancyv1

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
	VacancyService_GetVacancyById_FullMethodName     = "/vacancy.VacancyService/GetVacancyById"
	VacancyService_ListVacancy_FullMethodName        = "/vacancy.VacancyService/ListVacancy"
	VacancyService_CreateVacancy_FullMethodName      = "/vacancy.VacancyService/CreateVacancy"
	VacancyService_GetVacancyByUserId_FullMethodName = "/vacancy.VacancyService/GetVacancyByUserId"
	VacancyService_UpdateVacancy_FullMethodName      = "/vacancy.VacancyService/UpdateVacancy"
	VacancyService_StopVacancy_FullMethodName        = "/vacancy.VacancyService/StopVacancy"
	VacancyService_DeleteVacancy_FullMethodName      = "/vacancy.VacancyService/DeleteVacancy"
)

// VacancyServiceClient is the client API for VacancyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VacancyServiceClient interface {
	GetVacancyById(ctx context.Context, in *GetVacancyByIdRequest, opts ...grpc.CallOption) (*GetVacancyByIdResponse, error)
	ListVacancy(ctx context.Context, in *ListVacancyRequest, opts ...grpc.CallOption) (*ListVacancyResponse, error)
	CreateVacancy(ctx context.Context, in *CreateVacancyRequest, opts ...grpc.CallOption) (*CreateVacancyResponse, error)
	GetVacancyByUserId(ctx context.Context, in *GetVacancyByUserIdRequest, opts ...grpc.CallOption) (*GetVacancyByUserIdResponse, error)
	UpdateVacancy(ctx context.Context, in *UpdateVacancyRequest, opts ...grpc.CallOption) (*UpdateVacancyResponse, error)
	StopVacancy(ctx context.Context, in *StopVacancyRequest, opts ...grpc.CallOption) (*StopVacancyResponse, error)
	DeleteVacancy(ctx context.Context, in *DeleteVacancyRequest, opts ...grpc.CallOption) (*DeleteVacancyResponse, error)
}

type vacancyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVacancyServiceClient(cc grpc.ClientConnInterface) VacancyServiceClient {
	return &vacancyServiceClient{cc}
}

func (c *vacancyServiceClient) GetVacancyById(ctx context.Context, in *GetVacancyByIdRequest, opts ...grpc.CallOption) (*GetVacancyByIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetVacancyByIdResponse)
	err := c.cc.Invoke(ctx, VacancyService_GetVacancyById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vacancyServiceClient) ListVacancy(ctx context.Context, in *ListVacancyRequest, opts ...grpc.CallOption) (*ListVacancyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListVacancyResponse)
	err := c.cc.Invoke(ctx, VacancyService_ListVacancy_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vacancyServiceClient) CreateVacancy(ctx context.Context, in *CreateVacancyRequest, opts ...grpc.CallOption) (*CreateVacancyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateVacancyResponse)
	err := c.cc.Invoke(ctx, VacancyService_CreateVacancy_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vacancyServiceClient) GetVacancyByUserId(ctx context.Context, in *GetVacancyByUserIdRequest, opts ...grpc.CallOption) (*GetVacancyByUserIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetVacancyByUserIdResponse)
	err := c.cc.Invoke(ctx, VacancyService_GetVacancyByUserId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vacancyServiceClient) UpdateVacancy(ctx context.Context, in *UpdateVacancyRequest, opts ...grpc.CallOption) (*UpdateVacancyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateVacancyResponse)
	err := c.cc.Invoke(ctx, VacancyService_UpdateVacancy_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vacancyServiceClient) StopVacancy(ctx context.Context, in *StopVacancyRequest, opts ...grpc.CallOption) (*StopVacancyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StopVacancyResponse)
	err := c.cc.Invoke(ctx, VacancyService_StopVacancy_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vacancyServiceClient) DeleteVacancy(ctx context.Context, in *DeleteVacancyRequest, opts ...grpc.CallOption) (*DeleteVacancyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteVacancyResponse)
	err := c.cc.Invoke(ctx, VacancyService_DeleteVacancy_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VacancyServiceServer is the server API for VacancyService service.
// All implementations must embed UnimplementedVacancyServiceServer
// for forward compatibility.
type VacancyServiceServer interface {
	GetVacancyById(context.Context, *GetVacancyByIdRequest) (*GetVacancyByIdResponse, error)
	ListVacancy(context.Context, *ListVacancyRequest) (*ListVacancyResponse, error)
	CreateVacancy(context.Context, *CreateVacancyRequest) (*CreateVacancyResponse, error)
	GetVacancyByUserId(context.Context, *GetVacancyByUserIdRequest) (*GetVacancyByUserIdResponse, error)
	UpdateVacancy(context.Context, *UpdateVacancyRequest) (*UpdateVacancyResponse, error)
	StopVacancy(context.Context, *StopVacancyRequest) (*StopVacancyResponse, error)
	DeleteVacancy(context.Context, *DeleteVacancyRequest) (*DeleteVacancyResponse, error)
	mustEmbedUnimplementedVacancyServiceServer()
}

// UnimplementedVacancyServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedVacancyServiceServer struct{}

func (UnimplementedVacancyServiceServer) GetVacancyById(context.Context, *GetVacancyByIdRequest) (*GetVacancyByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVacancyById not implemented")
}
func (UnimplementedVacancyServiceServer) ListVacancy(context.Context, *ListVacancyRequest) (*ListVacancyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVacancy not implemented")
}
func (UnimplementedVacancyServiceServer) CreateVacancy(context.Context, *CreateVacancyRequest) (*CreateVacancyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVacancy not implemented")
}
func (UnimplementedVacancyServiceServer) GetVacancyByUserId(context.Context, *GetVacancyByUserIdRequest) (*GetVacancyByUserIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVacancyByUserId not implemented")
}
func (UnimplementedVacancyServiceServer) UpdateVacancy(context.Context, *UpdateVacancyRequest) (*UpdateVacancyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateVacancy not implemented")
}
func (UnimplementedVacancyServiceServer) StopVacancy(context.Context, *StopVacancyRequest) (*StopVacancyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopVacancy not implemented")
}
func (UnimplementedVacancyServiceServer) DeleteVacancy(context.Context, *DeleteVacancyRequest) (*DeleteVacancyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteVacancy not implemented")
}
func (UnimplementedVacancyServiceServer) mustEmbedUnimplementedVacancyServiceServer() {}
func (UnimplementedVacancyServiceServer) testEmbeddedByValue()                        {}

// UnsafeVacancyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VacancyServiceServer will
// result in compilation errors.
type UnsafeVacancyServiceServer interface {
	mustEmbedUnimplementedVacancyServiceServer()
}

func RegisterVacancyServiceServer(s grpc.ServiceRegistrar, srv VacancyServiceServer) {
	// If the following call pancis, it indicates UnimplementedVacancyServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&VacancyService_ServiceDesc, srv)
}

func _VacancyService_GetVacancyById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVacancyByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VacancyServiceServer).GetVacancyById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VacancyService_GetVacancyById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VacancyServiceServer).GetVacancyById(ctx, req.(*GetVacancyByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VacancyService_ListVacancy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListVacancyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VacancyServiceServer).ListVacancy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VacancyService_ListVacancy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VacancyServiceServer).ListVacancy(ctx, req.(*ListVacancyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VacancyService_CreateVacancy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVacancyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VacancyServiceServer).CreateVacancy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VacancyService_CreateVacancy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VacancyServiceServer).CreateVacancy(ctx, req.(*CreateVacancyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VacancyService_GetVacancyByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVacancyByUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VacancyServiceServer).GetVacancyByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VacancyService_GetVacancyByUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VacancyServiceServer).GetVacancyByUserId(ctx, req.(*GetVacancyByUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VacancyService_UpdateVacancy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateVacancyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VacancyServiceServer).UpdateVacancy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VacancyService_UpdateVacancy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VacancyServiceServer).UpdateVacancy(ctx, req.(*UpdateVacancyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VacancyService_StopVacancy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopVacancyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VacancyServiceServer).StopVacancy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VacancyService_StopVacancy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VacancyServiceServer).StopVacancy(ctx, req.(*StopVacancyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VacancyService_DeleteVacancy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteVacancyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VacancyServiceServer).DeleteVacancy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VacancyService_DeleteVacancy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VacancyServiceServer).DeleteVacancy(ctx, req.(*DeleteVacancyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VacancyService_ServiceDesc is the grpc.ServiceDesc for VacancyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VacancyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vacancy.VacancyService",
	HandlerType: (*VacancyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVacancyById",
			Handler:    _VacancyService_GetVacancyById_Handler,
		},
		{
			MethodName: "ListVacancy",
			Handler:    _VacancyService_ListVacancy_Handler,
		},
		{
			MethodName: "CreateVacancy",
			Handler:    _VacancyService_CreateVacancy_Handler,
		},
		{
			MethodName: "GetVacancyByUserId",
			Handler:    _VacancyService_GetVacancyByUserId_Handler,
		},
		{
			MethodName: "UpdateVacancy",
			Handler:    _VacancyService_UpdateVacancy_Handler,
		},
		{
			MethodName: "StopVacancy",
			Handler:    _VacancyService_StopVacancy_Handler,
		},
		{
			MethodName: "DeleteVacancy",
			Handler:    _VacancyService_DeleteVacancy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vacancy/vacancy.proto",
}
