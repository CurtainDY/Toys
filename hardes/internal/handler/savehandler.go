package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"hardes/internal/logic"
	"hardes/internal/svc"
	"hardes/internal/types"
)

func saveHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SaveReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewSaveLogic(r.Context(), svcCtx)
		resp, err := l.Save(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
