package handlers

import "net/http"

var (
	httpHandlers map[string]http.Handler
)

func RegisterHandler(name string, handler http.Handler) {
	if httpHandlers == nil {
		httpHandlers = make(map[string]http.Handler)
	}
	if _, ok := httpHandlers[name]; ok {
		panic("already registered")
	}
	httpHandlers[name] = handler
}

func GetHandler(name string) http.Handler {
	if httpHandlers == nil {
		return nil
	}
	handler, _ := httpHandlers[name]
	return handler
}
