package main

import (
	"GITpro/feature_postgres/simple_connection"
	"GITpro/feature_postgres/simple_sql"
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	conn, err := simple_connection.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}
	if err := simple_sql.CreateTable(ctx, conn); err != nil {
		panic(err)
	}

	if err := simple_sql.InsertRow(ctx, conn); err != nil {
		panic(err)
	}
	fmt.Println("успех")
}
