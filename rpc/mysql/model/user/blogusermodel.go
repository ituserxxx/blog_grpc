package user

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ BlogUserModel = (*customBlogUserModel)(nil)

type (
	// BlogUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBlogUserModel.
	BlogUserModel interface {
		blogUserModel
	}

	customBlogUserModel struct {
		*defaultBlogUserModel
	}
)

// NewBlogUserModel returns a model for the database table.
func NewBlogUserModel(conn sqlx.SqlConn) BlogUserModel {
	return &customBlogUserModel{
		defaultBlogUserModel: newBlogUserModel(conn),
	}
}
