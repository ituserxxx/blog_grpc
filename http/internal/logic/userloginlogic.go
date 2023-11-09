package logic

import (
	"brl/http/internal/svc"
	"brl/http/internal/types"
	"brl/rpc/user/userclient"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(ctx context.Context,req *types.LoginReq) (resp *types.LoginResp, err error) {
	res, err := l.svcCtx.UserRpc.UserLogin(ctx, &userclient.UserLoginReq{Username: req.Username,Password: req.Password})
	if err != nil {
		return nil,err
	}
	return &types.LoginResp{Id:res.Id,Username: res.Username},nil
}
