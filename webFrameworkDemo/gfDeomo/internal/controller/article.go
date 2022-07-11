package controller

import (
	v1 "gfDeomo/api/v1"
	"gfDeomo/internal/service"
	"github.com/gogf/gf/v2/net/ghttp"
)

var (
	Article = new(article)
)

type article struct{}

func (a *article) Index(r *ghttp.Request) {

	var req v1.ContentGetListCommonReq

	if errrrHandle(r, r.Parse(&req)) {
		return
	}

	res, err := service.Article().List(r.Context(), req)

	if errrrHandle(r, err) {
		return
	}
	r.Response.WriteJson(res)
}

func errrrHandle(r *ghttp.Request, err error) (isReturn bool) {
	if err != nil {
		r.SetError(err)
		return true
	}
	return false
}
