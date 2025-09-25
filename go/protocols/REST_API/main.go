package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux = setHandlers(mux)

	http.ListenAndServe(":9999", mux)
}

func setHandlers(mux *http.ServeMux) *http.ServeMux {
	mux.HandleFunc("GET /users", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("page") != "" {
			// pass ...?page=# in the request URL
			w.Write([]byte("users response page: " + r.URL.Query().Get("page")))
		} else {
			w.Write([]byte("users endpoint"))
		}
	})

	mux.HandleFunc("GET /users/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("users subtree endpoint"))
	})

	mux.HandleFunc("PUT /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("users dynamic path id: " + r.PathValue("id")))
	})

	mux.HandleFunc("PUT /users/jon", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("specific endpoint is prioritized"))
	})

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		w.Write([]byte("root endpoint"))
	})

	return mux
}
