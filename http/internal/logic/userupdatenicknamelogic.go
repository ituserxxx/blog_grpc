package logic

import (
	"brl/rpc/user/userclient"
	"context"

	"brl/http/internal/svc"
	"brl/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserUpdateNickNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserUpdateNickNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUpdateNickNameLogic {
	return &UserUpdateNickNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserUpdateNickNameLogic) UserUpdateNickName(req *types.UserUpdateNickNameReq) (resp *types.UserUpdateNickNameResp, err error) {
	_,err = l.svcCtx.UserRpc.UserUpdateNickName(l.ctx,&userclient.UserUpdateNickNameReq{
		Id: req.Id,
		Nickname: req.Nickname,
	})
	if err != nil {
		return nil,err
	}
	return &types.UserUpdateNickNameResp{},nil
}
