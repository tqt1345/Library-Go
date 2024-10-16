package controller

import (
	"html/template"
	"log"
	"net/http"

	"github.com/tqt1345/Library-Go/model"
)

type Header struct {
	Heading  string
	NavItems []NavItem
}

type NavItem struct {
	Title string
	Url   string
}

type NavItems []NavItem

func (n *NavItems) Add(title, url string) {
	// *n = append(*n, NavItem{Title: title, Url: title})
	*n = append(*n, NavItem{Title: title, Url: url})
}

// HEADER TEMPLATES

func BookCatalogueHeaderTemplate(w http.ResponseWriter, r *http.Request) {
	serveHeaderTemplate(w, "Book Catalogue")
}

func IndexHeaderTemplate(w http.ResponseWriter, r *http.Request) {
	serveHeaderTemplate(w, "Home")
}

// END HEADER TEMPLATES

// CONTENT TEMPLATES

func AllBooksTemplate(w http.ResponseWriter, r *http.Request) {
	books, err := repo.FindAllBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Print(err.Error())
		return
	}

	if books == nil {
		books = []model.Book{}
	}

	tmpl, err := template.ParseFiles(wd + "/view/templates/catalogue-table.html")
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(ContentType, Html)
	tmpl.Execute(w, books)
}

// END CONTENT TEMPLATES
