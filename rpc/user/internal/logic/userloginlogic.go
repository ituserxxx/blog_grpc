package logic

import (
	"context"
	"errors"

	"brl/rpc/user/internal/svc"
	"brl/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserLoginLogic) UserLogin(in *user.UserLoginReq) (*user.UserLoginResp, error) {
	uInfo, err := l.svcCtx.ModelUser.FindOneByUsername(l.ctx, in.Username)
	if err != nil {
		return nil, err
	}
	if uInfo == nil {
		return nil, errors.New("no User")
	}
	if uInfo.Password != in.Password {
		return nil, errors.New("not password")
	}
	return &user.UserLoginResp{
		Id:       uInfo.Id,
		Username: uInfo.Username,
	}, nil
}
