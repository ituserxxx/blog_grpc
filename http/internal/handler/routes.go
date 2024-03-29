// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	user "brl/http/internal/handler/user"
	"brl/http/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/from/:name",
				Handler: user.HttpHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/add",
				Handler: user.AddUserHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/del",
				Handler: user.DelUserHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/UpdateNickName",
				Handler: user.UserUpdateNickNameHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/info",
				Handler: user.GetUserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/list",
				Handler: user.GetUserListHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/api/user"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: user.UserLoginHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/user"),
	)
}
