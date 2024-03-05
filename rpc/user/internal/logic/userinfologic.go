package logic

import (
	"context"
	"errors"

	"brl/rpc/user/internal/svc"
	"brl/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *user.UserInfoReq) (*user.UserInfoResp, error) {
	uInfo, err := l.svcCtx.ModelUser.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	if uInfo == nil {
		return nil, errors.New("no User")
	}

	return &user.UserInfoResp{
		Id:       uInfo.Id,
		Username: uInfo.Username,
	}, nil
}
