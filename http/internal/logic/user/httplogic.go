package user

import (
	"brl/rpc/user/userclient"
	"context"

	"brl/http/internal/svc"
	"brl/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HttpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHttpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HttpLogic {
	return &HttpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HttpLogic) Http(req *types.Request) (resp *types.Response, err error) {
	res, err := l.svcCtx.UserRpc.Ping(l.ctx, &userclient.Request{})
	if err != nil {
		return nil, err
	}
	return &types.Response{Message: res.Pong}, nil
}
