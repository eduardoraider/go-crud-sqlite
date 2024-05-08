package main

import (
	"database/sql"
	"fmt"
	"github.com/eduardoraider/go-crud-sqlite/internal/api"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var err error

	// Open database connection
	db, err := sql.Open("sqlite3", "./db/books.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create table if not exists
	if _, err := db.Exec(api.SchemaSQL); err != nil {
		log.Fatal(err)
	}

	// Start server
	fmt.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", api.NewRouter(db)))
}
