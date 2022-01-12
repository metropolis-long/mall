package logic

import (
	"context"
	"strconv"

	"mall/order/api/internal/svc"
	"mall/order/api/internal/types"
	"mall/user/rpc/userclient"

	"github.com/tal-tech/go-zero/core/logx"
)

type FindUserLikeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindUserLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) FindUserLikeLogic {
	return FindUserLikeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindUserLikeLogic) FindUserLike(req types.UserLikeReq) (resp *types.UserLikeReply, err error) {
	// todo: add your logic here and delete this line
	data, err := l.svcCtx.UserRpc.FindUserLike(l.ctx, &userclient.UserLikeRequest{
		Id: req.Id,
	})
	likes:=[]types.Like{}
	for _,e:=range data.Likes{
		elem:=types.Like{
			Id: strconv.FormatInt(e.Id,10),
			Title: e.Title,
		}
		likes=append(likes, elem)
	}
	return &types.UserLikeReply{
		User: types.User{
			Id: data.User.Id,
			Name: data.User.Name,
		},
		Likes: likes,
	},nil
}
