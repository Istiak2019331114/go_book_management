package models

import _ "fmt"

type Author struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Genre  string  `json:"genre"`
	Author *Author `json:"author"`
}
