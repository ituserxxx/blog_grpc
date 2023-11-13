package user

import (
	"brl/rpc/user/userclient"
	"context"

	"brl/http/internal/svc"
	"brl/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.GetUserInfoReq) (resp *types.GetUserInfoResp, err error) {
	res, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &userclient.UserInfoReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return &types.GetUserInfoResp{
		Id:       res.Id,
		Username: res.Username,
	}, nil
}
