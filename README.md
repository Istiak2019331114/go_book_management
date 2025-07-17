## 📘 API Endpoints

| Method | Endpoint          | Description               |
|--------|-------------------|---------------------------|
| GET    | `/books`          | Retrieve all books        |
| GET    | `/books/{id}`     | Retrieve a book by ID     |
| POST   | `/books`          | Add a new book            |
| PUT    | `/books/{id}`     | Update an existing book   |
| DELETE | `/books/{id}`     | Delete a book             |
| GET    | `/find/{genre}`   | Search books by genre     |
| GET    | `/authors`        | Retrieve all authors      |
| GET    | `/authors/{id}`   | Retrieve an author by ID  |

# Go Book Management

## Run the App Directly

```bash
go run main.go
```

## Build the Go Binary

```bash
go build -o go-book .
```

## Run the Built Binary

```bash
./go-book
```

