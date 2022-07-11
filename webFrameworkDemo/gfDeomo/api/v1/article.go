package v1

import "github.com/gogf/gf/v2/frame/g"

type ArticleIndexReq struct {
	g.Meta `path:"/article" method:"get" tags:"文章" summary:"展示Article列表页面"`
	ContentGetListCommonReq
}

type ContentGetListCommonReq struct {
	Type       string `json:"type"   in:"query" dc:"内容模型"`
	CategoryId uint   `json:"cate"   in:"query" dc:"栏目ID"`
	Sort       int    `json:"sort"   in:"query" dc:"排序类型(0:最新, 默认。1:活跃, 2:热度)"`
}

type ArticleIndexRes struct {
	g.Meta `mime:"application/json"`

	ArticleId string `json:"articleId"`
}
