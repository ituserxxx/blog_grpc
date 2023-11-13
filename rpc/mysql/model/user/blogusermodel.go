package user

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BlogUserModel = (*customBlogUserModel)(nil)

type (
	// BlogUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBlogUserModel.
	BlogUserModel interface {
		blogUserModel
		// 此处自增model方法
		List(ctx context.Context, page, limit int) ([]*BlogUser, error)
		Conner(ctx context.Context) sqlx.SqlConn
	}

	customBlogUserModel struct {
		*defaultBlogUserModel
	}
)

func (c *customBlogUserModel) List(ctx context.Context, page, limit int) ([]*BlogUser, error) {
	query := fmt.Sprintf("select %s from %s limit ?,?", blogUserRows, c.table)
	var resp []*BlogUser
	err := c.conn.QueryRowsCtx(ctx, &resp, query, (page-1)*limit, limit)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

/*
事务操作

	err := conn.TransactCtx(context.Background(), func(ctx context.Context, session sqlx.Session) error {
	       r, err := session.ExecCtx(ctx, "insert into user (id, name) values (?, ?)", 1, "test")
	       if err != nil {
	           return err
	       }
	       r ,err =session.ExecCtx(ctx, "insert into user (id, name) values (?, ?)", 2, "test")
	       if err != nil {
	           return err
	       }
	   })
*/
func (c *customBlogUserModel) Conner(ctx context.Context) sqlx.SqlConn {
	return c.conn
}

// NewBlogUserModel returns a model for the database table.
func NewBlogUserModel(conn sqlx.SqlConn) BlogUserModel {
	return &customBlogUserModel{
		defaultBlogUserModel: newBlogUserModel(conn),
	}
}
