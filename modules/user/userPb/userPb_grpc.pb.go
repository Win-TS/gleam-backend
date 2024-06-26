// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: modules/user/userPb/userPb.proto

package gleam_backend

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

// UserGrpcServiceClient is the client API for UserGrpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserGrpcServiceClient interface {
	SearchUser(ctx context.Context, in *SearchUserReq, opts ...grpc.CallOption) (*SearchUserRes, error)
	GetUserProfile(ctx context.Context, in *GetUserProfileReq, opts ...grpc.CallOption) (*GetUserProfileRes, error)
	GetBatchUserProfiles(ctx context.Context, in *GetBatchUserProfileReq, opts ...grpc.CallOption) (*GetBatchUserProfileRes, error)
	GetUserFriends(ctx context.Context, in *GetUserFriendsReq, opts ...grpc.CallOption) (*GetUserFriendsRes, error)
}

type userGrpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserGrpcServiceClient(cc grpc.ClientConnInterface) UserGrpcServiceClient {
	return &userGrpcServiceClient{cc}
}

func (c *userGrpcServiceClient) SearchUser(ctx context.Context, in *SearchUserReq, opts ...grpc.CallOption) (*SearchUserRes, error) {
	out := new(SearchUserRes)
	err := c.cc.Invoke(ctx, "/UserGrpcService/SearchUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userGrpcServiceClient) GetUserProfile(ctx context.Context, in *GetUserProfileReq, opts ...grpc.CallOption) (*GetUserProfileRes, error) {
	out := new(GetUserProfileRes)
	err := c.cc.Invoke(ctx, "/UserGrpcService/GetUserProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userGrpcServiceClient) GetBatchUserProfiles(ctx context.Context, in *GetBatchUserProfileReq, opts ...grpc.CallOption) (*GetBatchUserProfileRes, error) {
	out := new(GetBatchUserProfileRes)
	err := c.cc.Invoke(ctx, "/UserGrpcService/GetBatchUserProfiles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userGrpcServiceClient) GetUserFriends(ctx context.Context, in *GetUserFriendsReq, opts ...grpc.CallOption) (*GetUserFriendsRes, error) {
	out := new(GetUserFriendsRes)
	err := c.cc.Invoke(ctx, "/UserGrpcService/GetUserFriends", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserGrpcServiceServer is the server API for UserGrpcService service.
// All implementations must embed UnimplementedUserGrpcServiceServer
// for forward compatibility
type UserGrpcServiceServer interface {
	SearchUser(context.Context, *SearchUserReq) (*SearchUserRes, error)
	GetUserProfile(context.Context, *GetUserProfileReq) (*GetUserProfileRes, error)
	GetBatchUserProfiles(context.Context, *GetBatchUserProfileReq) (*GetBatchUserProfileRes, error)
	GetUserFriends(context.Context, *GetUserFriendsReq) (*GetUserFriendsRes, error)
	mustEmbedUnimplementedUserGrpcServiceServer()
}

// UnimplementedUserGrpcServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserGrpcServiceServer struct {
}

func (UnimplementedUserGrpcServiceServer) SearchUser(context.Context, *SearchUserReq) (*SearchUserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchUser not implemented")
}
func (UnimplementedUserGrpcServiceServer) GetUserProfile(context.Context, *GetUserProfileReq) (*GetUserProfileRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserProfile not implemented")
}
func (UnimplementedUserGrpcServiceServer) GetBatchUserProfiles(context.Context, *GetBatchUserProfileReq) (*GetBatchUserProfileRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBatchUserProfiles not implemented")
}
func (UnimplementedUserGrpcServiceServer) GetUserFriends(context.Context, *GetUserFriendsReq) (*GetUserFriendsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserFriends not implemented")
}
func (UnimplementedUserGrpcServiceServer) mustEmbedUnimplementedUserGrpcServiceServer() {}

// UnsafeUserGrpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserGrpcServiceServer will
// result in compilation errors.
type UnsafeUserGrpcServiceServer interface {
	mustEmbedUnimplementedUserGrpcServiceServer()
}

func RegisterUserGrpcServiceServer(s grpc.ServiceRegistrar, srv UserGrpcServiceServer) {
	s.RegisterService(&UserGrpcService_ServiceDesc, srv)
}

func _UserGrpcService_SearchUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserGrpcServiceServer).SearchUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserGrpcService/SearchUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserGrpcServiceServer).SearchUser(ctx, req.(*SearchUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserGrpcService_GetUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserProfileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserGrpcServiceServer).GetUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserGrpcService/GetUserProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserGrpcServiceServer).GetUserProfile(ctx, req.(*GetUserProfileReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserGrpcService_GetBatchUserProfiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBatchUserProfileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserGrpcServiceServer).GetBatchUserProfiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserGrpcService/GetBatchUserProfiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserGrpcServiceServer).GetBatchUserProfiles(ctx, req.(*GetBatchUserProfileReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserGrpcService_GetUserFriends_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserFriendsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserGrpcServiceServer).GetUserFriends(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserGrpcService/GetUserFriends",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserGrpcServiceServer).GetUserFriends(ctx, req.(*GetUserFriendsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserGrpcService_ServiceDesc is the grpc.ServiceDesc for UserGrpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserGrpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "UserGrpcService",
	HandlerType: (*UserGrpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchUser",
			Handler:    _UserGrpcService_SearchUser_Handler,
		},
		{
			MethodName: "GetUserProfile",
			Handler:    _UserGrpcService_GetUserProfile_Handler,
		},
		{
			MethodName: "GetBatchUserProfiles",
			Handler:    _UserGrpcService_GetBatchUserProfiles_Handler,
		},
		{
			MethodName: "GetUserFriends",
			Handler:    _UserGrpcService_GetUserFriends_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modules/user/userPb/userPb.proto",
}
