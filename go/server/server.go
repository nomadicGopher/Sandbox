// In production, if the program contains any special files like .wasm, it will need to be configured to allow fetching of that extension.
package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

var (
	port int
	root string
)

func init() {
	flag.IntVar(&port, "port", 8080, "Port to serve on")
	flag.StringVar(&root, "root", "", "Directory to serve files from (default: current working directory)")
	flag.Parse()
}

func main() {
	if root == "" {
		var err error
		root, err = os.Getwd()
		if err != nil {
			fmt.Println("Error getting current directory:", err)
			return
		}
	}

	fs := http.FileServer(http.Dir(root))
	http.Handle("/", fs)

	fmt.Printf("Serving files from %s on :%d\n", root, port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
