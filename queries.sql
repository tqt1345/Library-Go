-- name: FindAllBooks :many 
SELECT * FROM books
ORDER BY title;

-- name: FindBookById :one
SELECT * FROM books WHERE id = ?;

-- name: FindBookByTitle :many
SELECT * FROM books WHERE title LIKE '%' || ? || '%';
