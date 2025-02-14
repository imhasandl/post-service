module github.com/imhasandl/post-service

go 1.23.5

require google.golang.org/grpc v1.70.0

require github.com/golang-jwt/jwt/v5 v5.2.1

require (
	github.com/imhasandl/post-service/internal/protos v0.0.0-20250214064643-e32bb57a3dad
	github.com/joho/godotenv v1.5.1
	github.com/lib/pq v1.10.9 // indirect
	golang.org/x/net v0.32.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241202173237-19429a94021a // indirect
	google.golang.org/protobuf v1.36.5 // indirect
)

require (
	github.com/google/uuid v1.6.0
	golang.org/x/sys v0.30.0 // indirect
)
