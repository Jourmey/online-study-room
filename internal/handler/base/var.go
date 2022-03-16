package base

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"online-study-room/internal/types"
)

func Result(w http.ResponseWriter, msg interface{}, err error) {
	var r types.CommonResult
	if err != nil {
		r.Code = 1
		r.Msg = err.Error()
		r.Result = msg
	} else {
		r.Code = 0
		r.Msg = "success"
		r.Result = msg
	}

	httpx.OkJson(w, r)
}
