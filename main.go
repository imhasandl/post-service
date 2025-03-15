package main

import (
	"database/sql"
	"log"
	"net"
	"os"
	"path/filepath"

	server "github.com/imhasandl/post-service/cmd/server"
	"github.com/imhasandl/post-service/internal/database"
	"github.com/imhasandl/post-service/internal/rabbitmq"
	pb "github.com/imhasandl/post-service/protos"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(filepath.Join("./", ".env")); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalf("Set Port in env")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatalf("Set db connection in env")
	}

	tokenSecret := os.Getenv("TOKEN_SECRET")
	if tokenSecret == "" {
		log.Fatalf("Set db connection in env")
	}

	rabbitmqURL := os.Getenv("RABBITMQ_URL")
	if rabbitmqURL == "" {
		log.Fatalf("Set rabbit mq url path")
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listed: %v", err)
	}

	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error opening database: %s", err)
	}
	defer dbConn.Close()
	dbQueries := database.New(dbConn)

	rebbitmq, err := rabbitmq.NewRabbitMQ(rabbitmqURL)
	if err != nil {
		log.Fatalf("Error connecting to rabbit mq: %v", err)
	}
	defer rebbitmq.Close()

	postServer := server.NewServer(dbQueries, tokenSecret, rebbitmq)

	s := grpc.NewServer()
	pb.RegisterPostServiceServer(s, postServer)

	reflection.Register(s)
	log.Printf("Server listening on %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to lister: %v", err)
	}
}
