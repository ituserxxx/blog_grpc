package handler

import (
	"net/http"

	"brl/http/internal/logic"
	"brl/http/internal/svc"
	"brl/http/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewAddUserLogic(r.Context(), svcCtx)
		resp, err := l.AddUser(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
