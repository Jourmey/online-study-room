// Code generated by goctl. DO NOT EDIT.
package types

type CommonResult struct {
	Code   int         `json:"code"`
	Result interface{} `json:"result"`
	Msg    string      `json:"msg"`
}

type GetRequest struct {
	Id int `path:"id"`
}
