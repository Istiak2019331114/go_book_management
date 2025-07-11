package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"go_book_management/handlers"
	"go_book_management/storage"
	"log"
	"net/http"
)

func main() {
	storage.Init()
	r := chi.NewRouter()
	r.Route("/books", func(r chi.Router) {
		r.Get("/", handlers.GetBooks)
		r.Get("/{id}", handlers.GetBook)
		r.Post("/", handlers.AddBook)
		r.Put("/{id}", handlers.UpdateBook)
		r.Delete("/{id}", handlers.DeleteBook)
	})
	r.Route("/authors", func(r chi.Router) {
		r.Get("/", handlers.GetAuthors)
		r.Get("/{id}", handlers.GetAuthor)
	})
	r.Get("/find/{genre}", handlers.FindByGenre)

	fmt.Println("Listening on :8000")
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatal(err)
	}
}
