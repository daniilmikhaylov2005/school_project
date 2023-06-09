// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: api/api.proto

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

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/api.User/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/api.User/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.User/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.User/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _User_CreateUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _User_GetUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/api.proto",
}

// MagazineClient is the client API for Magazine service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MagazineClient interface {
	CreateClass(ctx context.Context, in *CreateClassRequest, opts ...grpc.CallOption) (*CreateClassResponse, error)
	GetClass(ctx context.Context, in *GetClassRequest, opts ...grpc.CallOption) (*GetClassResponse, error)
	GetClassGrades(ctx context.Context, in *GetClassGradesRequest, opts ...grpc.CallOption) (*GetClassGradesResponse, error)
	CreateGrade(ctx context.Context, in *CreateGradeRequest, opts ...grpc.CallOption) (*CreateGradeResponse, error)
	GetGrades(ctx context.Context, in *GetGradesRequest, opts ...grpc.CallOption) (*GetGradesResponse, error)
}

type magazineClient struct {
	cc grpc.ClientConnInterface
}

func NewMagazineClient(cc grpc.ClientConnInterface) MagazineClient {
	return &magazineClient{cc}
}

func (c *magazineClient) CreateClass(ctx context.Context, in *CreateClassRequest, opts ...grpc.CallOption) (*CreateClassResponse, error) {
	out := new(CreateClassResponse)
	err := c.cc.Invoke(ctx, "/api.Magazine/CreateClass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *magazineClient) GetClass(ctx context.Context, in *GetClassRequest, opts ...grpc.CallOption) (*GetClassResponse, error) {
	out := new(GetClassResponse)
	err := c.cc.Invoke(ctx, "/api.Magazine/GetClass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *magazineClient) GetClassGrades(ctx context.Context, in *GetClassGradesRequest, opts ...grpc.CallOption) (*GetClassGradesResponse, error) {
	out := new(GetClassGradesResponse)
	err := c.cc.Invoke(ctx, "/api.Magazine/GetClassGrades", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *magazineClient) CreateGrade(ctx context.Context, in *CreateGradeRequest, opts ...grpc.CallOption) (*CreateGradeResponse, error) {
	out := new(CreateGradeResponse)
	err := c.cc.Invoke(ctx, "/api.Magazine/CreateGrade", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *magazineClient) GetGrades(ctx context.Context, in *GetGradesRequest, opts ...grpc.CallOption) (*GetGradesResponse, error) {
	out := new(GetGradesResponse)
	err := c.cc.Invoke(ctx, "/api.Magazine/GetGrades", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MagazineServer is the server API for Magazine service.
// All implementations must embed UnimplementedMagazineServer
// for forward compatibility
type MagazineServer interface {
	CreateClass(context.Context, *CreateClassRequest) (*CreateClassResponse, error)
	GetClass(context.Context, *GetClassRequest) (*GetClassResponse, error)
	GetClassGrades(context.Context, *GetClassGradesRequest) (*GetClassGradesResponse, error)
	CreateGrade(context.Context, *CreateGradeRequest) (*CreateGradeResponse, error)
	GetGrades(context.Context, *GetGradesRequest) (*GetGradesResponse, error)
	mustEmbedUnimplementedMagazineServer()
}

// UnimplementedMagazineServer must be embedded to have forward compatible implementations.
type UnimplementedMagazineServer struct {
}

func (UnimplementedMagazineServer) CreateClass(context.Context, *CreateClassRequest) (*CreateClassResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateClass not implemented")
}
func (UnimplementedMagazineServer) GetClass(context.Context, *GetClassRequest) (*GetClassResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClass not implemented")
}
func (UnimplementedMagazineServer) GetClassGrades(context.Context, *GetClassGradesRequest) (*GetClassGradesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClassGrades not implemented")
}
func (UnimplementedMagazineServer) CreateGrade(context.Context, *CreateGradeRequest) (*CreateGradeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGrade not implemented")
}
func (UnimplementedMagazineServer) GetGrades(context.Context, *GetGradesRequest) (*GetGradesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGrades not implemented")
}
func (UnimplementedMagazineServer) mustEmbedUnimplementedMagazineServer() {}

// UnsafeMagazineServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MagazineServer will
// result in compilation errors.
type UnsafeMagazineServer interface {
	mustEmbedUnimplementedMagazineServer()
}

func RegisterMagazineServer(s grpc.ServiceRegistrar, srv MagazineServer) {
	s.RegisterService(&Magazine_ServiceDesc, srv)
}

func _Magazine_CreateClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MagazineServer).CreateClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Magazine/CreateClass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MagazineServer).CreateClass(ctx, req.(*CreateClassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Magazine_GetClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MagazineServer).GetClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Magazine/GetClass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MagazineServer).GetClass(ctx, req.(*GetClassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Magazine_GetClassGrades_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClassGradesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MagazineServer).GetClassGrades(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Magazine/GetClassGrades",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MagazineServer).GetClassGrades(ctx, req.(*GetClassGradesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Magazine_CreateGrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGradeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MagazineServer).CreateGrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Magazine/CreateGrade",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MagazineServer).CreateGrade(ctx, req.(*CreateGradeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Magazine_GetGrades_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGradesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MagazineServer).GetGrades(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Magazine/GetGrades",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MagazineServer).GetGrades(ctx, req.(*GetGradesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Magazine_ServiceDesc is the grpc.ServiceDesc for Magazine service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Magazine_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.Magazine",
	HandlerType: (*MagazineServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateClass",
			Handler:    _Magazine_CreateClass_Handler,
		},
		{
			MethodName: "GetClass",
			Handler:    _Magazine_GetClass_Handler,
		},
		{
			MethodName: "GetClassGrades",
			Handler:    _Magazine_GetClassGrades_Handler,
		},
		{
			MethodName: "CreateGrade",
			Handler:    _Magazine_CreateGrade_Handler,
		},
		{
			MethodName: "GetGrades",
			Handler:    _Magazine_GetGrades_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/api.proto",
}
