package postgres

import (
	_ "github.com/lib/pq"
	"github.com/gofaith/go-zero/core/stores/sqlx"
)

const postgreDriverName = "postgres"

func NewPostgre(datasource string, opts ...sqlx.SqlOption) sqlx.SqlConn {
	return sqlx.NewSqlConn(postgreDriverName, datasource, opts...)
}
