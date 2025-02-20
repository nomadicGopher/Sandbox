// In production, if the program contains any special files like .wasm, it will need to be configured to allow fetching of that extension.
package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

var port int

func init() {
	flag.IntVar(&port, "port", 8080, "port to serve on")
	flag.Parse()
}

func main() {
	// Get the current working directory
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	// Serve static files from the current directory
	fs := http.FileServer(http.Dir(dir))
	http.Handle("/", fs)

	fmt.Printf("Starting server on :%d\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
