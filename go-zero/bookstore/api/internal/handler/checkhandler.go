package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"golabs/go-zero/bookstore/api/internal/logic"
	"golabs/go-zero/bookstore/api/internal/svc"
	"golabs/go-zero/bookstore/api/internal/types"
)

func CheckHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CheckReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCheckLogic(r.Context(), svcCtx)
		resp, err := l.Check(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
