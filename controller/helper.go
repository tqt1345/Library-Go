package controller

import (
	"html/template"
	"log"
	"net/http"
)

// HELPER FUNCTIONS

func servePage(w http.ResponseWriter, r *http.Request, filePath string) {
	w.Header().Set(ContentType, Html)
	http.ServeFile(w, r, wd+filePath)
}

func serveHeaderTemplate(w http.ResponseWriter, title string) {
	h := Header{title, nv}

	tmpl, err := template.ParseFiles(wd + "/view/templates/header.html")
	if err != nil {
		internalServerError(w, err)
		return
	}

	w.Header().Set(ContentType, Html)
	tmpl.Execute(w, h)
}

func internalServerError(w http.ResponseWriter, err error) {
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	log.Print(err.Error())
}

func badRequest(w http.ResponseWriter, err error) {
	http.Error(w, "Bad Request", http.StatusBadRequest)
	log.Print(err.Error())
}

func notFound(w http.ResponseWriter, err error) {
	http.Error(w, "Not Found", http.StatusNotFound)
	log.Print(err.Error())
}

// END HELPER FUNCTIONS
