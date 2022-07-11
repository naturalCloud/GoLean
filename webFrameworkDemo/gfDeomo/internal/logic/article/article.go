package article

import (
	"context"
	v1 "gfDeomo/api/v1"
)

func New() *sArticle {
	return &sArticle{}
}

type sArticle struct{}

// List 文章列表
func (a *sArticle) List(ctx context.Context, req v1.ContentGetListCommonReq) ([]v1.ArticleIndexRes, error) {
	return []v1.ArticleIndexRes{
		{
			ArticleId: "id",
		},
	}, nil
}
