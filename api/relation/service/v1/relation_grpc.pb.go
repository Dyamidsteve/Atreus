// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: relation/service/v1/relation.proto

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

const (
	RelationService_GetFollowerRelationList_FullMethodName = "/relation.service.v1.RelationService/GetFollowerRelationList"
	RelationService_GetFollowRelationList_FullMethodName   = "/relation.service.v1.RelationService/GetFollowRelationList"
	RelationService_RelationAction_FullMethodName          = "/relation.service.v1.RelationService/RelationAction"
	RelationService_IsFollow_FullMethodName                = "/relation.service.v1.RelationService/IsFollow"
)

// RelationServiceClient is the client API for RelationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RelationServiceClient interface {
	// 获取粉丝列表(客户端)
	GetFollowerRelationList(ctx context.Context, in *RelationFollowerListRequest, opts ...grpc.CallOption) (*RelationFollowerListReply, error)
	// 获取关注列表(客户端)
	GetFollowRelationList(ctx context.Context, in *RelationFollowListRequest, opts ...grpc.CallOption) (*RelationFollowListReply, error)
	// 关注或取关用户(客户端)
	RelationAction(ctx context.Context, in *RelationActionRequest, opts ...grpc.CallOption) (*RelationActionReply, error)
	// 根据userId和toUserId判断是否关注(user)
	IsFollow(ctx context.Context, in *IsFollowRequest, opts ...grpc.CallOption) (*IsFollowReply, error)
}

type relationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRelationServiceClient(cc grpc.ClientConnInterface) RelationServiceClient {
	return &relationServiceClient{cc}
}

func (c *relationServiceClient) GetFollowerRelationList(ctx context.Context, in *RelationFollowerListRequest, opts ...grpc.CallOption) (*RelationFollowerListReply, error) {
	out := new(RelationFollowerListReply)
	err := c.cc.Invoke(ctx, RelationService_GetFollowerRelationList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) GetFollowRelationList(ctx context.Context, in *RelationFollowListRequest, opts ...grpc.CallOption) (*RelationFollowListReply, error) {
	out := new(RelationFollowListReply)
	err := c.cc.Invoke(ctx, RelationService_GetFollowRelationList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) RelationAction(ctx context.Context, in *RelationActionRequest, opts ...grpc.CallOption) (*RelationActionReply, error) {
	out := new(RelationActionReply)
	err := c.cc.Invoke(ctx, RelationService_RelationAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) IsFollow(ctx context.Context, in *IsFollowRequest, opts ...grpc.CallOption) (*IsFollowReply, error) {
	out := new(IsFollowReply)
	err := c.cc.Invoke(ctx, RelationService_IsFollow_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RelationServiceServer is the server API for RelationService service.
// All implementations must embed UnimplementedRelationServiceServer
// for forward compatibility
type RelationServiceServer interface {
	// 获取粉丝列表(客户端)
	GetFollowerRelationList(context.Context, *RelationFollowerListRequest) (*RelationFollowerListReply, error)
	// 获取关注列表(客户端)
	GetFollowRelationList(context.Context, *RelationFollowListRequest) (*RelationFollowListReply, error)
	// 关注或取关用户(客户端)
	RelationAction(context.Context, *RelationActionRequest) (*RelationActionReply, error)
	// 根据userId和toUserId判断是否关注(user)
	IsFollow(context.Context, *IsFollowRequest) (*IsFollowReply, error)
	mustEmbedUnimplementedRelationServiceServer()
}

// UnimplementedRelationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRelationServiceServer struct {
}

func (UnimplementedRelationServiceServer) GetFollowerRelationList(context.Context, *RelationFollowerListRequest) (*RelationFollowerListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFollowerRelationList not implemented")
}
func (UnimplementedRelationServiceServer) GetFollowRelationList(context.Context, *RelationFollowListRequest) (*RelationFollowListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFollowRelationList not implemented")
}
func (UnimplementedRelationServiceServer) RelationAction(context.Context, *RelationActionRequest) (*RelationActionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RelationAction not implemented")
}
func (UnimplementedRelationServiceServer) IsFollow(context.Context, *IsFollowRequest) (*IsFollowReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsFollow not implemented")
}
func (UnimplementedRelationServiceServer) mustEmbedUnimplementedRelationServiceServer() {}

// UnsafeRelationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RelationServiceServer will
// result in compilation errors.
type UnsafeRelationServiceServer interface {
	mustEmbedUnimplementedRelationServiceServer()
}

func RegisterRelationServiceServer(s grpc.ServiceRegistrar, srv RelationServiceServer) {
	s.RegisterService(&RelationService_ServiceDesc, srv)
}

func _RelationService_GetFollowerRelationList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RelationFollowerListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).GetFollowerRelationList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RelationService_GetFollowerRelationList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).GetFollowerRelationList(ctx, req.(*RelationFollowerListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_GetFollowRelationList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RelationFollowListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).GetFollowRelationList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RelationService_GetFollowRelationList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).GetFollowRelationList(ctx, req.(*RelationFollowListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_RelationAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RelationActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).RelationAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RelationService_RelationAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).RelationAction(ctx, req.(*RelationActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_IsFollow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsFollowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).IsFollow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RelationService_IsFollow_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).IsFollow(ctx, req.(*IsFollowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RelationService_ServiceDesc is the grpc.ServiceDesc for RelationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RelationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "relation.service.v1.RelationService",
	HandlerType: (*RelationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFollowerRelationList",
			Handler:    _RelationService_GetFollowerRelationList_Handler,
		},
		{
			MethodName: "GetFollowRelationList",
			Handler:    _RelationService_GetFollowRelationList_Handler,
		},
		{
			MethodName: "RelationAction",
			Handler:    _RelationService_RelationAction_Handler,
		},
		{
			MethodName: "IsFollow",
			Handler:    _RelationService_IsFollow_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "relation/service/v1/relation.proto",
}
