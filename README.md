# book-management-system
This project is a book management system API. It's written in Go utilizing Gorilla Mux for routing and GORM library for interacting with a MySQL database.

### Requirements
The following are required to run the application:
- Golang v1.15 or higher
- MySQL v5.7 or higher

## Getting Started
To run the application:
1. Clone this repository
2. Run `go mod download` command to download all the dependencies.
3. Create a MySQL database.
4. Edit the connection string "dsn" in the Connect() function in the `config` package to match your database details.
5. Run `go run main.go` command to start the application.

## API Endpoints

| Endpoint | Method | Description |
| -------- | ------ | ----------- |
| /books   | POST   | Create a new book |
| /books   | GET    | Get all books |
| /books/{bookId} | GET    | Get a book by ID |
| /books/{bookId} | PUT    | Update a book by ID |
| /books/{bookId} | DELETE | Delete a book by ID |

## How to Use
Use any http client program to make HTTP requests to the running server.
### **Create a new Book**
#### **Request**
POST /books
```json
{
    "name": "The Alchemist",
    "author": "Paulo Coelho",
    "publication": "HarperOne"
}
```
#### **Response**
201 Created
```json
{
    "ID":1, // auto-increment
    "CreatedAt":"2022-01-23T09:17:23.212135216-05:00",
    "UpdatedAt":"2022-01-23T09:17:23.212135216-05:00",
    "DeletedAt":null,
    "name":"The Alchemist",
    "author":"Paulo Coelho",
    "publication":"HarperOne"
}
```
### **Get all Books**
#### **Request**
GET /books
#### **Response**
200 OK
```json
[
    {
    "ID":1,
    "CreatedAt":"2022-01-23T09:17:23.212+01:00",
    "UpdatedAt":"2022-01-23T09:17:23.212+01:00",
    "DeletedAt":null,
    "name":"The Alchemist",
    "author":"Paulo Coelho",
    "publication":"HarperOne"
    },
    {
    "ID":2,
    "CreatedAt":"2022-01-23T09:17:44.835+01:00",
    "UpdatedAt":"2022-01-23T09:17:44.835+01:00",
    "DeletedAt":null,
    "name":"The Da Vinci Code",
    "author":"Dan Brown",
    "publication":"Doubleday"
    }
]
```
### **Get a Book by ID**
#### **Request**
GET /books/1
#### **Response**
200 OK
```json
{
    "ID":1,
    "CreatedAt":"2022-01-23T09:17:23.212+01:00",
    "UpdatedAt":"2022-01-23T09:17:23.212+01:00",
    "DeletedAt":null,
    "name":"The Alchemist",
    "author":"Paulo Coelho",
    "publication":"HarperOne"
}
```
### **Update a Book by ID**
#### **Request**
PUT /books/1
```json
{
    "name": "New Name",
    "author": "New Author",
    "publication": "New Publication"
}
```
#### **Response**
200 OK
```json
{
    "ID":1,
    "CreatedAt":"2022-01-23T09:17:23.212+01:00",
    "UpdatedAt":"2023-04-11T10:18:37.281+01:00",
    "DeletedAt":null,
    "name": "New Name",
    "author": "New Author",
    "publication": "New Publication"
}
```
### **Delete a Book by ID**
#### **Request**
DELETE /books/2
#### **Response**
200 OK
```json
{
    "ID":0,
    "DeletedAt":"2023-05-16T06:54:14.762+01:00"
}
```
