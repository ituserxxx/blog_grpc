package logic

import (
	"brl/rpc/user/internal/svc"
	"brl/rpc/user/user"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserListLogic) UserList(in *user.UserListReq) (*user.UserListResp, error) {

	ls, err := l.svcCtx.ModelUser.List(l.ctx, 1, 10)
	if err != nil {
		return nil, err
	}
	var ut = make([]*user.UserListItem, 0)
	for _, u1 := range ls {
		ut = append(ut, &user.UserListItem{
			Id:       u1.Id,
			Nickname: u1.Nickname.String,
			Username: u1.Username,
		})
	}
	return &user.UserListResp{
		Total: 10,
		List:  ut,
	}, nil
}
