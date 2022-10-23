package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"golabs/go-zero/bookstore/api/internal/logic"
	"golabs/go-zero/bookstore/api/internal/svc"
	"golabs/go-zero/bookstore/api/internal/types"
)

func AddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAddLogic(r.Context(), svcCtx)
		resp, err := l.Add(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
