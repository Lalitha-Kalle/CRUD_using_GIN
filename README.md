# 📚 Bookstore CRUD API with Gin (Go)

A simple in-memory CRUD (Create, Read, Update, Delete) API built using the Gin web framework in Go. This project is ideal for beginners who want to learn how to build RESTful APIs without involving a database.

---

## 🚀 Features

- ✅ Get all books
- ✅ Get a book by ID
- ✅ Create a new book
- ✅ Update a book by ID
- ✅ Delete a book by ID

---

## 🛠 Technologies Used

- [Go (Golang)](https://golang.org/)
- [Gin Web Framework](https://github.com/gin-gonic/gin)

---

## 📁 File Structure

```
.
├── main.go         # Main Go application file
└── README.md       # Project documentation
```

---

## ▶️ Getting Started

### 1. Install Go

Make sure Go is installed on your system. You can download it from [golang.org](https://golang.org/dl/).

### 2. Install Gin

Run the following command to install Gin:

```bash
go get -u github.com/gin-gonic/gin
```

### 3. Run the application

Clone the repo or create a file named main.go and paste the code. Then run:

```bash
go run main.go
```

Server will start on:

```
http://localhost:2000
```

---

## 📬 API Endpoints

### GET /books

Returns all books.

```bash
curl http://localhost:2000/books
```

---

### GET /books/:id

Returns a single book by ID.

```bash
curl http://localhost:2000/books/1
```

---

### POST /books

Creates a new book.

Example:

```bash
curl -X POST http://localhost:2000/books \
  -H "Content-Type: application/json" \
  -d '{"id": "3", "title": "Clean Code", "author": "Robert C. Martin"}'
```

---

### PUT /books/:id

Updates an existing book by ID.

Example:

```bash
curl -X PUT http://localhost:2000/books/1 \
  -H "Content-Type: application/json" \
  -d '{"id": "1", "title": "Deep Work", "author": "Cal Newport"}'
```

### DELETE /books/:id

Deletes an existing book by ID.

Example:

```bash
curl -X DELETE http://localhost:2000/books/1
```


---
