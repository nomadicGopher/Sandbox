// https://docs.gofiber.io/template/html/TEMPLATES_CHEATSHEET/

package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
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

func convertPortsToStrings(httpsPort *uint, httpPort *uint) (httpsPortStr string, httpPortStr string) {
	httpsPortStr = strconv.Itoa(int(*httpsPort))
	httpPortStr = strconv.Itoa(int(*httpPort))
	return
}

func startServer(ssl bool, httpsPortStr, httpPortStr, certFile, keyFile string) {
	if ssl {
		fmt.Println("HTTPS server starting at https://localhost:" + httpsPortStr)
		if err := http.ListenAndServeTLS(":"+httpsPortStr, certFile, keyFile, nil); err != nil {
			log.Fatalf("Error starting HTTPS server: %v", err)
		}
	} else {
		fmt.Println("HTTP server starting at http://localhost:" + httpPortStr)
		if err := http.ListenAndServe(":"+httpPortStr, nil); err != nil {
			log.Fatalf("Error starting HTTP server: %v", err)
		}
	}
}

func main() {
	ssl := flag.Bool("ssl", false, "Enable/disable SSL layer for the server.")
	certFile := flag.String("cert", "server.crt", "Path to the SSL certificate file.")
	keyFile := flag.String("key", "server.key", "Path to the SSL key file.")
	httpsPort := flag.Uint("https", 8443, "Port to serve over for HTTPS server.")
	httpPort := flag.Uint("http", 8080, "Port to serve over for HTTP server.")
	flag.Parse()

	setHandlers()

	httpsPortStr, httpPortStr := convertPortsToStrings(httpsPort, httpPort)

	startServer(*ssl, httpsPortStr, httpPortStr, *certFile, *keyFile)
}
