package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"mall/order/api/internal/logic"
	"mall/order/api/internal/svc"
	"mall/order/api/internal/types"
)

func findUserLikeHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserLikeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewFindUserLikeLogic(r.Context(), ctx)
		resp, err := l.FindUserLike(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
