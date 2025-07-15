package cmd

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go_book_management/auth"
	"go_book_management/handlers"
	"go_book_management/storage"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var port int

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		storage.Init()
		r := chi.NewRouter()
		r.Use(middleware.Logger)

		r.Post("/login", auth.Login)

		r.Route("/books", func(r chi.Router) {
			r.Get("/", handlers.GetBooks)
			r.Get("/{id}", handlers.GetBook)

			r.With(auth.JWTMiddleware).Post("/", handlers.AddBook)
			r.With(auth.JWTMiddleware).Put("/{id}", handlers.UpdateBook)
			r.With(auth.JWTMiddleware).Delete("/{id}", handlers.DeleteBook)
		})

		r.Route("/authors", func(r chi.Router) {
			r.Get("/", handlers.GetAuthors)
			r.Get("/{id}", handlers.GetAuthor)
		})

		r.Get("/find/{genre}", handlers.FindByGenre)

		addr := fmt.Sprintf(":%d", port)
		fmt.Println("Listening on", addr)
		log.Fatal(http.ListenAndServe(addr, r))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().IntVarP(&port, "port", "p", 4000, "Port to run the server on")
}
