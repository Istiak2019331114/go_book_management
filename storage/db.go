package storage

import "go_book_management/models"

var Books = make(map[string]models.Book)
var Authors = make(map[string]models.Author)

func Init() {
	Books["1"] = models.Book{ID: "1", Title: "Book1", Genre: "Comedy", Author: &models.Author{ID: "1", FirstName: "Istaikur", LastName: "Rahman"}}
	Books["2"] = models.Book{ID: "2", Title: "Book2", Genre: "Sci Fi", Author: &models.Author{ID: "2", FirstName: "Pulok", LastName: "Saha"}}

	Authors["1"] = models.Author{ID: "1", FirstName: "Istaikur", LastName: "Rahman"}
	Authors["2"] = models.Author{ID: "2", FirstName: "Pulok", LastName: "Saha"}
}
