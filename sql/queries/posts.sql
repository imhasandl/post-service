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

-- name: ChangePost :one
UPDATE posts SET body = $1
WHERE id = $2
RETURNING *;

-- name: DeletePost :one
DELETE FROM posts WHERE id = $1 RETURNING *;

-- name: LikePost :one
UPDATE posts
SET likes = likes + 1,
    liked_by = array_append(liked_by, $2),
    updated_at = NOW()
WHERE id = $1
  AND NOT $2 = ANY(liked_by)
RETURNING *;

-- name: UnlikePost :one
UPDATE posts
SET likes = likes - 1,
    liked_by = array_remove(liked_by, $2),
    updated_at = NOW()
WHERE id = $1
  AND $2 = ANY(liked_by)
RETURNING *;

-- name: GetLikers :many
SELECT unnest(liked_by) AS liker_id
FROM posts
WHERE id = $1;