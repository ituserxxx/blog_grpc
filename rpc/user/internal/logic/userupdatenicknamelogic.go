package logic

import (
	userModel "brl/rpc/mysql/model/user"
	"context"
	"database/sql"

	"brl/rpc/user/internal/svc"
	"brl/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserUpdateNickNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserUpdateNickNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUpdateNickNameLogic {
	return &UserUpdateNickNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserUpdateNickNameLogic) UserUpdateNickName(in *user.UserUpdateNickNameReq) (*user.EmptyResp, error) {
	err := l.svcCtx.ModelUser.Update(l.ctx, &userModel.BlogUser{
		Id:       in.Id,
		Nickname: sql.NullString{String: in.Nickname, Valid: true},
	})
	if err != nil {
		return nil, err
	}
	return &user.EmptyResp{}, nil
}
