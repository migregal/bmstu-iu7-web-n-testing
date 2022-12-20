package main

import (
	"encoding/json"
	"fmt"
	"golang_benchmarks/handlers"
	"net/http"

	gorilla "github.com/gorilla/mux"
)

func init() {
	h := gorilla.NewRouter()
	h.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Hello, World")
			return
		}

		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		decoder := json.NewDecoder(r.Body)
		var req Req
		if err := decoder.Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if len(req.Data) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, req.Data[len(req.Data)-1])
	})
	h.HandleFunc("/{name}", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello, %s", gorilla.Vars(r)["name"])
	})
	handlers.RegisterHandler("gorilla", h)
}
