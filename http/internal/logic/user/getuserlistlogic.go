package user

import (
	"brl/http/internal/svc"
	"brl/http/internal/types"
	"brl/rpc/user/userclient"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserListLogic) GetUserList(req *types.UserListReq) (resp *types.UserListResp, err error) {

	ls, err := l.svcCtx.UserRpc.UserList(l.ctx, &userclient.UserListReq{})
	if err != nil {
		return nil, err
	}
	var rl = make([]*types.UserListItem, 0)
	for _, item := range ls.List {
		rl = append(rl, &types.UserListItem{
			Id:       item.Id,
			Username: item.Username,
			Nickname: item.Nickname,
		})
	}
	return &types.UserListResp{
		Total: ls.Total,
		List:  rl,
	}, nil
}
