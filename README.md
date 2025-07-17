# Go Book Management API with JWT Authentication

This is a REST API for managing books and authors built with Go.  
It now includes JWT-based authentication to protect modification routes.

---

## Features

- User login with JWT token issued in an HTTP-only cookie  
- Public endpoints for retrieving books and authors  
- Protected endpoints for adding, updating, and deleting books (require JWT auth)  
- Middleware to verify JWT token on protected routes  

---

## Setup & Run

1. Clone the repo:

```bash
git clone https://github.com/yourusername/go-book-management.git
cd go-book-management
````

2. Install dependencies:

```bash
go mod tidy
```

3. Run the server:

```bash
go run main.go
```

Server listens on port `3000`.

---

## Authentication

* POST `/login` with JSON body:

  ```json
  {
    "username": "admin",
    "password": "whoami"
  }
  ```

* On success, a JWT token cookie named `token` is set (HttpOnly).

* Use this cookie in subsequent requests to protected endpoints.

* Token expires in 1 hour.

---

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

---
