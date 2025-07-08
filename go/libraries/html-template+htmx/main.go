// https://docs.gofiber.io/template/html/TEMPLATES_CHEATSHEET/

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Window struct {
	PageTitle string
}

func setHandlers() {
	http.HandleFunc("/", func(responseWriter http.ResponseWriter, request *http.Request) {
		windowData := Window{PageTitle: "Hello World with Go and HTMX"}
		index := template.Must(template.ParseFiles("./index.html"))
		if err := index.Execute(responseWriter, windowData); err != nil {
			http.Error(responseWriter, "Error rendering template", http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/update", func(responseWriter http.ResponseWriter, request *http.Request) {
		updateFilePath := "templates/update.html"
		updateContent, err := os.ReadFile(updateFilePath)
		if err != nil {
			http.Error(responseWriter, "Unable to read "+updateFilePath, http.StatusInternalServerError)
			return
		}
		_, err = responseWriter.Write(updateContent)
		if err != nil {
			log.Fatalf("Error writing var(updateContent) to the server: %v", err)
		}
	})
}

func main() {
	setHandlers()

	fmt.Println("HTTP server starting at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting HTTP server: %v", err)
	}
}
