// https://docs.gofiber.io/template/html/TEMPLATES_CHEATSHEET/

package main

import (
	"html/template"
	"net/http"
)

type PageData struct {
	Title string
}

func main() {
	// Parse HTML template
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := PageData{Title: "Hello World with Go and HTMX"}
		tmpl.Execute(w, data)
	})

	// Endpoint for HTMX interaction
	http.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<p><strong>Dynamic Update:</strong> Hello from HTMX!</p>"))
	})

	// Start server
	http.ListenAndServe(":8080", nil)
}
