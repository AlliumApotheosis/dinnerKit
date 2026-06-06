package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func GetDBConnection() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres@localhost:5432/dinnerkitdb")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}

	return conn, nil
}

func GetDinnerKitTables(conn *pgx.Conn) ([]string, error) {
	rows, err := conn.Query(context.Background(), "SELECT table_name FROM information_schema.tables WHERE table_schema = 'public' AND table_type = 'BASE TABLE';")
	tables, err := pgx.CollectRows(rows, pgx.RowTo[string])
	return tables, err
}
