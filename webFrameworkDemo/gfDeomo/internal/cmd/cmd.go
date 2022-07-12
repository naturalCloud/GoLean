package cmd

import (
	"context"
	"fmt"
	"gfDeomo/internal/middleware"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"gfDeomo/internal/controller"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()

			s.Use(middleware.ExceptionHandle(ctx))
			s.Use(ghttp.MiddlewareHandlerResponse)

			s.Group("/hello", func(group *ghttp.RouterGroup) {
				group.Bind(
					controller.Hello)
			})

			s.Group("/article", func(group *ghttp.RouterGroup) {
				group.Bind(
					controller.Article)
			})

			fmt.Println(11111)
			s.EnableAdmin()
			s.Run()
			return nil
		},
	}
)
