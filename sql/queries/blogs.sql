-- name: CreateBlog :one
INSERT INTO blogs (title, content, tags)
VALUES (
    ?,
    ?,
    ?
)
RETURNING *;
