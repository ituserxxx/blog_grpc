package logic

import (
	"brl/rpc/user/userclient"
	"context"

	"brl/http/internal/svc"
	"brl/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddUserLogic) AddUser(req *types.AddUserReq) (resp *types.AddUserResp, err error) {
	id,err := l.svcCtx.UserRpc.AddUser(l.ctx,&userclient.AddUserReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil,err
	}
	return &types.AddUserResp{Id: id.Id},nil
}
