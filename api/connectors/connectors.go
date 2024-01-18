package connectors

import (
	"fmt"
	"net/http"
)

func handleGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Rota GET")
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Rota POST")
}

func handleDelete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Rota DELETE")
}

func Router() {
	http.HandleFunc("/connectors/:connector/configs", handleGet)
	http.HandleFunc("/connectors/:connector", handleDelete)

	http.Handle("/connectors", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handleGet(w, r)
		case http.MethodPost:
			handlePost(w, r)
		default:
			http.NotFound(w, r)
		}
	}))
}
