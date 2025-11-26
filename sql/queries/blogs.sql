-- name: CreateBlog :one
INSERT INTO blogs (title, content, category, tags)
VALUES (
    ?,
    ?,
    ?,
    ?
)
RETURNING *;

-- name: UpdateBlog :one
UPDATE blogs
SET title = ?, content = ?, category = ?, tags = ?, updatedAt = current_timestamp
WHERE id = ?
RETURNING *;

-- name: DeleteBlog :one
DELETE FROM blogs
WHERE id = ?
RETURNING title;

-- name: GetBlog :one
SELECT * FROM blogs
WHERE id = ?;
