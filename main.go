package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/imhasandl/post-service/cmd/helper"
	server "github.com/imhasandl/post-service/cmd/server"
	"github.com/imhasandl/post-service/internal/database"
	"github.com/imhasandl/post-service/internal/rabbitmq"
	"github.com/imhasandl/post-service/internal/redis"
	pb "github.com/imhasandl/post-service/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/lib/pq"
)

func main() {
	envConfig := helper.GetENVSecrets()

	lis, err := net.Listen("tcp", envConfig.Port)
	if err != nil {
		log.Fatalf("failed to listed: %v", err)
	}

	dbConn, err := sql.Open("postgres", envConfig.DBURL)
	if err != nil {
		log.Fatalf("Error opening database: %s", err)
	}
	defer dbConn.Close()
	dbQueries := database.New(dbConn)

	rebbitmq, err := rabbitmq.NewRabbitMQ(envConfig.Rabbitmq)
	if err != nil {
		log.Fatalf("Error connecting to rabbit mq: %v", err)
	}
	defer rebbitmq.Close()

	redis, err := redis.NewRedisClient(envConfig.RedisSecret)
	if err != nil {
		log.Fatalf("Error connecting to Redis: %v", err)
	}
	defer redis.Close()

	postServer := server.NewServer(dbQueries, envConfig.TokenSecret, rebbitmq, redis)

	s := grpc.NewServer()
	pb.RegisterPostServiceServer(s, postServer)

	reflection.Register(s)
	log.Printf("Server listening on %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to lister: %v", err)
	}
}
