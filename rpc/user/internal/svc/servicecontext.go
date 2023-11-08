package svc

import (
	"brl/rpc/user/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)
import "brl/rpc/mysql/model/user"

type ServiceContext struct {
	Config config.Config
	BlogUserModel     user.BlogUserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.Datasource)
	return &ServiceContext{
		Config: c,
		BlogUserModel:     user.NewBlogUserModel(sqlConn),
	}
}
