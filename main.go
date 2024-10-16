package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"github.com/tqt1345/Library-Go/controller"
	"github.com/tqt1345/Library-Go/model"
)

const (
	JsonMime    = "application/json"
	HtmlMime    = "text/html"
	ContentType = "Content-Type"
)

var (
	repo *model.Repository
	ctx  context.Context
)

// This exists to init everything the program needs in one area and return any errors
func run() error {
	ctx = context.Background()

	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	log.Print("Database connected...")

	repo = model.NewRepo(db)
	log.Print("Repository loaded...")

	port := ":8080"

	controller.Init(repo)
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
