package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
	"github.com/tqt1345/Library-Go/database"
)

const (
	JsonMime    = "application/json"
	HtmlMime    = "text/html"
	ContentType = "Content-Type"
)

var (
	queries *database.Queries
	ctx     context.Context
	ddl     string
)

func ApiIndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JsonMime)

	json.NewEncoder(w).Encode("Hello world!")
}

func ApiAllBooksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, JsonMime)
	books, err := queries.FindAllBooks(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Print(err)
	}

	if books == nil {
		books = []database.Book{}
	}
	json.NewEncoder(w).Encode(books)
}

func ApiBookByIdHandler(w http.ResponseWriter, r *http.Request) {
	strId := r.PathValue("id")
	i, err := strconv.Atoi(strId)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id := int64(i)

	book, err := queries.FindBookById(ctx, id)
	if err != nil {
		notFound := fmt.Sprintf("No book found with id: %d ", id)
		log.Print(notFound)
		http.Error(w, notFound, http.StatusNotFound)
		return
	}

	w.Header().Set(ContentType, JsonMime)
	json.NewEncoder(w).Encode(book)
}

func ApiBookByTitleHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	var title string

	if params["title"] == nil {
		log.Print("Bad Request")
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	title = params.Get("title")

	books, err := queries.FindBookByTitle(ctx, sql.NullString{String: title, Valid: true})
	if err != nil {
		log.Print("Not Found")
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	w.Header().Set(ContentType, JsonMime)
	json.NewEncoder(w).Encode(books)
}

func ApiBookByAuthorHandler(w http.ResponseWriter, r *http.Request) {
}

func run() error {
	ctx = context.Background()

	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	// if _, err := db.ExecContext(ctx, ddl); err != nil {
	// 	return err
	// }

	log.Print("Database connected...")

	queries = database.New(db)
	log.Print("Queries loaded...")

	port := ":8080"

	http.HandleFunc("GET /api/", ApiIndexHandler)
	http.HandleFunc("GET /api/books/all", ApiAllBooksHandler)
	http.HandleFunc("GET /api/books/{id}", ApiBookByIdHandler)
	http.HandleFunc("GET /api/books/title", ApiBookByTitleHandler)
	log.Print("Routes loaded...")

	log.Printf("Server started on port%s", port)
	http.ListenAndServe(port, nil)
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
