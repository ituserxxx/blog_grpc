package logic

import (
	"context"

	"brl/rpc/user/internal/svc"
	"brl/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddUserLogic) AddUser(in *user.AddUserReq) (*user.AddUserResp, error) {
	// todo: add your logic here and delete this line
	
	return &user.AddUserResp{}, nil
}
