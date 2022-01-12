package logic

import (
	"context"
	"fmt"
	"strconv"

	"mall/user/rpc/internal/data"
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
	query:=&data.Users{}
	tmpId, _ := strconv.ParseInt(in.Id, 10, 64)
	res,err:=query.FindUserLike(tmpId)
	if err!=nil {
		return nil,err
	}
	likes := []*user.Like{}
	for _,s :=range res.Likes{
		elem:=&user.Like{
			Id: int64(s.ID),
			Title: s.Title,
			Ip: s.Ip,
		}
		fmt.Println("=========================================",s.Ua)
		likes=append(likes,elem)
	}
	u:=&user.UserResponse{
		Id: strconv.Itoa(res.User.UserId),
		Name: res.User.Name,
	}
	return &user.UserLikeResponse{
		User:u,
		Likes: likes,
	}, nil
}
