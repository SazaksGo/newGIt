package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func InsertRow(ctx context.Context, conn *pgx.Conn) error {

	sqlQuery := ``

	_, err := conn.Exec(ctx, sqlQuery)
	if err != nil {
		return err
	}
}
