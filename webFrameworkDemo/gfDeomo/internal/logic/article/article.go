package article

import (
	"gfDeomo/internal/model/entity"
	"gfDeomo/internal/service"

	"github.com/gogf/gf/v2/net/ghttp"
)

type sArticle struct{}

var articleIns = New()

func New() *sArticle {
	return &sArticle{}
}

func init() {
	service.RegisterArticle(articleIns)
}

// Abc    aaa
func (a *sArticle) Index(r *ghttp.Request, articles entity.Articles) bool {
	return false
}
