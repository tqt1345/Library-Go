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
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set(ContentType, Html)
	tmpl.Execute(w, h)
}

// END HELPER FUNCTIONS
