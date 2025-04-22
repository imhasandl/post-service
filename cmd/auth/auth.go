package auth

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/metadata"
)

// GetBearerTokenFromGrpc extracts the Bearer token from the gRPC context metadata.
// It returns the token string or an error if the token is not found or is invalid.
func GetBearerTokenFromGrpc(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", fmt.Errorf("metadata is not found in context")
	}

	authHeader := md.Get("authorization")
	if len(authHeader) == 0 {
		return "", fmt.Errorf("authorization header is not found")
	}

	bearerToken := authHeader[0]
	if !strings.HasPrefix(bearerToken, "Bearer ") {
		return "", fmt.Errorf("invalid authorization header format")
	}

	return strings.TrimPrefix(bearerToken, "Bearer "), nil
}
