package main

import (
	"encoding/json"
	"fmt"
	"golang_benchmarks/handlers"
	"net/http"

	"goji.io"
	"goji.io/pat"
)

func init() {
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/"), func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, "Hello, World")
	})
	mux.HandleFunc(pat.Post("/"), func(w http.ResponseWriter, r *http.Request) {
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
	mux.HandleFunc(pat.Get("/:name"), func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello, %s", pat.Param(r, "name"))
	})

	handlers.RegisterHandler("goji", mux)
}
