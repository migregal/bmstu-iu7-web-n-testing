package main

import (
	"fmt"
	"net/http"
	"testing"
)

type Req struct {
	Data []string `json:"data"`
}

func generateRequest(len int) Req {
	data := make([]string, len)

	for i := 0; i < len; i++ {
		data[i] = fmt.Sprintf("data string %d", i)
	}

	return Req{Data: data}
}

type mockResponseWriter struct{}

func (m *mockResponseWriter) Header() http.Header {
	return http.Header{}
}

func (m *mockResponseWriter) Write(p []byte) (int, error) {
	return len(p), nil
}

func (m *mockResponseWriter) WriteHeader(code int) {}

func benchRequest(b *testing.B, router http.Handler, r *http.Request) {
	w := mockResponseWriter{}
	u := r.URL
	r.RequestURI = u.RequestURI()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		router.ServeHTTP(&w, r)

		// clear caches
		r.Form = nil
		r.PostForm = nil
		r.MultipartForm = nil
	}
}
