package article

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BlogArticleModel = (*customBlogArticleModel)(nil)

type (
	// BlogArticleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBlogArticleModel.
	BlogArticleModel interface {
		blogArticleModel
		List(ctx context.Context, page, limit int) ([]*BlogArticle, int64, error)
		withSession(session sqlx.Session) BlogArticleModel
	}

	customBlogArticleModel struct {
		*defaultBlogArticleModel
	}
)

// NewBlogArticleModel returns a model for the database table.
func NewBlogArticleModel(conn sqlx.SqlConn) BlogArticleModel {
	return &customBlogArticleModel{
		defaultBlogArticleModel: newBlogArticleModel(conn),
	}
}

func (m *customBlogArticleModel) withSession(session sqlx.Session) BlogArticleModel {
	return NewBlogArticleModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customBlogArticleModel) List(ctx context.Context, page, limit int) ([]*BlogArticle, int64, error) {
	query := fmt.Sprintf("select %s from %s limit ?,?", blogArticleRows, m.table)
	var resp []*BlogArticle

	err := m.conn.QueryRowsCtx(ctx, &resp, query, (page-1)*limit, limit)
	switch err {
	case nil:
		break
	case sqlc.ErrNotFound:
		return nil, 0, ErrNotFound

	}
	qc := fmt.Sprintf("select count(*) as count from %s", m.table)
	var count int64
	err = m.conn.QueryRow(&count, qc)

	switch err {
	case nil:
		return resp, count, nil
	case sqlc.ErrNotFound:
		return resp, 0, ErrNotFound
	default:
		return resp, 0, err
	}

}
