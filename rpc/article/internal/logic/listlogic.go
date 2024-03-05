package logic

import (
	"brl/rpc/article/article"
	"brl/rpc/article/internal/svc"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListLogic) List(in *article.ListReq) (*article.ListResp, error) {
	// todo: add your logic here and delete this line
	ls, total, err := l.svcCtx.ArticleModel.List(l.ctx, 1, 10)
	if err != nil {
		return nil, err
	}
	var ut = make([]*article.ListItem, 0)
	for _, v := range ls {
		ut = append(ut, &article.ListItem{
			Id:         v.Id,
			Title:      v.Title,
			Content:    v.Content,
			CreateTime: v.CreateTime.String(),
			//UpdateTime: string(v.UpdateTime),
			Status: v.Status,
		})
	}

	return &article.ListResp{
		Total: total,
	}, nil
}
