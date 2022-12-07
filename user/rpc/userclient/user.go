// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package userclient

import (
	"context"

	"blog/user/rpc/types/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateUserReq = user.CreateUserReq
	CreateUserRes = user.CreateUserRes
	DelUserReq    = user.DelUserReq
	DelUserRes    = user.DelUserRes
	UpdateUserReq = user.UpdateUserReq
	UpdateUserRes = user.UpdateUserRes
	UserInfoReq   = user.UserInfoReq
	UserInfoRes   = user.UserInfoRes

	User interface {
		CreateUser(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*CreateUserRes, error)
		UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*UpdateUserRes, error)
		DelUser(ctx context.Context, in *DelUserReq, opts ...grpc.CallOption) (*DelUserRes, error)
		UserInfo(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserInfoRes, error)
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

func (m *defaultUser) CreateUser(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*CreateUserRes, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.CreateUser(ctx, in, opts...)
}

func (m *defaultUser) UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*UpdateUserRes, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.UpdateUser(ctx, in, opts...)
}

func (m *defaultUser) DelUser(ctx context.Context, in *DelUserReq, opts ...grpc.CallOption) (*DelUserRes, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.DelUser(ctx, in, opts...)
}

func (m *defaultUser) UserInfo(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserInfoRes, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.UserInfo(ctx, in, opts...)
}
