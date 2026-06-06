package main

import (
	"context"
	"fmt"
	"net/http"
	//"github.com/jackc/pgx/v5"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		conn, err := GetDBConnection()
		if err != nil {
			fmt.Fprintf(w, "Error connecting to database: %v", err)
			return
		}
		defer conn.Close(context.Background())

		tables, err := GetDinnerKitTables(conn)
		if err != nil {
			fmt.Fprintf(w, "Error querying tables: %v", err)
			return
		}
		fmt.Fprintf(w, "tables:\n %v", tables)

		//fmt.Fprintf(w, "Hello, Foxer! welcome to dinnerKit!")
	})
	http.ListenAndServe(":6336", nil)
}
