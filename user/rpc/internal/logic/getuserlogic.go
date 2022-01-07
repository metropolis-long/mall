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

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.IdRequest) (*user.UserResponse, error) {
	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")

	// todo: add your logic here and delete this line
	data:=data.Users{}
	users,err:=data.GetUser(1)
	if err!=nil{
		fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
		return nil,err
	}
	fmt.Println(users.Name)
	return &user.UserResponse{
		Id:     strconv.Itoa(users.UserId),
		Name:   users.Name,
		Gender: "female",
	}, nil
}
