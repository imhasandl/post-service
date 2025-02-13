package main

import (
	"context"

	"github.com/imhasandl/post-service/internal/database"
	pb "github.com/imhasandl/post-service/internal/protos"
)
	
type server struct {
	pb.UnimplementedPostServiceServer
	db *database.Queries
}

func NewServer(db *database.Queries) *server {
	return &server{
		pb.UnimplementedPostServiceServer{},
		db,
	}
}

func (*server) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.CreatePostResponse, error) {
	return nil, nil
}

func (*server) GetPostByUsername(ctx context.Context, req *pb.GetPostByUsernameRequest) (*pb.GetPostByUsernameResponse, error) {
	return nil, nil
}

func (*server) ReportPost(ctx context.Context, req *pb.ReportPostRequest) (*pb.ReportPostResponse, error) {
	return nil, nil
}