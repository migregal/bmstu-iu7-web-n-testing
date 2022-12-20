package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"golang_benchmarks/handlers"
)

func BenchmarkHttpSimple(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	benchRequest(b, handlers.GetHandler("http"), req)
}

func BenchmarkHttpParam(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/gopher?name=gopher", nil)
	benchRequest(b, handlers.GetHandler("http"), req)
}

func BenchmarkHttpPostData(b *testing.B) {
	json_data, err := json.Marshal(generateRequest(10))
	if err != nil {
		b.Fatal(err)
	}

	req, _ := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(json_data))
	req.Header.Set("Content-Type", "application/json")
	benchRequest(b, handlers.GetHandler("http"), req)
}
