package server

import (
	"context"

	"github.com/google/uuid"
	"github.com/imhasandl/post-service/cmd/helper"
	"github.com/imhasandl/post-service/internal/database"
	pb "github.com/imhasandl/post-service/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type server struct {
	pb.UnimplementedPostServiceServer
	db          *database.Queries
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
	accessToken, err := helper.GetBearerTokenFromGrpc(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't get bearer token: %v - CreatePost", err)
	}

	userID, err := helper.ValidateJWT(accessToken, s.tokenSecret)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't get user id from token: %v - CreatePost", err)
	}

	postParams := database.CreatePostParams{
		ID:       uuid.New(),
		PostedBy: userID.String(),
		Body:     req.GetBody(),
	}

	post, err := s.db.CreatePost(ctx, postParams)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't create post to database: %v - CreatePost", err)
	}

	createdAtProto := timestamppb.New(post.CreatedAt)
	updatedAtProto := timestamppb.New(post.UpdatedAt)

	return &pb.CreatePostResponse{
		Post: &pb.Post{
			Id:        post.ID.String(),
			CreatedAt: createdAtProto,
			UpdatedAt: updatedAtProto,
			PostedBy:  post.PostedBy,
			Body:      post.Body,
			Likes:     post.Likes,
			Views:     post.Views,
		},
	}, nil
}

func (s *server) GetAllPosts(ctx context.Context, req *pb.GetAllPostsRequest) (*pb.GetAllPostsResponse, error) {
	posts, err := s.db.GetAllPosts(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't get all posts from db: %v - GetAllPosts", err)
	}

	pbPosts := make([]*pb.Post, len(posts))
	for i, post := range posts {
		createdAtProto := timestamppb.New(post.CreatedAt)
		updatedAtProto := timestamppb.New(post.UpdatedAt)

		pbPosts[i] = &pb.Post{
			Id:        post.ID.String(),
			CreatedAt: createdAtProto,
			UpdatedAt: updatedAtProto,
			PostedBy:  post.PostedBy,
			Body:      post.Body,
			Likes:     post.Likes,
			Views:     post.Views,
		}
	}

	return &pb.GetAllPostsResponse{
		Posts: pbPosts,
	}, nil
}

func (s *server) GetPostByID(ctx context.Context, req *pb.GetPostByIDRequest) (*pb.GetPostByIDResponse, error) {
	postID, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't parse id from request: %v - GetPostByID", err)
	}

	post, err := s.db.GetPostByID(ctx, postID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't get post from db: %v - GetPostByID", err)
	}

	createdAtProto := timestamppb.New(post.CreatedAt)
	updatedAtProto := timestamppb.New(post.UpdatedAt)

	return &pb.GetPostByIDResponse{
		Post: &pb.Post{
			Id:        post.ID.String(),
			CreatedAt: createdAtProto,
			UpdatedAt: updatedAtProto,
			PostedBy:  post.PostedBy,
			Body:      post.Body,
			Views:     post.Views,
			Likes:     post.Likes,
		},
	}, nil
}

func (s *server) ReportPost(ctx context.Context, req *pb.ReportPostRequest) (*pb.ReportPostResponse, error) {
	postID, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't report post: %v - ReportPost", err)
	}

	reportParams := database.ReportPostParams{
		ID:         postID,
		ReportedBy: req.GetReportedBy(),
		Reason:     req.GetReason(),
	}

	report, err := s.db.ReportPost(ctx, reportParams)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't report post to db: %v - ReportPost", err)
	}

	reportedAt := timestamppb.New(report.ReportedAt)

	return &pb.ReportPostResponse{
		ReportPost: &pb.ReportPost{
			Id:         report.ID.String(),
			ReportedAt: reportedAt,
			ReportedBy: report.ReportedBy,
			Reason:     report.Reason,
		},
	}, nil
}
