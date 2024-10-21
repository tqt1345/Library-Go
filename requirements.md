# Requirements

## Core Features

User registration and login
View book catalogue
Search for books by title or author
Select a book to loan
Loans last for 1 week
can return books early (cancels the loan early)
three max total loans

## TODO

### Book catalogue

- [ ] Display a table for each book entity
  - [x] book title
  - [ ] book authors
  - [x] book description
  - [x] link to book page

- [ ] Implement book and author many to many relationship
- [ ] Book details page shows the books authors

### Domain objects

Users

- user id (primary key)
- first name
- last name
- email
- username
- password
- role
- loan

Loans

- loan id (primary key)
- loan start date
- loan return date
- loan due date
- book

Books

- book id (primary key)
- title
- description

Author

- author id
- first name
- last name

User -> Loans (Many to one)
Loans -> Books (One to one)
Books -> Author (Many to many)

## Tech stack

- Go
- Sqlite3
- htmx
- templ
