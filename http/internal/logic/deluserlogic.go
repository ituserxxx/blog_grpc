package logic

import (
	"brl/rpc/user/userclient"
	"context"

	"brl/http/internal/svc"
	"brl/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelUserLogic {
	return &DelUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelUserLogic) DelUser(req *types.DelUserReq) (resp *types.DelUserResp, err error) {
	_,err = l.svcCtx.UserRpc.UserDel(l.ctx,&userclient.UserDelReq{Id: req.Id})
	return &types.DelUserResp{},nil
}
