package handlers

import (
	_ "fmt"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"go_book_management/models"
	"go_book_management/storage"
	"net/http"
)

func GetAuthors(w http.ResponseWriter, r *http.Request) {
	var authors []models.Author
	for _, author := range storage.Authors {
		authors = append(authors, author)
	}
	err := json.NewEncoder(w).Encode(authors)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetAuthor(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	author, ok := storage.Authors[id]
	if !ok {
		http.Error(w, "Author not found", http.StatusNotFound)
		return
	}
	err := json.NewEncoder(w).Encode(author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
