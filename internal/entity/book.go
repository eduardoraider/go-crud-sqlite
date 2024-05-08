package entity

import (
	"database/sql"
)

// Book represents a book entity
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// GetBooks retrieves all books from the database
func GetBooks(db *sql.DB) ([]Book, error) {
	rows, err := db.Query("SELECT id, title, author FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author); err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

// AddBook adds a new book to the database
func AddBook(db *sql.DB, book Book) error {
	_, err := db.Exec("INSERT INTO books (title, author) VALUES (?, ?)", book.Title, book.Author)
	if err != nil {
		return err
	}
	return nil
}

// UpdateBook updates an existing book in the database
func UpdateBook(db *sql.DB, book Book) error {
	_, err := db.Exec("UPDATE books SET title = ?, author = ? WHERE id = ?", book.Title, book.Author, book.ID)
	if err != nil {
		return err
	}
	return nil
}

// DeleteBook deletes a book from the database by its ID
func DeleteBook(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM books WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
