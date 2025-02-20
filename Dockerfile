# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o out && ./out  # Build the executable as 'out'

# Final stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/out /app/out  
COPY .env .env

EXPOSE 50052

CMD ["./out"]  # Run the 'out' executable