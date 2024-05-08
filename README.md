# Go CRUD with SQLite

This project implements a CRUD (Create, Read, Update, Delete) API in Go using SQLite as the database.


## Usage

To run the CRUD API, use the following command:

```
go run cmd/crud-api/main.go
```

### Endpoints

- `GET /books`: Retrieve all books.
- `POST /books/create`: Create a new book.
- `PUT /books/update`: Update an existing book.
- `DELETE /books/delete`: Delete an existing book.

### Using CLI (curl)

#### Get all books:

```bash
curl http://localhost:8080/books
```

#### Create a book:

```bash
curl -X POST -H "Content-Type: application/json" -d '{"title": "Book 2", "author": "Author 2"}' http://localhost:8080/books/create
```

#### Update a book:

```bash
curl -X PUT -H "Content-Type: application/json" -d '{"id": 1, "title": "Updated Book 1", "author": "Updated Author 1"}' http://localhost:8080/books/update
```

#### Delete a book:

```bash
curl -X DELETE -H "Content-Type: application/json" -d '{"id": 2}' http://localhost:8080/books/delete
```

### Using Postman

#### Get all books:

- Method: GET
- URL: http://localhost:8080/books

#### Create a book:

- Method: POST
- URL: http://localhost:8080/books/create
- Headers:
    - Key: Content-Type
    - Value: application/json
- Body (raw JSON):
  ```json
  {
      "title": "Book 2",
      "author": "Author 2"
  }
  ```

#### Update a book:

- Method: PUT
- URL: http://localhost:8080/books/update
- Headers:
    - Key: Content-Type
    - Value: application/json
- Body (raw JSON):
  ```json
  {
      "id": 1,
      "title": "Updated Book 1",
      "author": "Updated Author 1"
  }
  ```

#### Delete a book:

- Method: DELETE
- URL: http://localhost:8080/books/delete
- Headers:
    - Key: Content-Type
    - Value: application/json
- Body (raw JSON):
  ```json
  {
      "id": 2
  }
  ```


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

by **Eduardo Raider** - Software Engineer