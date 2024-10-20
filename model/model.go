package model

import (
	"database/sql"
	"log"
	"strconv"
)

const (
	bookPath   = "/api/books/"
	authorPath = "/api/authors"
)

type Repository struct {
	db *sql.DB
}

func NewRepo(d *sql.DB) *Repository {
	return &Repository{db: d}
}

type Book struct {
	Title       string
	Description string
	ID          int
	CoverImage  string
	Link        string
}

func (b *Book) UpdateLink() {
	b.Link = bookPath + strconv.Itoa(b.ID)
}

func (r *Repository) FindBookById(id int) (Book, error) {
	sql := `SELECT * FROM books WHERE id = ?`

	var b Book
	err := r.db.QueryRow(sql, id).Scan(&b.ID, &b.Title, &b.Description, &b.CoverImage)
	log.Print(b)
	if err != nil {
		log.Print(err)
		return b, err
	}

	b.UpdateLink()

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
		err := rows.Scan(&b.ID, &b.Title, &b.Description, &b.CoverImage)
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

	for i := range books {
		books[i].UpdateLink()
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
		err := rows.Scan(&b.ID, &b.Title, &b.Description, &b.CoverImage)
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

	for i := range books {
		books[i].UpdateLink()
	}

	return books, nil
}

type Author struct {
	FirstName string
	LastName  string
	ID        int
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

func (r *Repository) FindAuthorsByFirstName(name string) ([]Author, error) {
	sql := `SELECT * FROM authors WHERE first_name LIKE '%' || ? || '%'`

	var authors []Author

	rows, err := r.db.Query(sql, name)
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
