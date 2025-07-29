# Go Book Management

A simple book management API built with Go, Chi router, JWT authentication, and Cobra CLI. This app supports CRUD operations for books and authors, along with authentication.

---

## 📘 API Endpoints

| Method | Endpoint        | Description              | Auth Required?         |
|--------|------------------|--------------------------|-------------------------|
| POST   | `/login`         | Login and get JWT cookie | No                      |
| GET    | `/books`         | Retrieve all books       | No                      |
| GET    | `/books/{id}`    | Retrieve a book by ID    | No                      |
| POST   | `/books`         | Add a new book           | Yes (JWT token cookie)  |
| PUT    | `/books/{id}`    | Update an existing book  | Yes (JWT token cookie)  |
| DELETE | `/books/{id}`    | Delete a book            | Yes (JWT token cookie)  |
| GET    | `/find/{genre}`  | Search books by genre    | No                      |
| GET    | `/authors`       | Retrieve all authors     | No                      |
| GET    | `/authors/{id}`  | Retrieve an author by ID | No                      |

---

## Docker Usage

### Build Docker Image

```bash
docker build -t go-book-app .
```

### Run the Container (on port 8080)

```bash
docker run -p 8080:8080 go-book-app --port 8080
```

---

## Kubernetes Deployment (Using Kind)

### Step 1: Push Image to Docker Hub (or use `kind load` for local dev)

If using Docker Hub:

```bash
docker tag go-book-app istiaka2i/bookapp:v1.0
docker push istiaka2i/bookapp:v1.0
```

If using Kind locally:

```bash
kind load docker-image go-book-app
```

### Step 2: Apply Kubernetes Manifest

```bash
kubectl apply -f go-book-app.yaml
```

### Step 3: Forward Port to Access App Locally

```bash
kubectl port-forward svc/go-book-service 4000:4000
```

Then access the app at:

```
http://localhost:4000
```

---

## 🔌 Port Mapping Explained

| Level           | Port | Description                                                |
|-----------------|------|------------------------------------------------------------|
| Container Port  | 4000 | The Go app runs on this port inside the container          |
| Target Port     | 4000 | Kubernetes service forwards traffic to this container port |
| Node Port       | 30080| Exposes service on the node (not usable directly in `kind`)|
| Host Port       | 4000 | Exposed using `kubectl port-forward` for local access      |

**Note:** Since `kind` runs inside Docker, `NodePort` is not directly reachable. Use `kubectl port-forward` for access.

---

## Useful Docker Commands

| Task                         | Command                               |
|-----------------------------|---------------------------------------|
| View running containers     | `docker ps`                           |
| View container logs         | `docker logs <container_id>`          |
| Stop a container            | `docker stop <container_id>`          |
| Debug into a container      | `docker exec -it <container_id> sh`   |
| Clean up stopped containers | `docker container prune`              |

---
