-- name: CreateComment :one
INSERT INTO comments (id, created_at, post_id, user_id, comment_text)
VALUES (
   $1,
   NOW(),
   $2,
   $3,
   $4
) RETURNING *;

-- name: DeleteComment :exec
DELETE FROM comments
WHERE id = $1;

-- name: GetCommentByID :one
SELECT id, user_id FROM comments
WHERE id = $1;