package auth

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// TokenType represents the type of JWT token used in the authentication system.
type TokenType string

// Constants for different token types
const (
	// TokenTypeAccess -
	TokenTypeAccess TokenType = "media-access"
)

// ValidateJWT validates the provided JWT token and returns the claims if valid.
// It checks token signature and expiration using the provided secret key.
func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error) {
	claimsStruct := jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(
		tokenString,
		&claimsStruct,
		func(token *jwt.Token) (interface{}, error) { return []byte(tokenSecret), nil },
	)
	if err != nil {
		return uuid.Nil, errors.New("no auth header included in request")
	}

	userIDString, err := token.Claims.GetSubject()
	if err != nil {
		return uuid.Nil, errors.New("no auth header included in request")
	}

	issuer, err := token.Claims.GetIssuer()
	if err != nil {
		return uuid.Nil, errors.New("no auth header included in request")
	}
	if issuer != string(TokenTypeAccess) {
		return uuid.Nil, errors.New("invalid issuer")
	}

	id, err := uuid.Parse(userIDString)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid user ID: %w", err)
	}

	return id, nil
}
