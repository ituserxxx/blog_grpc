// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package userclient

import (
	"context"

	"brl/rpc/user/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddUserReq            = user.AddUserReq
	AddUserResp           = user.AddUserResp
	EmptyResp             = user.EmptyResp
	Request               = user.Request
	Response              = user.Response
	UserDelReq            = user.UserDelReq
	UserInfoReq           = user.UserInfoReq
	UserInfoResp          = user.UserInfoResp
	UserListItem          = user.UserListItem
	UserListReq           = user.UserListReq
	UserListResp          = user.UserListResp
	UserLoginReq          = user.UserLoginReq
	UserLoginResp         = user.UserLoginResp
	UserUpdateNickNameReq = user.UserUpdateNickNameReq

	User interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
		AddUser(ctx context.Context, in *AddUserReq, opts ...grpc.CallOption) (*AddUserResp, error)
		UserLogin(ctx context.Context, in *UserLoginReq, opts ...grpc.CallOption) (*UserLoginResp, error)
		UserUpdateNickName(ctx context.Context, in *UserUpdateNickNameReq, opts ...grpc.CallOption) (*EmptyResp, error)
		UserDel(ctx context.Context, in *UserDelReq, opts ...grpc.CallOption) (*EmptyResp, error)
		UserInfo(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserInfoResp, error)
		UserList(ctx context.Context, in *UserListReq, opts ...grpc.CallOption) (*UserListResp, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}

func (m *defaultUser) AddUser(ctx context.Context, in *AddUserReq, opts ...grpc.CallOption) (*AddUserResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.AddUser(ctx, in, opts...)
}

func (m *defaultUser) UserLogin(ctx context.Context, in *UserLoginReq, opts ...grpc.CallOption) (*UserLoginResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.UserLogin(ctx, in, opts...)
}

func (m *defaultUser) UserUpdateNickName(ctx context.Context, in *UserUpdateNickNameReq, opts ...grpc.CallOption) (*EmptyResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.UserUpdateNickName(ctx, in, opts...)
}

func (m *defaultUser) UserDel(ctx context.Context, in *UserDelReq, opts ...grpc.CallOption) (*EmptyResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.UserDel(ctx, in, opts...)
}

func (m *defaultUser) UserInfo(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserInfoResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.UserInfo(ctx, in, opts...)
}

func (m *defaultUser) UserList(ctx context.Context, in *UserListReq, opts ...grpc.CallOption) (*UserListResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.UserList(ctx, in, opts...)
}
