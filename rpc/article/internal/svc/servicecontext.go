package svc

import (
	"brl/rpc/article/internal/config"
	"brl/rpc/mysql/model/article"
	_ "brl/rpc/mysql/model/article"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config       config.Config
	ArticleModel article.BlogArticleModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.Datasource)
	return &ServiceContext{
		Config:       c,
		ArticleModel: article.NewBlogArticleModel(sqlConn),
	}
}
