package controller

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"

	"gfDeomo/api/v1"
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
