package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

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
}
