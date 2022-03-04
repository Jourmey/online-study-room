package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"online-study-room/internal/svc"
)

func HelloHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		httpx.OkJson(w, "hello online study room!")
	}
}
