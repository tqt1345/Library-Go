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

func FindAll() {
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

func (r *Repository) FindBooksByTitle(title string) ([]Book, error) {
	sql := `SELECT * FROM books WHERE title LIKE '%' || ? || '%'`

	var books []Book

	rows, err := r.db.Query(sql, title)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var b Book
		err := rows.Scan(&b.ID, &b.Title, &b.Description)
		if err != nil {
			log.Print(err)
			return nil, err
		}
		books = append(books, b)
	}

	if err := rows.Err(); err != nil {
		log.Print(err)
		return nil, err
	}

	return books, nil
}

func (r *Repository) FindAllBooks() ([]Book, error) {
	sql := `SELECT * FROM books`

	var books []Book

	rows, err := r.db.Query(sql)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var b Book
		err := rows.Scan(&b.ID, &b.Title, &b.Description)
		if err != nil {
			log.Print(err)
			return nil, err
		}
		books = append(books, b)
	}

	if err := rows.Err(); err != nil {
		log.Print(err)
		return nil, err
	}

	return books, nil
}

func (r *Repository) FindAllAuthors() ([]Author, error) {
	sql := `SELECT * FROM authors`

	var authors []Author

	rows, err := r.db.Query(sql)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var a Author
		err := rows.Scan(&a.ID, &a.FirstName, &a.LastName)
		if err != nil {
			log.Print(err)
			return nil, err
		}
		authors = append(authors, a)
	}

	if err := rows.Err(); err != nil {
		log.Print(err)
		return nil, err
	}

	return authors, nil
}

// func (r *Repository) FindAuthorsByFirstName(name string) ([]Author, error) {
//
// }
