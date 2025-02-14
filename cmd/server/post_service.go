package server

import (
	"context"

	"github.com/imhasandl/post-service/internal/database"
	pb "github.com/imhasandl/post-service/internal/protos"
)
	
type server struct {
	pb.UnimplementedPostServiceServer
	db *database.Queries
	tokenSecret string
}

func NewServer(db *database.Queries, tokenSecret string) *server {
	return &server{
		pb.UnimplementedPostServiceServer{},
		db,
		tokenSecret,
	}
}

func (s *server) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.CreatePostResponse, error) {
	// accessToken, err := helper.GetBearerTokenFromGrpc(ctx)
	// if err != nil {
	// 	return nil, status.Errorf(codes.Internal, "can't get bearer token: %v - CreatePost", err)
	// }

	// userID, err := helper.ValidateJWT(accessToken, s.tokenSecret)
	// if err != nil {
	// 	return nil, status.Errorf(codes.Internal, "can't get user id from token: %v - CreatePost", err)
	// }

	// postParams := database.CreatePostParams{
	// 	ID: uuid.New(),
	// 	PostedBy: ,
	// 	Body: req.GetBody(),
	// }
	return nil, nil
}

func (s *server) ReportPost(ctx context.Context, req *pb.ReportPostRequest) (*pb.ReportPostResponse, error) {
	return nil, nil
}
