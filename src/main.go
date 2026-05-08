package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		conn, err := GetDBConnection()
		if err != nil {
			fmt.Fprintf(w, "Error connecting to database: %v", err)
			return
		}
		defer conn.Close(context.Background())

		rows, err := conn.Query(context.Background(), "SELECT table_name FROM information_schema.tables WHERE table_schema = 'public' AND table_type = 'BASE TABLE';")
		tables, err := pgx.CollectRows(rows, pgx.RowTo[string])
		fmt.Fprintf(w, "tables:\n %v", tables)
		//fmt.Fprintf(w, "Hello, Foxer! welcome to dinnerKit!")
	})
	http.ListenAndServe(":6336", nil)
}
