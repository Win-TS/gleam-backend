// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: modules/group/groupPb/groupPb.proto

package gleam_backend

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

// GroupGrpcServiceClient is the client API for GroupGrpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GroupGrpcServiceClient interface {
	DeleteUserData(ctx context.Context, in *DeleteUserDataReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UserHighestStreak(ctx context.Context, in *UserHighestStreakReq, opts ...grpc.CallOption) (*UserHighestStreakRes, error)
}

type groupGrpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGroupGrpcServiceClient(cc grpc.ClientConnInterface) GroupGrpcServiceClient {
	return &groupGrpcServiceClient{cc}
}

func (c *groupGrpcServiceClient) DeleteUserData(ctx context.Context, in *DeleteUserDataReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/GroupGrpcService/DeleteUserData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupGrpcServiceClient) UserHighestStreak(ctx context.Context, in *UserHighestStreakReq, opts ...grpc.CallOption) (*UserHighestStreakRes, error) {
	out := new(UserHighestStreakRes)
	err := c.cc.Invoke(ctx, "/GroupGrpcService/UserHighestStreak", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GroupGrpcServiceServer is the server API for GroupGrpcService service.
// All implementations must embed UnimplementedGroupGrpcServiceServer
// for forward compatibility
type GroupGrpcServiceServer interface {
	DeleteUserData(context.Context, *DeleteUserDataReq) (*emptypb.Empty, error)
	UserHighestStreak(context.Context, *UserHighestStreakReq) (*UserHighestStreakRes, error)
	mustEmbedUnimplementedGroupGrpcServiceServer()
}

// UnimplementedGroupGrpcServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGroupGrpcServiceServer struct {
}

func (UnimplementedGroupGrpcServiceServer) DeleteUserData(context.Context, *DeleteUserDataReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserData not implemented")
}
func (UnimplementedGroupGrpcServiceServer) UserHighestStreak(context.Context, *UserHighestStreakReq) (*UserHighestStreakRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserHighestStreak not implemented")
}
func (UnimplementedGroupGrpcServiceServer) mustEmbedUnimplementedGroupGrpcServiceServer() {}

// UnsafeGroupGrpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GroupGrpcServiceServer will
// result in compilation errors.
type UnsafeGroupGrpcServiceServer interface {
	mustEmbedUnimplementedGroupGrpcServiceServer()
}

func RegisterGroupGrpcServiceServer(s grpc.ServiceRegistrar, srv GroupGrpcServiceServer) {
	s.RegisterService(&GroupGrpcService_ServiceDesc, srv)
}

func _GroupGrpcService_DeleteUserData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserDataReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupGrpcServiceServer).DeleteUserData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GroupGrpcService/DeleteUserData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupGrpcServiceServer).DeleteUserData(ctx, req.(*DeleteUserDataReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupGrpcService_UserHighestStreak_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserHighestStreakReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupGrpcServiceServer).UserHighestStreak(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GroupGrpcService/UserHighestStreak",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupGrpcServiceServer).UserHighestStreak(ctx, req.(*UserHighestStreakReq))
	}
	return interceptor(ctx, in, info, handler)
}

// GroupGrpcService_ServiceDesc is the grpc.ServiceDesc for GroupGrpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GroupGrpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "GroupGrpcService",
	HandlerType: (*GroupGrpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteUserData",
			Handler:    _GroupGrpcService_DeleteUserData_Handler,
		},
		{
			MethodName: "UserHighestStreak",
			Handler:    _GroupGrpcService_UserHighestStreak_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modules/group/groupPb/groupPb.proto",
}
