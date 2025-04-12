// https://docs.gofiber.io/template/html/TEMPLATES_CHEATSHEET/

package main

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
)

type Window struct {
	PageTitle string
}

func convertPortsToStrings(httpsPort *uint, httpPort *uint) (httpsPortStr string, httpPortStr string) {
	httpsPortStr = strconv.Itoa(int(*httpsPort))
	httpPortStr = strconv.Itoa(int(*httpPort))
	return
}

func setHandlers() {
	http.HandleFunc("/", func(responseWriter http.ResponseWriter, request *http.Request) {
		windowData := Window{PageTitle: "Hello World with Go and HTMX"}
		index := template.Must(template.ParseFiles("./index.html"))
		index.Execute(responseWriter, windowData)
	})

	http.HandleFunc("/update", func(responseWriter http.ResponseWriter, request *http.Request) {
		updateFilePath := "templates/update.html"
		updateContent, err := os.ReadFile(updateFilePath)
		if err != nil {
			http.Error(responseWriter, "Unable to read "+updateFilePath, http.StatusInternalServerError)
			return
		}
		responseWriter.Write(updateContent)
	})
}

func main() {
	ssl := flag.Bool("ssl", false, "Enable/disable SSL layer for the server.")
	certFile := flag.String("cert", "server.crt", "Path to the SSL certificate file.")
	keyFile := flag.String("key", "server.key", "Path to the SSL key file.")
	httpsPort := flag.Uint("https", 443, "Port to serve over for HTTPS server.")
	httpPort := flag.Uint("http", 80, "Port to serve over for HTTP server.")
	flag.Parse()

	setHandlers()

	httpsPortStr, httpPortStr := convertPortsToStrings(httpsPort, httpPort)

	if *ssl {
		fmt.Println("HTTPS server starting at http://localhost:", httpsPortStr)
		err := http.ListenAndServeTLS(":"+httpsPortStr, *certFile, *keyFile, nil)
		if err != nil {
			fmt.Println("Error starting HTTPS server:", err)
		}
	} else {
		fmt.Println("HTTP server starting at http://localhost:", httpPortStr)
		err := http.ListenAndServe(":"+httpPortStr, nil)
		if err != nil {
			fmt.Println("Error starting HTTP server:", err)
		}
	}
}
