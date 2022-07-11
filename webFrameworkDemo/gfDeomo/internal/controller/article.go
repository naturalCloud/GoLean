package controller

import (
	"context"
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

	if errorHandle(r, r.Parse(&req)) {
		return
	}

	res, err := service.Article().List(r.Context(), req)

	if errorHandle(r, err) {
		return
	}
	r.Response.WriteJson(res)
}

func (a *article) Index2(ctx context.Context, req *v1.ArticleIndexReq) (v1.ArticleIndexRes, error) {
	list, err := service.Article().List(ctx, v1.ContentGetListCommonReq{})
	return list, err
}

func errorHandle(r *ghttp.Request, err error) (isReturn bool) {
	if err != nil {
		r.SetError(err)
		return true
	}
	return false
}
