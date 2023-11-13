package user

import (
	"net/http"

	"brl/http/internal/logic/user"
	"brl/http/internal/svc"
	"brl/http/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DelUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DelUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewDelUserLogic(r.Context(), svcCtx)
		resp, err := l.DelUser(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
