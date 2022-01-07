package logic

import (
	"context"

	"mall/user/rpc/internal/svc"
	"mall/user/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type FindUserLikeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserLikeLogic {
	return &FindUserLikeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindUserLikeLogic) FindUserLike(in *user.UserLikeRequest) (*user.UserLikeResponse, error) {
	// todo: add your logic here and delete this line

	return &user.UserLikeResponse{}, nil
}
