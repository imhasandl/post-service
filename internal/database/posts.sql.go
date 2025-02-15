// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: posts.sql

package database

import (
	"context"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const changePost = `-- name: ChangePost :one
UPDATE posts SET body = $1
WHERE id = $2
RETURNING id, created_at, updated_at, posted_by, body, likes, views, liked_by
`

type ChangePostParams struct {
	Body string
	ID   uuid.UUID
}

func (q *Queries) ChangePost(ctx context.Context, arg ChangePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, changePost, arg.Body, arg.ID)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.PostedBy,
		&i.Body,
		&i.Likes,
		&i.Views,
		pq.Array(&i.LikedBy),
	)
	return i, err
}

const createPost = `-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, posted_by, body, views, likes)
VALUES (
   $1,
   NOW(),
   NOW(),
   $2,
   $3,
   0,
   0
) RETURNING id, created_at, updated_at, posted_by, body, likes, views, liked_by
`

type CreatePostParams struct {
	ID       uuid.UUID
	PostedBy string
	Body     string
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createPost, arg.ID, arg.PostedBy, arg.Body)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.PostedBy,
		&i.Body,
		&i.Likes,
		&i.Views,
		pq.Array(&i.LikedBy),
	)
	return i, err
}

const deletePost = `-- name: DeletePost :one
DELETE FROM posts WHERE id = $1 RETURNING id, created_at, updated_at, posted_by, body, likes, views, liked_by
`

func (q *Queries) DeletePost(ctx context.Context, id uuid.UUID) (Post, error) {
	row := q.db.QueryRowContext(ctx, deletePost, id)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.PostedBy,
		&i.Body,
		&i.Likes,
		&i.Views,
		pq.Array(&i.LikedBy),
	)
	return i, err
}

const getAllPosts = `-- name: GetAllPosts :many
SELECT id, created_at, updated_at, posted_by, body, likes, views, liked_by FROM posts
`

func (q *Queries) GetAllPosts(ctx context.Context) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, getAllPosts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.PostedBy,
			&i.Body,
			&i.Likes,
			&i.Views,
			pq.Array(&i.LikedBy),
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getLikers = `-- name: GetLikers :many
SELECT unnest(liked_by) AS liker_id
FROM posts
WHERE id = $1
`

func (q *Queries) GetLikers(ctx context.Context, id uuid.UUID) ([]interface{}, error) {
	rows, err := q.db.QueryContext(ctx, getLikers, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []interface{}
	for rows.Next() {
		var liker_id interface{}
		if err := rows.Scan(&liker_id); err != nil {
			return nil, err
		}
		items = append(items, liker_id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPostByID = `-- name: GetPostByID :one
SELECT id, created_at, updated_at, posted_by, body, likes, views, liked_by FROM posts
WHERE id = $1
`

func (q *Queries) GetPostByID(ctx context.Context, id uuid.UUID) (Post, error) {
	row := q.db.QueryRowContext(ctx, getPostByID, id)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.PostedBy,
		&i.Body,
		&i.Likes,
		&i.Views,
		pq.Array(&i.LikedBy),
	)
	return i, err
}

const likePost = `-- name: LikePost :one
UPDATE posts
SET likes = likes + 1,
    liked_by = array_append(liked_by, $2),
    updated_at = NOW()
WHERE id = $1
  AND NOT $2 = ANY(liked_by)
RETURNING id, created_at, updated_at, posted_by, body, likes, views, liked_by
`

type LikePostParams struct {
	ID          uuid.UUID
	ArrayAppend interface{}
}

func (q *Queries) LikePost(ctx context.Context, arg LikePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, likePost, arg.ID, arg.ArrayAppend)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.PostedBy,
		&i.Body,
		&i.Likes,
		&i.Views,
		pq.Array(&i.LikedBy),
	)
	return i, err
}

const unlikePost = `-- name: UnlikePost :one
UPDATE posts
SET likes = likes - 1,
    liked_by = array_remove(liked_by, $2),
    updated_at = NOW()
WHERE id = $1
  AND $2 = ANY(liked_by)
RETURNING id, created_at, updated_at, posted_by, body, likes, views, liked_by
`

type UnlikePostParams struct {
	ID          uuid.UUID
	ArrayRemove interface{}
}

func (q *Queries) UnlikePost(ctx context.Context, arg UnlikePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, unlikePost, arg.ID, arg.ArrayRemove)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.PostedBy,
		&i.Body,
		&i.Likes,
		&i.Views,
		pq.Array(&i.LikedBy),
	)
	return i, err
}
