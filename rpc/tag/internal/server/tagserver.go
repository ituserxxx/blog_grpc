// Code generated by goctl. DO NOT EDIT.
// Source: tag.proto

package server

import (
	"context"

	"tag/internal/logic"
	"tag/internal/svc"
	"tag/tag"
)

type TagServer struct {
	svcCtx *svc.ServiceContext
	tag.UnimplementedTagServer
}

func NewTagServer(svcCtx *svc.ServiceContext) *TagServer {
	return &TagServer{
		svcCtx: svcCtx,
	}
}

func (s *TagServer) Ping(ctx context.Context, in *tag.Request) (*tag.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}