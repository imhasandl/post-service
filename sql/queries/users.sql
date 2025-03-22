-- name: GetSubscribers :many
SELECT subscribers FROM users
WHERE id = $1;