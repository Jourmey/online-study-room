package user

import (
	"net/http"
	"online-study-room/internal/handler/base"
	"online-study-room/internal/logic/user"
	"online-study-room/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserSigninHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			UserName string `json:"user_name"`
			Password string `json:"password"`
		}
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewLogic(r.Context(), ctx)
		code, err := l.Signin(req.UserName, req.Password)
		base.Result(w, code, err)
	}
}

func UserRegisterHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			UserName string `json:"user_name"`
			Password string `json:"password"`
			NikeName string `json:"nike_name"`
		}
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewLogic(r.Context(), ctx)
		code, err := l.Register(req.UserName, req.Password, req.NikeName)
		base.Result(w, code, err)
	}
}

func UserByIdHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Id int `path:"id"`
		}

		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewLogic(r.Context(), ctx)
		resp, err := l.User(req.Id)
		base.Result(w, resp, err)
	}
}
