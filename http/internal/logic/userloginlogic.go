package logic

import (
	"context"
	"brl/rpc/user/userclient"
	"brl/http/internal/types"
	"brl/http/internal/svc"
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
	// todo: add your logic here and delete this line
	l.svcCtx.UserRpc.Ping(ctx,&userclient.Request{})
	return
}
