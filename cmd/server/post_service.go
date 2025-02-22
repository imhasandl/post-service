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

	return &pb.CreatePostResponse{
		Post: &pb.Post{
			Id:        post.ID.String(),
			CreatedAt: timestamppb.New(post.CreatedAt),
			UpdatedAt: timestamppb.New(post.UpdatedAt),
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
		pbPosts[i] = &pb.Post{
			Id:        post.ID.String(),
			CreatedAt: timestamppb.New(post.CreatedAt),
			UpdatedAt: timestamppb.New(post.UpdatedAt),
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

	err = s.db.IncrementPostViews(ctx, post.ID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't increment post views: %v - GetPost", err)
	}

	return &pb.GetPostByIDResponse{
		Post: &pb.Post{
			Id:        post.ID.String(),
			CreatedAt: timestamppb.New(post.CreatedAt),
			UpdatedAt: timestamppb.New(post.UpdatedAt),
			PostedBy:  post.PostedBy,
			Body:      post.Body,
			Views:     post.Views,
			Likes:     post.Likes,
		},
	}, nil
}

func (s *server) ChangePost(ctx context.Context, req *pb.ChangePostRequest) (*pb.ChangePostResponse, error) {
	postID, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "can't parse post id: %v - ChangePost", err)
	}

	accessToken, err := helper.GetBearerTokenFromGrpc(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "can't get bearer token: %v - ChangePost", err)
	}

	userID, err := helper.ValidateJWT(accessToken, s.tokenSecret)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token: %v - ChangePost", err)
	}

	post, err := s.db.GetPostByID(ctx, postID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't get post by id: %v - ChangePost", err)
	}

	if post.PostedBy != userID.String() {
		return nil, status.Errorf(codes.PermissionDenied, "you are not allowed to change this post: %v - ChangePost", err)
	}

	changedPostParams := database.ChangePostParams{
		Body: req.GetBody(),
		ID:   post.ID,
	}

	changedPost, err := s.db.ChangePost(ctx, changedPostParams)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "can't change post: %v - ChangePost", err)
	}

	createdAtProto := timestamppb.New(changedPost.CreatedAt)
	updatedAtProto := timestamppb.New(changedPost.UpdatedAt)

	return &pb.ChangePostResponse{
		Post: &pb.Post{
			Id:        changedPost.ID.String(),
			CreatedAt: createdAtProto,
			UpdatedAt: updatedAtProto,
			PostedBy:  changedPost.PostedBy,
			Body:      changedPost.Body,
			Views:     changedPost.Views,
			Likes:     changedPost.Likes,
		},
	}, nil
}

func (s *server) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {
	postID, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't parse bearer's token to uuid: %v - DeletePost", err)
	}

	accessToken, err := helper.GetBearerTokenFromGrpc(ctx)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "can't get bearer token from header: %v - DeletePost", err)
	}

	userID, err := helper.ValidateJWT(accessToken, s.tokenSecret)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "can't validate provided token: %v - DeletePost", err)
	}

	post, err := s.db.GetPostByID(ctx, postID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't get post using id: %v - DeletePost", err)
	}

	if post.ID != userID {
		return nil, status.Errorf(codes.PermissionDenied, "you are not allowed to change this post: %v - DeletePost", err)
	}

	_, err = s.db.DeletePost(ctx, postID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't delete post: %v - DeletePost", err)
	}

	return &pb.DeletePostResponse{
		Result: "Deleted successfully",
	}, nil
}

func (s *server) LikePost(ctx context.Context, req *pb.LikePostRequest) (*pb.LikePostResponse, error) {
	postID, err := uuid.Parse(req.GetPostId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't parse post id: %v - LikePost", err)
	}

	likePostParams := database.LikePostParams{
		ID:          postID,
		ArrayAppend: req.GetLikedBy(),
	}

	likedPost, err := s.db.LikePost(ctx, likePostParams)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't like post: %v - LikePost", err)
	}

	return &pb.LikePostResponse{
		Post: &pb.Post{
			Id:        likedPost.ID.String(),
			CreatedAt: timestamppb.New(likedPost.CreatedAt),
			UpdatedAt: timestamppb.New(likedPost.UpdatedAt),
			PostedBy:  likedPost.PostedBy,
			Body:      likedPost.Body,
			Likes:     likedPost.Likes,
			Views:     likedPost.Views,
			LikedBy:   likedPost.LikedBy,
		},
	}, nil
}

func (s *server) UnlikePost(ctx context.Context, req *pb.UnlikePostRequest) (*pb.UnlikePostResponse, error) {
	postID, err := uuid.Parse(req.GetPostId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't parse post id: %v - LikePost", err)
	}

	likePostParams := database.LikePostParams{
		ID:          postID,
		ArrayAppend: req.GetUnlikedBy(),
	}

	likedPost, err := s.db.LikePost(ctx, likePostParams)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't like post: %v - LikePost", err)
	}

	return &pb.UnlikePostResponse{
		Post: &pb.Post{
			Id:        likedPost.ID.String(),
			CreatedAt: timestamppb.New(likedPost.CreatedAt),
			UpdatedAt: timestamppb.New(likedPost.UpdatedAt),
			PostedBy:  likedPost.PostedBy,
			Body:      likedPost.Body,
			Likes:     likedPost.Likes,
			Views:     likedPost.Views,
			LikedBy:   likedPost.LikedBy,
		},
	}, nil
}

