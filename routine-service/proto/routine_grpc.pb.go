// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.3
// source: routine.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	RoutineService_CreateRoutine_FullMethodName  = "/routine.RoutineService/CreateRoutine"
	RoutineService_GetRoutine_FullMethodName     = "/routine.RoutineService/GetRoutine"
	RoutineService_UpdateRoutine_FullMethodName  = "/routine.RoutineService/UpdateRoutine"
	RoutineService_DeleteRoutine_FullMethodName  = "/routine.RoutineService/DeleteRoutine"
	RoutineService_GetAllRoutines_FullMethodName = "/routine.RoutineService/GetAllRoutines"
)

// RoutineServiceClient is the client API for RoutineService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RoutineServiceClient interface {
	CreateRoutine(ctx context.Context, in *CreateRoutineRequest, opts ...grpc.CallOption) (*RoutineResponse, error)
	GetRoutine(ctx context.Context, in *GetRoutineRequest, opts ...grpc.CallOption) (*RoutineResponse, error)
	UpdateRoutine(ctx context.Context, in *UpdateRoutineRequest, opts ...grpc.CallOption) (*RoutineResponse, error)
	DeleteRoutine(ctx context.Context, in *DeleteRoutineRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetAllRoutines(ctx context.Context, in *GetAllRoutinesRequest, opts ...grpc.CallOption) (*GetAllRoutinesResponse, error)
}

type routineServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRoutineServiceClient(cc grpc.ClientConnInterface) RoutineServiceClient {
	return &routineServiceClient{cc}
}

func (c *routineServiceClient) CreateRoutine(ctx context.Context, in *CreateRoutineRequest, opts ...grpc.CallOption) (*RoutineResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RoutineResponse)
	err := c.cc.Invoke(ctx, RoutineService_CreateRoutine_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routineServiceClient) GetRoutine(ctx context.Context, in *GetRoutineRequest, opts ...grpc.CallOption) (*RoutineResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RoutineResponse)
	err := c.cc.Invoke(ctx, RoutineService_GetRoutine_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routineServiceClient) UpdateRoutine(ctx context.Context, in *UpdateRoutineRequest, opts ...grpc.CallOption) (*RoutineResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RoutineResponse)
	err := c.cc.Invoke(ctx, RoutineService_UpdateRoutine_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routineServiceClient) DeleteRoutine(ctx context.Context, in *DeleteRoutineRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, RoutineService_DeleteRoutine_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routineServiceClient) GetAllRoutines(ctx context.Context, in *GetAllRoutinesRequest, opts ...grpc.CallOption) (*GetAllRoutinesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllRoutinesResponse)
	err := c.cc.Invoke(ctx, RoutineService_GetAllRoutines_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoutineServiceServer is the server API for RoutineService service.
// All implementations must embed UnimplementedRoutineServiceServer
// for forward compatibility
type RoutineServiceServer interface {
	CreateRoutine(context.Context, *CreateRoutineRequest) (*RoutineResponse, error)
	GetRoutine(context.Context, *GetRoutineRequest) (*RoutineResponse, error)
	UpdateRoutine(context.Context, *UpdateRoutineRequest) (*RoutineResponse, error)
	DeleteRoutine(context.Context, *DeleteRoutineRequest) (*emptypb.Empty, error)
	GetAllRoutines(context.Context, *GetAllRoutinesRequest) (*GetAllRoutinesResponse, error)
	mustEmbedUnimplementedRoutineServiceServer()
}

// UnimplementedRoutineServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRoutineServiceServer struct {
}

func (UnimplementedRoutineServiceServer) CreateRoutine(context.Context, *CreateRoutineRequest) (*RoutineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoutine not implemented")
}
func (UnimplementedRoutineServiceServer) GetRoutine(context.Context, *GetRoutineRequest) (*RoutineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoutine not implemented")
}
func (UnimplementedRoutineServiceServer) UpdateRoutine(context.Context, *UpdateRoutineRequest) (*RoutineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRoutine not implemented")
}
func (UnimplementedRoutineServiceServer) DeleteRoutine(context.Context, *DeleteRoutineRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRoutine not implemented")
}
func (UnimplementedRoutineServiceServer) GetAllRoutines(context.Context, *GetAllRoutinesRequest) (*GetAllRoutinesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllRoutines not implemented")
}
func (UnimplementedRoutineServiceServer) mustEmbedUnimplementedRoutineServiceServer() {}

// UnsafeRoutineServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoutineServiceServer will
// result in compilation errors.
type UnsafeRoutineServiceServer interface {
	mustEmbedUnimplementedRoutineServiceServer()
}

func RegisterRoutineServiceServer(s grpc.ServiceRegistrar, srv RoutineServiceServer) {
	s.RegisterService(&RoutineService_ServiceDesc, srv)
}

func _RoutineService_CreateRoutine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoutineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutineServiceServer).CreateRoutine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoutineService_CreateRoutine_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutineServiceServer).CreateRoutine(ctx, req.(*CreateRoutineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoutineService_GetRoutine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoutineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutineServiceServer).GetRoutine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoutineService_GetRoutine_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutineServiceServer).GetRoutine(ctx, req.(*GetRoutineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoutineService_UpdateRoutine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoutineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutineServiceServer).UpdateRoutine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoutineService_UpdateRoutine_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutineServiceServer).UpdateRoutine(ctx, req.(*UpdateRoutineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoutineService_DeleteRoutine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRoutineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutineServiceServer).DeleteRoutine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoutineService_DeleteRoutine_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutineServiceServer).DeleteRoutine(ctx, req.(*DeleteRoutineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoutineService_GetAllRoutines_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRoutinesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutineServiceServer).GetAllRoutines(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoutineService_GetAllRoutines_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutineServiceServer).GetAllRoutines(ctx, req.(*GetAllRoutinesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RoutineService_ServiceDesc is the grpc.ServiceDesc for RoutineService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RoutineService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "routine.RoutineService",
	HandlerType: (*RoutineServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRoutine",
			Handler:    _RoutineService_CreateRoutine_Handler,
		},
		{
			MethodName: "GetRoutine",
			Handler:    _RoutineService_GetRoutine_Handler,
		},
		{
			MethodName: "UpdateRoutine",
			Handler:    _RoutineService_UpdateRoutine_Handler,
		},
		{
			MethodName: "DeleteRoutine",
			Handler:    _RoutineService_DeleteRoutine_Handler,
		},
		{
			MethodName: "GetAllRoutines",
			Handler:    _RoutineService_GetAllRoutines_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "routine.proto",
}