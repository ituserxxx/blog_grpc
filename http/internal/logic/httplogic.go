package logic

import (
	"context"

	"brl/http/internal/types"
	"brl/http/internal/svc"
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
	// todo: add your logic here and delete this line
	// l.svcCtx.UserRpc.Ping()
	return
}
