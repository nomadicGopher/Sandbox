// https://docs.gofiber.io/template/html/TEMPLATES_CHEATSHEET/

package main

import (
	"flag"
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

		// Load & reference a single templates
		index := template.Must(template.ParseFiles("index.html"))
		if err := index.Execute(responseWriter, windowData); err != nil {
			http.Error(responseWriter, "Error rendering template", http.StatusInternalServerError)
		}

		// Load & reference multiple templates
		/*
			templates := template.Must(template.ParseFiles("header.html", "body.html", "footer.html"))
			if err := templates.ExecuteTemplate(responseWriter, "body.html", windowData); err != nil {
				http.Error(responseWriter, "Error rendering body template", http.StatusInternalServerError)
			}
		*/

		// Define an inline template using template.New and Parse
		/*
					inlineTemplate := template.Must(template.New("index").Parse(`
			            <!DOCTYPE html>
			            <html lang="en">
			            <head>
			                <meta charset="UTF-8">
			                <meta name="viewport" content="width=device-width, initial-scale=1.0">
			                <title>{{.PageTitle}}</title>
			            </head>
			            <body>
			                <h1>{{.PageTitle}}</h1>
			                <p>Welcome to the Go and HTMX example!</p>
			            </body>
			            </html>
			        `))

					// Render the inline template with the provided data
					if err := inlineTemplate.Execute(responseWriter, windowData); err != nil {
						http.Error(responseWriter, "Error rendering template", http.StatusInternalServerError)
					}
		*/
	})

	http.HandleFunc("/update", func(responseWriter http.ResponseWriter, request *http.Request) {
		injectableContentPath := "templates/update.html"
		injectableContent, err := os.ReadFile(injectableContentPath)
		if err != nil {
			http.Error(responseWriter, "Unable to read "+injectableContentPath, http.StatusInternalServerError)
			return
		}
		_, err = responseWriter.Write(injectableContent)
		if err != nil {
			log.Fatalf("Error writing injectable content to the server: %v", err)
		}
	})
}

func main() {
	port := flag.Int("port", 9999, "What port should be used to serve files?")
	flag.Parse()

	setHandlers()

	fmt.Printf("HTTP server starting at http://localhost:%d\n", *port)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil); err != nil {
		log.Fatalf("Error starting HTTP server: %v", err)
	}
}
