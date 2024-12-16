package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	fmt.Println("Base Server is Running")
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("you are at base server"))
	})
	http.ListenAndServe(":3000", r)
}
