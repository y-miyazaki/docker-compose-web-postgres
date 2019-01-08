package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// TestResponse responses common json data.
type TestResponse struct {
	Message string `json:"message,omitempty"`
}

func main() {
	r := chi.NewRouter()
	InitializeRouter(r)
	Router(r)
	http.ListenAndServe(":8080", r)
}

// InitializeRouter initializes Mux and middleware
func InitializeRouter(r *chi.Mux) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
}

// Router sets router.
func Router(r *chi.Mux) {
	r.HandleFunc("/", test)
}

// test function.
func test(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test")
	response := TestResponse{
		Message: "test",
	}
	ResponseJSON(w, 200, response)
}

// ResponseJSON function response as json with ResponseWriter
func ResponseJSON(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		json, _ := json.Marshal(data)
		_, _ = w.Write(json)
	}
}
