package api

import (
	"database/sql"
	"encoding/json"
	"github.com/eduardoraider/go-crud-sqlite/internal/entity"
	"net/http"
)

// SchemaSQL defines the schema for the books table
const SchemaSQL = `
CREATE TABLE IF NOT EXISTS books (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT,
    author TEXT
);
`

// NewRouter returns a new router instance
func NewRouter(db *sql.DB) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/books", getBooksHandler(db))
	mux.HandleFunc("/books/create", createBookHandler(db))
	mux.HandleFunc("/books/update", updateBookHandler(db))
	mux.HandleFunc("/books/delete", deleteBookHandler(db))
	return mux
}

func getBooksHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		books, err := entity.GetBooks(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(books)
	}
}

func createBookHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newBook entity.Book
		if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := entity.AddBook(db, newBook); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func updateBookHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var updatedBook entity.Book
		if err := json.NewDecoder(r.Body).Decode(&updatedBook); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := entity.UpdateBook(db, updatedBook); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func deleteBookHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var bookID struct {
			ID int `json:"id"`
		}
		if err := json.NewDecoder(r.Body).Decode(&bookID); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := entity.DeleteBook(db, bookID.ID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
