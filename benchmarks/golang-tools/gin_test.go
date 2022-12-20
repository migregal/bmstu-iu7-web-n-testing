package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"golang_benchmarks/handlers"
)

func BenchmarkGinSimple(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	benchRequest(b, handlers.GetHandler("gin"), req)
}

func BenchmarkGinParam(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/gopher", nil)
	benchRequest(b, handlers.GetHandler("gin"), req)
}

func BenchmarkGinPostData(b *testing.B) {
	json_data, err := json.Marshal(generateRequest(10))
	if err != nil {
		b.Fatal(err)
	}

	req, _ := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(json_data))
	req.Header.Set("Content-Type", "application/json")
	benchRequest(b, handlers.GetHandler("gin"), req)
}
