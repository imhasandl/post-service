[![CI](https://github.com/imhasandl/post-service/actions/workflows/ci.yml/badge.svg)](https://github.com/imhasandl/post-service/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/imhasandl/post-service)](https://goreportcard.com/report/github.com/imhasandl/post-service)
[![GoDoc](https://godoc.org/github.com/imhasandl/post-service?status.svg)](https://godoc.org/github.com/imhasandl/post-service)
[![Coverage](https://codecov.io/gh/imhasandl/post-service/branch/main/graph/badge.svg)](https://codecov.io/gh/imhasandl/post-service)
[![Go Version](https://img.shields.io/github/go-mod/go-version/imhasandl/post-service)](https://golang.org/doc/devel/release.html)nication.

# Post Service

A microservice for post management in a social media application, built with Go and gRPC.

## Overview posts by user ID

* Update existing posts
The Post Service is responsible for managing posts, comments, and interactions in the social media platform. It provides functionality for creating, retrieving, updating, and deleting posts, as well as liking/unliking posts, adding comments, and reporting inappropriate content. The service uses gRPC for communication with other services in the microservices architecture.

## Prerequisitesror handling and logging.

- Go 1.23 or latered
- PostgreSQL database
- RabbitMQ (for event-driven communication with other services)

## Configurationfers

Create a `.env` file in the root directory with the following variables:

```envtory.
PORT=":50051"2.  Install dependencies using `go mod tidy`.
DB_URL="postgres://username:password@host:port/database?sslmode=disable"Run the service using `go run main.go`.
# DB_URL="postgres://username:password@db:port/database?sslmode=disable" // FOR DOCKER COMPOSE
TOKEN_SECRET="YOUR_JWT_SECRET_KEY"
RABBITMQ_URL="amqp://username:password@host:port"
```umentation is not ready yet.
## gRPC MethodsThe service implements the following gRPC methods:### CreatePostCreates a new post.#### Request Format```json{   "body": "This is the content of my post"}```#### Response Format```json{   "post": {      "id": "UUID of the created post",      "created_at": "2023-01-01T12:00:00Z",      "updated_at": "2023-01-01T12:00:00Z",      "posted_by": "UUID of the user who created the post",      "body": "This is the content of my post"
   }
}
```

### ChangePost

Updates an existing post.

#### Request Format

```json
{
   "id": "UUID of the post",
   "body": "This is the updated content of my post"
}
```

#### Response Format

```json
{
   "post": {
      "id": "UUID of the post",
      "created_at": "2023-01-01T12:00:00Z",
      "updated_at": "2023-01-01T12:30:00Z",
      "posted_by": "UUID of the user who created the post",
      "body": "This is the updated content of my post"
   }
}
```

### DeletePost

Deletes a post.

#### Request Format

```json
{
   "id": "UUID of the post"
}
```

#### Response Format

```json
{
   "result": "Post successfully deleted"
}
```

### LikePost

Likes a post.

#### Request Format

```json
{
   "post_id": "UUID of the post",
   "liked_by": "UUID of the user liking the post"
}
```

#### Response Format

```json
{
   "post": {
      "id": "UUID of the post",
      "created_at": "2023-01-01T12:00:00Z",
      "updated_at": "2023-01-01T12:00:00Z",
      "posted_by": "UUID of the user who created the post",
      "body": "Content of the post"
   }
}
```

### UnlikePost

Removes a like from a post.

#### Request Format

```json
{
   "post_id": "UUID of the post",
   "unliked_by": "UUID of the user unliking the post"
}
```

#### Response Format

```json
{
   "post": {
      "id": "UUID of the post",
      "created_at": "2023-01-01T12:00:00Z",
      "updated_at": "2023-01-01T12:00:00Z",
      "posted_by": "UUID of the user who created the post",
      "body": "Content of the post"
   }
}
```

### GetLikersFromPost

Retrieves the list of users who liked a post.

#### Request Format

```json
{
   "post_id": "UUID of the post"
}
```

#### Response Format

```json
{
   "liked_by": ["UUID1", "UUID2", "UUID3"]
}
```

### GetPostByID

Retrieves a single post by its ID.

#### Request Format

```json
{
   "id": "UUID of the post"
}
```

#### Response Format

```json
{
   "post": {
      "id": "UUID of the post",
      "created_at": "2023-01-01T12:00:00Z",
      "updated_at": "2023-01-01T12:00:00Z",
      "posted_by": "UUID of the user who created the post",
      "body": "Content of the post"
   }
}
```

### GetAllPosts

Retrieves all posts.

#### Request Format

```json
{}
```

#### Response Format

```json
{
   "posts": [
      {
         "id": "UUID of post 1",
         "created_at": "2023-01-01T12:00:00Z",
         "updated_at": "2023-01-01T12:00:00Z",
         "posted_by": "UUID of the user who created the post",
         "body": "Content of post 1"
      },
      {
         "id": "UUID of post 2",
         "created_at": "2023-01-02T10:00:00Z",
         "updated_at": "2023-01-02T10:00:00Z",
         "posted_by": "UUID of the user who created the post",
         "body": "Content of post 2"
      }
   ]
}
```

### ReportPost

Reports a post for inappropriate content.

#### Request Format

```json
{
   "id": "UUID of the post",
   "reason": "Reason for reporting the post"
}
```

#### Response Format

```json
{
   "report_post": {
      "reported_at": "2023-01-01T12:00:00Z",
      "post_id": "UUID of the reported post",
      "reason": "Reason for reporting the post"
   }
}
```

### GetAllReports

Retrieves all reported posts.

#### Request Format

```json
{}
```

#### Response Format

```json
{
   "report_post": [
      {
         "reported_at": "2023-01-01T12:00:00Z",
         "post_id": "UUID of the reported post",
         "reason": "Reason for reporting the post"
      },
      {
         "reported_at": "2023-01-02T10:00:00Z",
         "post_id": "UUID of another reported post",
         "reason": "Reason for reporting the post"
      }
   ]
}
```

### CreateComment

Creates a new comment on a post.

#### Request Format

```json
{
   "post_id": "UUID of the post",
   "comment_text": "This is a comment on the post"
}
```

#### Response Format

```json
{
   "comment": {
      "id": "UUID of the comment",
      "created_at": "2023-01-01T12:00:00Z",
      "post_id": "UUID of the post",
      "commented_by": "UUID of the user who created the comment",
      "comment_text": "This is a comment on the post"
   }
}
```

### DeleteComment

Deletes a comment.

#### Request Format

```json
{
   "id": "UUID of the comment"
}
```

#### Response Format

```json
{
   "status": "Comment successfully deleted"
}
```

### ResetPosts

Resets all posts (for development purposes only).

#### Request Format

```json
{}
```

#### Response Format

```json
{
   "status": "All posts successfully reset"
}
```

## RabbitMQ Integration

The Post Service publishes events to RabbitMQ when significant post actions occur, enabling other services to react accordingly.

### Event Publication

The service publishes events to:
- **Exchange**: `posts.topic` (topic exchange)
- **Routing Keys**:
  - `post.created` - When a new post is created
  - `post.updated` - When a post is updated
  - `post.deleted` - When a post is deleted
  - `post.liked` - When a post is liked
  - `post.comment.added` - When a comment is added to a post
  - `post.reported` - When a post is reported

### Message Format Example

```json
{
   "event_type": "post.created",
   "post_id": "UUID of the post",
   "user_id": "UUID of the user who created the post",
   "timestamp": "2023-01-01T12:00:00Z",
   "data": {
      "body": "Content of the post"
   }
}
```

## Running the Service

```bash
go run main.go
```

## Docker Support

The service can be run using Docker:

```bash
# Build the Docker image
docker build -t post-service .

# Run the container
docker run -p 50051:50051 post-service
```

When deploying to different CPU architectures:

```bash
docker build --platform=linux/amd64 -t post-service .
```

For more details on Docker deployment, see [README.Docker.md](./README.Docker.md).
