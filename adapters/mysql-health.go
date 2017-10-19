package main

import (
	"database/sql"
	"flag"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var (
	user   = flag.String("user", "", "The database user name")
	passwd = flag.String("password", "", "The database password")
	db     = flag.String("database", "", "The database to connect to")
	query  = flag.String("query", "", "The test query")
	addr   = flag.String("address", "localhost:8080", "The address to listen on")
)

// Basic usage:
//   db-check --query="SELECT * from my-cool-table" \
//            --user=bdburns \
//            --passwd="you wish"
//
func main() {
	flag.Parse()
	db, err := sql.Open("localhost", fmt.Sprintf("%s:%s@/%s", *user, *passwd, *db))
	if err != nil {
		fmt.Printf("Error opening database: %v", err)
	}

	// Simple web handler that runs the query
	http.HandleFunc("", func(res http.ResponseWriter, req *http.Request) {
		_, err := db.Exec(*query)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			res.Write([]byte(err.Error()))
			return
		}
		res.WriteHeader(http.StatusOK)
		res.Write([]byte("OK"))
		return
	})
	// Startup the server
	http.ListenAndServe(*addr, nil)
}
