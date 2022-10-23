package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"golabs/go-zero/shorturl/api/internal/logic"
	"golabs/go-zero/shorturl/api/internal/svc"
	"golabs/go-zero/shorturl/api/internal/types"
)

func ExpandHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ExpandReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewExpandLogic(r.Context(), svcCtx)
		resp, err := l.Expand(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
