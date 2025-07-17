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

# Go Book Management – Docker Commands

## Build the Docker Image

```bash
docker build -t go-book-app .
```

## Run the Container

### Run and expose port 8080
```bash
docker run -p 8080:8080 go-book-app --port 8080
```

### Run and expose port 3000
```bash
docker run -p 3000:3000 go-book-app --port 3000
```

## List Running Containers

```bash
docker ps
```

## View Logs

```bash
docker logs <container_id>
```

## Stop a Container

```bash
docker stop <container_id>
```

## Rebuild After Code Changes

```bash
docker build -t go-book-app .
```

## Debug Inside Container

```bash
docker exec -it <container_id> sh
```

## Clean Up Stopped Containers

```bash
docker container prune
```
