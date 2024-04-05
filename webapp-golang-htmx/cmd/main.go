package main

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func init() {

}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "index.html", nil)
	})

	http.ListenAndServe(":3000", nil)
}
