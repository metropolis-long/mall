// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package userclient

import (
	"context"

	"mall/user/rpc/user"

	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	IdRequest        = user.IdRequest
	Like             = user.Like
	UserLikeRequest  = user.UserLikeRequest
	UserLikeResponse = user.UserLikeResponse
	UserResponse     = user.UserResponse

	User interface {
		GetUser(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*UserResponse, error)
		FindUserLike(ctx context.Context, in *UserLikeRequest, opts ...grpc.CallOption) (*UserLikeResponse, error)
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

func (m *defaultUser) GetUser(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetUser(ctx, in, opts...)
}

func (m *defaultUser) FindUserLike(ctx context.Context, in *UserLikeRequest, opts ...grpc.CallOption) (*UserLikeResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.FindUserLike(ctx, in, opts...)
}
