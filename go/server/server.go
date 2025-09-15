// In production, if the program contains any special files like .wasm, it will need to be configured to allow fetching of that extension.
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	root := flag.String("root", "", "Directory to serve files from (default: current working directory)")
	ssl := flag.Bool("ssl", false, "Enable/disable SSL layer for the server.")
	certFile := flag.String("cert", "server.crt", "Path to the SSL certificate file.")
	keyFile := flag.String("key", "server.key", "Path to the SSL key file.")
	httpsPort := flag.Uint("https", 8443, "Port to serve over for HTTPS server.")
	httpPort := flag.Uint("http", 8080, "Port to serve over for HTTP server.")
	flag.Parse()

	if *root == "" {
		var err error
		*root, err = os.Getwd()
		if err != nil {
			fmt.Println("Error getting current directory:", err)
			return
		}
	}

	fmt.Println("Serving directory: ", *root)

	fs := http.FileServer(http.Dir(*root))
	http.Handle("/", fs)

	if *ssl {
		httpsPortStr := strconv.Itoa(int(*httpsPort))
		fmt.Println("HTTPS server starting at https://localhost:" + httpsPortStr)
		if err := http.ListenAndServeTLS(":"+httpsPortStr, *certFile, *keyFile, nil); err != nil {
			log.Fatalf("Error starting HTTPS server: %v", err)
		}
	} else {
		httpPortStr := strconv.Itoa(int(*httpPort))
		fmt.Println("HTTP server starting at http://localhost:" + httpPortStr)
		if err := http.ListenAndServe(":"+httpPortStr, nil); err != nil {
			log.Fatalf("Error starting HTTP server: %v", err)
		}
	}
}
