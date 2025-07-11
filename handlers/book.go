package handlers

import (
	_ "fmt"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"go_book_management/models"
	"go_book_management/storage"
	"net/http"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book
	for _, book := range storage.Books {
		books = append(books, book)
	}
	err := json.NewEncoder(w).Encode(books)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	book, ok := storage.Books[id]
	if !ok {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	err := json.NewEncoder(w).Encode(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, ok := storage.Books[book.ID]
	if ok {
		http.Error(w, "Book already exists", http.StatusConflict)
		return
	}
	storage.Books[book.ID] = book
	storage.Authors[book.Author.ID] = *book.Author
	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	oldId := chi.URLParam(r, "id")
	if len(oldId) == 0 {
		http.Error(w, "No ID provided", http.StatusBadRequest)
		return
	}
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, ok := storage.Books[oldId]
	if !ok {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	if book.ID != oldId {
		http.Error(w, "You cannot change ID of a book.", http.StatusBadRequest)
		return
	}
	storage.Books[oldId] = book
	storage.Authors[book.Author.ID] = *book.Author
	_, err = w.Write([]byte("Book updated successfully"))
	if err != nil {
		http.Error(w, "Cannot write response", http.StatusInternalServerError)
	}
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	oldId := chi.URLParam(r, "id")
	if len(oldId) == 0 {
		http.Error(w, "No ID provided", http.StatusBadRequest)
		return
	}
	_, ok := storage.Books[oldId]
	if !ok {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	delete(storage.Authors, storage.Books[oldId].Author.ID)
	delete(storage.Books, oldId)
	_, err := w.Write([]byte("Data deleted successfully"))
	if err != nil {
		http.Error(w, "Cannot write response", http.StatusInternalServerError)
	}
}

func FindByGenre(w http.ResponseWriter, r *http.Request) {
	gen := chi.URLParam(r, "genre")
	if len(gen) == 0 {
		http.Error(w, "No genre provided", http.StatusBadRequest)
		return
	}
	for _, book := range storage.Books {
		if book.Genre == gen {
			err := json.NewEncoder(w).Encode(book)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
}
