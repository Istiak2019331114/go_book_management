## 📘 API Endpoints

| Method | Endpoint        | Description              | Auth Required?         |
| ------ | --------------- | ------------------------ | ---------------------- |
| POST   | `/login`        | Login and get JWT cookie | No                     |
| GET    | `/books`        | Retrieve all books       | No                     |
| GET    | `/books/{id}`   | Retrieve a book by ID    | No                     |
| POST   | `/books`        | Add a new book           | Yes (JWT token cookie) |
| PUT    | `/books/{id}`   | Update an existing book  | Yes (JWT token cookie) |
| DELETE | `/books/{id}`   | Delete a book            | Yes (JWT token cookie) |
| GET    | `/find/{genre}` | Search books by genre    | No                     |
| GET    | `/authors`      | Retrieve all authors     | No                     |
| GET    | `/authors/{id}` | Retrieve an author by ID | No                     |


# Go Book Management – Cobra CLI Commands

## Run the App with Default Port (4000)

```bash
go run main.go serve
```

## Run the App with Custom Port

```bash
go run main.go serve --port 8080
```

## Build the Go Binary

```bash
go build -o go-book .
```

## Run the Built Binary with Default Port

```bash
./go-book serve
```

## Run the Built Binary with Custom Port

```bash
./go-book serve --port 3000
```

## Cobra Help Command

```bash
go run main.go --help
```

## Cobra Subcommand Help

```bash
go run main.go serve --help
```

