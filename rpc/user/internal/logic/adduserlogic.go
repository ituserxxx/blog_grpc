package logic

import (
	userModel "brl/rpc/mysql/model/user"
	"brl/rpc/user/internal/svc"
	"brl/rpc/user/user"
	"context"

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
	inres, err := l.svcCtx.ModelUser.Insert(l.ctx, &userModel.BlogUser{
		Username: in.Username,
		Password: in.Password,
	})
	if err != nil {
		return nil, err
	}
	id, err := inres.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &user.AddUserResp{
		Id: id,
	}, nil
}
