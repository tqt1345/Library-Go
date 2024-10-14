package model

import (
	"database/sql"
	"log"
)

type Author struct {
	FirstName string
	LastName  string
	ID        int
}

type Book struct {
	Title       string
	Description string
	ID          int
}

type Repository struct {
	db *sql.DB
}

func NewRepo(d *sql.DB) *Repository {
	return &Repository{db: d}
}

func (r *Repository) FindBookById(id int) (Book, error) {
	sql := `SELECT * FROM books WHERE id = ?`

	var b Book
	err := r.db.QueryRow(sql, id).Scan(&b.ID, &b.Title, &b.Description)
	if err != nil {
		log.Print(err)
		return b, err
	}

	return b, nil
}
