-- name: ListBooks :many 
SELECT * FROM books
ORDER BY title;

-- name: BookId1 :one
SELECT * FROM books WHERE id = 1;
