package handler

import (
	"net/http"

	"github.com/nizonglonggit/usercenter/api/internal/logic"
	"github.com/nizonglonggit/usercenter/api/internal/svc"
	"github.com/nizonglonggit/usercenter/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func ApiHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewApiLogic(r.Context(), ctx)
		resp, err := l.Api(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