func (s *server) GetLikersFromPost(ctx context.Context, req *pb.GetLikersFromPostRequest) (*pb.GetLikersFromPostResponse, error) {
	postID, err := uuid.Parse(req.GetPostId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't parse post id: %v - GetLikersFromPost", err)
	}

	likers, err := s.db.GetLikersFromPost(ctx, postID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't get likers from post: %v - GetLikersFromPost", err)
	}

	likedBy := make([]string, len(likers))
	for i, liker := range likers {
		likedBy[i] = liker.(string)
	}

	return &pb.GetLikersFromPostResponse{
		LikedBy: likedBy,
	}, nil
}

func (s *server) CreateComment(ctx context.Context, req *pb.CreateCommentRequest) (*pb.CreateCommentResponse, error) {
	accessToken, err := helper.GetBearerTokenFromGrpc(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't get bearer token: %v - CreateComment", err)
	}

	userID, err := helper.ValidateJWT(accessToken, s.tokenSecret)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "can't get user id from token: %v - CreateComment", err)
	}

	postID, err := uuid.Parse(req.GetPostId())
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "can't parse post id: %v - CreateComment", err)
	} 

	createCommentParams := database.CreateCommentParams{
		ID: uuid.New(),
		PostID: postID,
		UserID: userID,
		CommentText: req.GetCommentText(),
	}

	comment, err := s.db.CreateComment(ctx, createCommentParams)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't create comment: %v - CreateComment", err)
	}

	return &pb.CreateCommentResponse{
		Comment: &pb.Comment{
			Id: comment.ID.String(),
			CreatedAt: timestamppb.New(comment.CreatedAt),
			PostId: comment.PostID.String(),
			UserId: comment.UserID.String(),		
			CommentText: comment.CommentText,	
		},
	}, nil
}

func (s *server) DeleteComment(ctx context.Context, req *pb.DeleteCommentRequest) (*pb.DeleteCommentResponse, error) {
	accessToken, err := helper.GetBearerTokenFromGrpc(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't get bearer token: %v - DeleteComment", err)
	}

	userID, err := helper.ValidateJWT(accessToken, s.tokenSecret)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "can't get user id from token: %v - CreateComment", err)
	}

	commentID, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't parse comment id: %v - DeleteComment", err)
	}

	commentParams, err := s.db.GetCommentByID(ctx, commentID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't get comment by id: %v - DeleteComment", err)
	}

	if commentParams.UserID != userID {
		return nil, status.Errorf(codes.PermissionDenied, "you are not allowed to delete this comment: %v - DeleteComment", err)
	}

	err = s.db.DeleteComment(ctx, commentParams.ID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't delete comment: %v - DeleteComment", err)
	}

	return &pb.DeleteCommentResponse{
		Status: "comment deleted successfully",
	}, nil
}

func (s *server) ReportPost(ctx context.Context, req *pb.ReportPostRequest) (*pb.ReportPostResponse, error) {
	postID, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't parse post id: %v - ReportPost", err)
	}

	accessToken, err := helper.GetBearerTokenFromGrpc(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't get bearer token: %v - ReportPost", err)
	}

	userID, err := helper.ValidateJWT(accessToken, s.tokenSecret)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't get user id from token: %v - ReportPost", err)
	}

	reportParams := database.ReportPostParams{
		ID:     postID,
		Reason: req.GetReason(),
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
			ReportedBy: userID.String(),
			Reason:     report.Reason,
		},
	}, nil
}

func (s *server) GetAllReports(ctx context.Context, req *pb.GetAllReportsRequest) (*pb.GetAllReportsResponse, error) {
	reports, err := s.db.GetAllReports(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't get reports from db: %v - GetAllReports", err)
	}

	pbReports := make([]*pb.ReportPost, len(reports))
	for i, report := range reports {
		reportedAt := timestamppb.New(report.ReportedAt)

		pbReports[i] = &pb.ReportPost{
			Id:         report.ID.String(),
			ReportedAt: reportedAt,
			ReportedBy: report.ReportedBy.String(),
			Reason:     report.Reason,
		}
	}

	return &pb.GetAllReportsResponse{
		ReportPost: pbReports,
	}, nil
}

func (s *server) ResetPosts(ctx context.Context, req *pb.ResetPostsRequest) (*pb.ResetPostsResponse, error) {
	err := s.db.ResetPosts(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't reset posts: %v - ResetPosts", err)
	}

	return &pb.ResetPostsResponse{
		Status: "All posts deleted successfully",
	}, nil
}
