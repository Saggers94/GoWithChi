package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	// Because we have set this with Get method, It will only accept GET requests
	router.Get("/hello", basicHandler) // We can just use the basicHandler function
	server := &http.Server{
		Addr:    ":3000",
		Handler: router, // We can just pass router because the chi router conforms to the http.Server handler
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Failed to listen to server: ", err)
	}
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
