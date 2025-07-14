package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go_book_management/auth"
	"go_book_management/handlers"
	"go_book_management/storage"
	"log"
	"net/http"
)

func main() {
	storage.Init()
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Public route
	r.Post("/login", auth.Login)

	// Unprotected routes
	r.Route("/books", func(r chi.Router) {
		r.Get("/", handlers.GetBooks)
		r.Get("/{id}", handlers.GetBook)

		// Protected routes
		r.With(auth.JWTMiddleware).Post("/", handlers.AddBook)
		r.With(auth.JWTMiddleware).Put("/{id}", handlers.UpdateBook)
		r.With(auth.JWTMiddleware).Delete("/{id}", handlers.DeleteBook)
	})

	// Unprotected routes
	r.Route("/authors", func(r chi.Router) {
		r.Get("/", handlers.GetAuthors)
		r.Get("/{id}", handlers.GetAuthor)
	})

	r.Get("/find/{genre}", handlers.FindByGenre)

	fmt.Println("Listening on :3000")
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatal(err)
	}
}
