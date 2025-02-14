-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, posted_by, body, views, likes)
VALUES (
   $1,
   NOW(),
   NOW(),
   $2,
   $3,
   0,
   0
) RETURNING *;

-- name: GetPostByID :one
SELECT * FROM posts
WHERE id = $1;

-- name: GetAllPosts :many
SELECT * FROM posts;