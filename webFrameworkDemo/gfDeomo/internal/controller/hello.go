package controller

import (
	"context"
	v1 "gfDeomo/api/v1"

	"github.com/gogf/gf/v2/net/ghttp"

	"github.com/gogf/gf/v2/frame/g"
)

var (
	Hello = cHello{}
)

type cHello struct {
	test1 interface{}
}

func (c *cHello) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
	g.RequestFromCtx(ctx).Response.Writeln("Hello World!")
	return
}

func (c cHello) Test1(r *ghttp.Request) {
	g.Log("aaa").Info(r.Context(), "11111111111111")
	r.Response.WriteJson(g.Map{"abc": 123445})
}
