package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()

	router.Get("/index", handler)
	server := &http.Server{
		Addr:    ":4000",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Failed to lister to server", err)
	}
	fmt.Println()
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
