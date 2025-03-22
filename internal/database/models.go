// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package database

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	PostID      uuid.UUID
	UserID      uuid.UUID
	CommentText string
}

type DeviceToken struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	DeviceToken string
	DeviceType  string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Message struct {
	ID         uuid.UUID
	SentAt     time.Time
	SenderID   uuid.UUID
	ReceiverID uuid.UUID
	Content    string
}

type Post struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	PostedBy  uuid.UUID
	Body      string
	Likes     int32
	Views     int32
	LikedBy   []string
}

type RefreshToken struct {
	Token      string
	UserID     uuid.UUID
	ExpiryTime time.Time
	CreatedAt  time.Time
}

type Report struct {
	ID         uuid.UUID
	ReportedAt time.Time
	ReportedBy uuid.UUID
	Reason     string
}

type User struct {
	ID               uuid.UUID
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Email            string
	Password         string
	Username         string
	Subscribers      []uuid.UUID
	SubscribedTo     []uuid.UUID
	IsPremium        bool
	VerificationCode int32
	IsVerified       bool
}
