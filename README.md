````ProductsGo API

A simple RESTful API for managing products, built with [Go Fiber](https://gofiber.io/).

## Features

- Health check endpoint
- Create new products
- List all products
- Middleware example for `/api` routes

## Prerequisites

- [Go](https://go.dev/) 1.18 or newer
- MongoDB (if your handlers use it for storage)

## Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/yourusername/productsGo.git
cd productsGo
````

### 2. Install dependencies

```bash
go mod tidy
```

### 3. Run the API

```bash
go run main.go
```

The API will be available at [http://localhost:9000](http://localhost:9000).

## API Endpoints

| Method | Endpoint        | Description           |
| ------ | --------------- | --------------------- |
| GET    | `/healthcheck`  | Health check endpoint |
| POST   | `/api/products` | Create a new product  |
| GET    | `/api/products` | List all products     |

## Example Requests

### Health Check

```bash
curl http://localhost:9000/healthcheck
```

### Create a Product

```bash
curl -X POST http://localhost:9000/api/products \
  -H "Content-Type: application/json" \
  -d '{"name":"Sample Product","description":"A test product","price":19.99}'
```

### List All Products

```bash
curl http://localhost:9000/api/products
```

## Middleware

All `/api` routes use a middleware that prints a message to the console for every request.

## Project Structure

```
.
├── main.go
├── internal/
│   └── handlers/
│       └── product.go
└── README.md
```

## License

MIT
