package main

import (
	"log"
	"net/http"
)

func main() {
	mux1 := http.NewServeMux()
	mux2 := http.NewServeMux()

	mux1.HandleFunc("GET /users", usersHandler)
	mux1.HandleFunc("GET /users/", usersSubtreeHandler)
	mux1.HandleFunc("PUT /users/{id}", usersIDHandler)
	mux1.HandleFunc("PUT /users/jon", usersJonHandler)
	mux1.HandleFunc("GET /", rootHandler)

	mux2.HandleFunc("GET /users", usersHandler)
	mux2.HandleFunc("GET /users/", usersSubtreeHandler)
	mux2.HandleFunc("PUT /users/{id}", usersIDHandler)
	mux2.HandleFunc("PUT /users/jon", usersJonHandler)
	mux2.HandleFunc("GET /", rootHandler)

	// Start both servers on different ports
	go func() {
		log.Println("Starting mux1 on :9999")
		http.ListenAndServe(":9999", mux1)
	}()
	log.Println("Starting mux2 on :8888")
	http.ListenAndServe(":8888", mux2)
}

// Handler Functions

func usersHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("page") != "" {
		// pass ...?page=# in the request URL
		w.Write([]byte("users response page: " + r.URL.Query().Get("page")))
	} else {
		w.Write([]byte("users endpoint"))
	}
}

func usersSubtreeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("users subtree endpoint"))
}

func usersIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("users dynamic path id: " + r.PathValue("id")))
}

func usersJonHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("specific endpoint is prioritized"))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("root endpoint"))
}
