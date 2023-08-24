// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.3
// - protoc             v4.23.4
// source: relation/service/v1/relation.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationRelationServiceGetFollowRelationList = "/relation.service.v1.RelationService/GetFollowRelationList"
const OperationRelationServiceGetFollowerRelationList = "/relation.service.v1.RelationService/GetFollowerRelationList"
const OperationRelationServiceGetFriendRelationList = "/relation.service.v1.RelationService/GetFriendRelationList"
const OperationRelationServiceRelationAction = "/relation.service.v1.RelationService/RelationAction"

type RelationServiceHTTPServer interface {
	// GetFollowRelationList 获取关注列表(客户端)
	GetFollowRelationList(context.Context, *RelationFollowListRequest) (*RelationFollowListReply, error)
	// GetFollowerRelationList 获取粉丝列表(客户端)
	GetFollowerRelationList(context.Context, *RelationFollowerListRequest) (*RelationFollowerListReply, error)
	// GetFriendRelationList 获取好友列表(客户端)
	GetFriendRelationList(context.Context, *RelationFriendListRequest) (*RelationFriendListReply, error)
	// RelationAction 关注或取关用户(客户端)
	RelationAction(context.Context, *RelationActionRequest) (*RelationActionReply, error)
}

func RegisterRelationServiceHTTPServer(s *http.Server, srv RelationServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/douyin/relation/follower/list", _RelationService_GetFollowerRelationList0_HTTP_Handler(srv))
	r.GET("/douyin/relation/follow/list", _RelationService_GetFollowRelationList0_HTTP_Handler(srv))
	r.POST("/douyin/relation/action", _RelationService_RelationAction0_HTTP_Handler(srv))
	r.GET("/douyin/relation/friend/list", _RelationService_GetFriendRelationList0_HTTP_Handler(srv))
}

func _RelationService_GetFollowerRelationList0_HTTP_Handler(srv RelationServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RelationFollowerListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRelationServiceGetFollowerRelationList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetFollowerRelationList(ctx, req.(*RelationFollowerListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RelationFollowerListReply)
		return ctx.Result(200, reply)
	}
}

func _RelationService_GetFollowRelationList0_HTTP_Handler(srv RelationServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RelationFollowListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRelationServiceGetFollowRelationList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetFollowRelationList(ctx, req.(*RelationFollowListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RelationFollowListReply)
		return ctx.Result(200, reply)
	}
}

func _RelationService_RelationAction0_HTTP_Handler(srv RelationServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RelationActionRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRelationServiceRelationAction)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RelationAction(ctx, req.(*RelationActionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RelationActionReply)
		return ctx.Result(200, reply)
	}
}

func _RelationService_GetFriendRelationList0_HTTP_Handler(srv RelationServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RelationFriendListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRelationServiceGetFriendRelationList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetFriendRelationList(ctx, req.(*RelationFriendListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RelationFriendListReply)
		return ctx.Result(200, reply)
	}
}

type RelationServiceHTTPClient interface {
	GetFollowRelationList(ctx context.Context, req *RelationFollowListRequest, opts ...http.CallOption) (rsp *RelationFollowListReply, err error)
	GetFollowerRelationList(ctx context.Context, req *RelationFollowerListRequest, opts ...http.CallOption) (rsp *RelationFollowerListReply, err error)
	GetFriendRelationList(ctx context.Context, req *RelationFriendListRequest, opts ...http.CallOption) (rsp *RelationFriendListReply, err error)
	RelationAction(ctx context.Context, req *RelationActionRequest, opts ...http.CallOption) (rsp *RelationActionReply, err error)
}

type RelationServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewRelationServiceHTTPClient(client *http.Client) RelationServiceHTTPClient {
	return &RelationServiceHTTPClientImpl{client}
}

func (c *RelationServiceHTTPClientImpl) GetFollowRelationList(ctx context.Context, in *RelationFollowListRequest, opts ...http.CallOption) (*RelationFollowListReply, error) {
	var out RelationFollowListReply
	pattern := "/douyin/relation/follow/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationRelationServiceGetFollowRelationList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RelationServiceHTTPClientImpl) GetFollowerRelationList(ctx context.Context, in *RelationFollowerListRequest, opts ...http.CallOption) (*RelationFollowerListReply, error) {
	var out RelationFollowerListReply
	pattern := "/douyin/relation/follower/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationRelationServiceGetFollowerRelationList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RelationServiceHTTPClientImpl) GetFriendRelationList(ctx context.Context, in *RelationFriendListRequest, opts ...http.CallOption) (*RelationFriendListReply, error) {
	var out RelationFriendListReply
	pattern := "/douyin/relation/friend/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationRelationServiceGetFriendRelationList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RelationServiceHTTPClientImpl) RelationAction(ctx context.Context, in *RelationActionRequest, opts ...http.CallOption) (*RelationActionReply, error) {
	var out RelationActionReply
	pattern := "/douyin/relation/action"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationRelationServiceRelationAction))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
