-- name: ReportPost :one
INSERT INTO reports (id, reported_at, reported_by, reason)
VALUES (
   $1,
   NOW(),
   $2,
   $3
) RETURNING *;

-- name: GetAllReports :many
SELECT * FROM reports;