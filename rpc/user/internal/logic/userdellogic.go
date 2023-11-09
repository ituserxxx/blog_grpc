package logic

import (
	"context"

	"brl/rpc/user/internal/svc"
	"brl/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDelLogic {
	return &UserDelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserDelLogic) UserDel(in *user.UserDelReq) (*user.EmptyResp, error) {
	err := l.svcCtx.BlogUserModel.Delete(l.ctx,in.Id)
	if err != nil {
		return nil, err
	}
	return &user.EmptyResp{}, nil
}
