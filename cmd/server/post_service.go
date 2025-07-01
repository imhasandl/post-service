package server

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/imhasandl/post-service/cmd/auth"
	"github.com/imhasandl/post-service/cmd/helper"
	"github.com/imhasandl/post-service/internal/database"
	"github.com/imhasandl/post-service/internal/rabbitmq"
	"github.com/imhasandl/post-service/internal/redis"
	pb "github.com/imhasandl/post-service/protos"
	"github.com/streadway/amqp"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type server struct {
	pb.UnimplementedPostServiceServer
	db          *database.Queries
	tokenSecret string
	rabbitmq    *rabbitmq.RabbitMQ
	redis       *redis.RedisClient
}

func NewServer(db *database.Queries, tokenSecret string, rabbitmq *rabbitmq.RabbitMQ, redis *redis.RedisClient) *server {
	return &server{
		pb.UnimplementedPostServiceServer{},
		db,
		tokenSecret,
		rabbitmq,
		redis,
	}
}

func (s *server) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.CreatePostResponse, error) {
	accessToken, err := auth.GetBearerTokenFromGrpc(ctx)
	if err != nil {
		return nil, helper.RespondWithErrorGRPC(ctx, codes.Internal, "can't get bearer token: CreatePost", err)
	}

	userID, err := auth.ValidateJWT(accessToken, s.tokenSecret)
	if err != nil {
		return nil, helper.RespondWithErrorGRPC(ctx, codes.Internal, "can't get user id from token: CreatePost", err)
	}

	postParams := database.CreatePostParams{
		ID:       uuid.New(),
		PostedBy: userID,
		Body:     req.GetBody(),
	}

	post, err := s.db.CreatePost(ctx, postParams)
	if err != nil {
		return nil, helper.RespondWithErrorGRPC(ctx, codes.Internal, "can't create post to database: CreatePost", err)
	}

	userSubscribers, err := s.db.GetSubscribers(ctx, post.PostedBy)
	if err != nil {
		return nil, helper.RespondWithErrorGRPC(ctx, codes.Internal, "can't get users subscribers from db - CreatePost", err)
	}

	for i := 0; i < len(userSubscribers); i++ {
		messageJSON, err := json.Marshal(map[string]any{
			"title":           "New Notification",
			"sender_username": userID,
			"receiver_id":     userSubscribers[i][0],
			"content":         fmt.Sprintf("New post from %v", userID.String()),
			"sent_at":         post.CreatedAt,
		})
		if err != nil {
			return nil, helper.RespondWithErrorGRPC(ctx, codes.Internal, "can't marshal notification message - CreatePost", err)
		}

		err = s.rabbitmq.Channel.Publish(
			rabbitmq.ExchangeName, // exchange
			rabbitmq.RoutingKey,   // routing key
			false,                 // mandatory
			false,                 // immediate
			amqp.Publishing{
				ContentType: "application/json",
				Body:        messageJSON,
			})
		if err != nil {
			return nil, helper.RespondWithErrorGRPC(ctx, codes.Internal, "can't publish message to RabbitMQ - CreatePost", err)
		}
	}

	return &pb.CreatePostResponse{
		Post: &pb.Post{
			Id:        post.ID.String(),
			CreatedAt: timestamppb.New(post.CreatedAt),
			UpdatedAt: timestamppb.New(post.UpdatedAt),
			PostedBy:  post.PostedBy.String(),
			Body:      post.Body,
			Likes:     post.Likes,
			Views:     post.Views,
		},
	}, nil
}

func (s *server) GetAllPosts(ctx context.Context, req *pb.GetAllPostsRequest) (*pb.GetAllPostsResponse, error) {
	posts, err := s.db.GetAllPosts(ctx)
	if err != nil {
		return nil, helper.RespondWithErrorGRPC(ctx, codes.Internal, "can't get all posts from db: GetAllPosts", err)
	}

	pbPosts := make([]*pb.Post, len(posts))
	for i, post := range posts {
		pbPosts[i] = &pb.Post{
			Id:        post.ID.String(),
			CreatedAt: timestamppb.New(post.CreatedAt),
			UpdatedAt: timestamppb.New(post.UpdatedAt),
			PostedBy:  post.PostedBy.String(),
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
		return nil, helper.RespondWithErrorGRPC(ctx, codes.Internal, "can't parse id from request: GetPostByID", err)
	}

	// Try to get post from Redis cache
	cacheKey := "post:" + postID.String()
	cachedPost, err := s.redis.Get(ctx, cacheKey)

	// If found in cache, unmarshal and return
	if err == nil {
		var post database.Post
		if err := json.Unmarshal([]byte(cachedPost), &post); err == nil {
			return &pb.GetPostByIDResponse{
				Post: &pb.Post{
					Id:        post.ID.String(),
					CreatedAt: timestamppb.New(post.CreatedAt),
					UpdatedAt: timestamppb.New(post.UpdatedAt),
					PostedBy:  post.PostedBy.String(),
					Body:      post.Body,
					Views:     post.Views,
					Likes:     post.Likes,
				},
			}, nil
		}
	}

	// Not found in cache, get from database
	post, err := s.db.GetPostByID(ctx, postID)
	if err != nil {
		return nil, helper.RespondWithErrorGRPC(ctx, codes.Internal, "can't get post from db: GetPostByID", err)
	}

	err = s.db.IncrementPostViews(ctx, post.ID)
	if err != nil {
		return nil, helper.RespondWithErrorGRPC(ctx, codes.Internal, "can't increment post views: GetPost", err)
	}

	// Cache the post for future requests (15 minutes)
	postBytes, err := json.Marshal(post)
	if err == nil {
		s.redis.Set(ctx, cacheKey, postBytes, 15*time.Minute)
	}

	return &pb.GetPostByIDResponse{
		Post: &pb.Post{
			Id:        post.ID.String(),
			CreatedAt: timestamppb.New(post.CreatedAt),
			UpdatedAt: timestamppb.New(post.UpdatedAt),
			PostedBy:  post.PostedBy.String(),
			Body:      post.Body,
			Views:     post.Views,
			Likes:     post.Likes,
		},
	}, nil
}

func (s *server) ChangePost(ctx context.Context, req *pb.ChangePostRequest) (*pb.ChangePostResponse, error) {
	postID, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, helper.RespondWithErrorGRPC(ctx, codes.InvalidArgument, "can't parse post id: ChangePost", err)
	}

	accessToken, err := auth.GetBearerTokenFromGrpc(ctx)
	if err != nil {
		return nil, helper.RespondWithErrorGRPC(ctx, codes.Unauthenticated, "can't get bearer token: ChangePost", err)
	}

	userID, err := auth.ValidateJWT(accessToken, s.tokenSecret)
	if err != nil {
		return nil, helper.RespondWithErrorGRPC(ctx, codes.Unauthenticated, "invalid token: ChangePost", err)
	}

	post, err := s.db.GetPostByID(ctx, postID)
	if err != nil {
		return nil, helper.RespondWithErrorGRPC(ctx, codes.Internal, "can't get post by id: ChangePost", err)
	}

	if post.PostedBy != userID {
		return nil, helper.RespondWithErrorGRPC(ctx, codes.PermissionDenied, "you are not allowed to change this post: ChangePost", err)
	}

	changedPostParams := database.ChangePostParams{
		Body: req.GetBody(),
		ID:   post.ID,
	}

	changedPost, err := s.db.ChangePost(ctx, changedPostParams)
	if err != nil {
		return nil, helper.RespondWithErrorGRPC(ctx, codes.PermissionDenied, "can't change post: ChangePost", err)
	}

	// Invalidate cache when post changes
	cacheKey := "post:" + changedPost.ID.String()
	s.redis.Delete(ctx, cacheKey)

	createdAtProto := timestamppb.New(changedPost.CreatedAt)
	updatedAtProto := timestamppb.New(changedPost.UpdatedAt)

	return &pb.ChangePostResponse{
		Post: &pb.Post{
			Id:        changedPost.ID.String(),
			CreatedAt: createdAtProto,
			UpdatedAt: updatedAtProto,
			PostedBy:  changedPost.PostedBy.String(),
			Body:      changedPost.Body,
			Views:     changedPost.Views,
			Likes:     changedPost.Likes,
		},
	}, nil
}

func (s *server) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {
	postID, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, helper.RespondWithErrorGRPC(ctx, codes.Internal, "can't parse bearer's token to uuid: DeletePost", err)
	}

	accessToken, err := auth.GetBearerTokenFromGrpc(ctx)
	if err != nil {
		return nil, helper.RespondWithErrorGRPC(ctx, codes.PermissionDenied, "can't get bearer token from header: DeletePost", err)
	}

	userID, err := auth.ValidateJWT(accessToken, s.tokenSecret)
	if err != nil {
		return nil, helper.RespondWithErrorGRPC(ctx, codes.PermissionDenied, "can't validate provided token: DeletePost", err)
	}

	post, err := s.db.GetPostByID(ctx, postID)
	if err != nil {
		return nil, helper.RespondWithErrorGRPC(ctx, codes.Internal, "can't get post using id: DeletePost", err)
	}

	if post.ID != userID {
		return nil, helper.RespondWithErrorGRPC(ctx, codes.PermissionDenied, "you are not allowed to change this post: DeletePost", err)
	}

	_, err = s.db.DeletePost(ctx, postID)
	if err != nil {
		return nil, helper.RespondWithErrorGRPC(ctx, codes.Internal, "can't delete post: DeletePost", err)
	}

	return &pb.DeletePostResponse{
		Result: "Deleted successfully",
	}, nil
}

func (s *server) LikePost(ctx context.Context, req *pb.LikePostRequest) (*pb.LikePostResponse, error) {
	postID, err := uuid.Parse(req.GetPostId())
	if err != nil {
		return nil, helper.RespondWithErrorGRPC(ctx, codes.Internal, "can't parse post id: LikePost", err)
	}

	likePostParams := database.LikePostParams{
		ID:          postID,
		ArrayAppend: req.GetLikedBy(),
	}

	post, err := s.db.LikePost(ctx, likePostParams)
	if err != nil {
		return nil, helper.RespondWithErrorGRPC(ctx, codes.Internal, "can't like post: LikePost", err)
	}
	messageJSON, err := json.Marshal(map[string]interface{}{
		"title":           "New Notification",
		"sender_username": req.LikedBy,
		"receiver_id":     post.PostedBy.String(),
		"content":         fmt.Sprintf("%v liked your post", req.LikedBy),
		"sent_at":         post.CreatedAt.String(),
	})
	if err != nil {
		return nil, helper.RespondWithErrorGRPC(ctx, codes.InvalidArgument, "can't create json message to send to rabbbitmq - LikePost", err)
	}

	err = s.rabbitmq.Channel.Publish(
		rabbitmq.ExchangeName, // exchange
		rabbitmq.RoutingKey,   // routing key
		false,                 // mandatory
		false,                 // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        messageJSON,
		},
	)
	if err != nil {
		return nil, helper.RespondWithErrorGRPC(ctx, codes.Internal, "can't send json to notification queue", err)
	}

	return &pb.LikePostResponse{
		Post: &pb.Post{
			Id:        post.ID.String(),
			CreatedAt: timestamppb.New(post.CreatedAt),
			UpdatedAt: timestamppb.New(post.UpdatedAt),
			PostedBy:  post.PostedBy.String(),
			Body:      post.Body,
			Likes:     post.Likes,
			Views:     post.Views,
			LikedBy:   post.LikedBy,
		},
	}, nil
}

func (s *server) UnlikePost(ctx context.Context, req *pb.UnlikePostRequest) (*pb.UnlikePostResponse, error) {
	postID, err := uuid.Parse(req.GetPostId())
	if err != nil {
		return nil, helper.RespondWithErrorGRPC(ctx, codes.Internal, "can't parse post id: UnlikePost", err)
	}

	likePostParams := database.LikePostParams{
		ID:          postID,
		ArrayAppend: req.GetUnlikedBy(),
	}

	likedPost, err := s.db.LikePost(ctx, likePostParams)
	if err != nil {
		return nil, helper.RespondWithErrorGRPC(ctx, codes.Internal, "can't like post: UnlikePost", err)
	}

	return &pb.UnlikePostResponse{
		Post: &pb.Post{
			Id:        likedPost.ID.String(),
			CreatedAt: timestamppb.New(likedPost.CreatedAt),
			UpdatedAt: timestamppb.New(likedPost.UpdatedAt),
			PostedBy:  likedPost.PostedBy.String(),
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
		return nil, helper.RespondWithErrorGRPC(ctx, codes.Internal, "can't parse post id: GetLikersFromPost", err)
	}

	likers, err := s.db.GetLikersFromPost(ctx, postID)
	if err != nil {
		return nil, helper.RespondWithErrorGRPC(ctx, codes.Internal, "can't get likers from post: GetLikersFromPost", err)
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
	accessToken, err := auth.GetBearerTokenFromGrpc(ctx)
	if err != nil {
		return nil, helper.RespondWithErrorGRPC(ctx, codes.Internal, "can't get bearer token: CreateComment", err)
	}

	userID, err := auth.ValidateJWT(accessToken, s.tokenSecret)
	if err != nil {
		return nil, helper.RespondWithErrorGRPC(ctx, codes.Unauthenticated, "can't get user id from token: CreateComment", err)
	}

	postID, err := uuid.Parse(req.GetPostId())
	if err != nil {
		return nil, helper.RespondWithErrorGRPC(ctx, codes.Unauthenticated, "can't parse post id: CreateComment", err)
	}

	post, err := s.db.GetPostByID(ctx, postID)
	if err != nil {
		return nil, helper.RespondWithErrorGRPC(ctx, codes.Internal, "can't get post from db using it's id - CreateComment", err)
	}

	createCommentParams := database.CreateCommentParams{
		ID:          uuid.New(),
		PostID:      postID,
		UserID:      userID,
		CommentText: req.GetCommentText(),
	}

	comment, err := s.db.CreateComment(ctx, createCommentParams)
	if err != nil {
		return nil, helper.RespondWithErrorGRPC(ctx, codes.Internal, "can't create comment: CreateComment", err)
	}

	// Create JSON message
	messageJSON, err := json.Marshal(map[string]interface{}{
		"title":           "New comment on your post",
		"sender_username": userID,
		"receiver_id":     post.PostedBy,
		"content":         req.GetCommentText(),
		"sent_at":         comment.CreatedAt,
	})
	if err != nil {
		return nil, helper.RespondWithErrorGRPC(ctx, codes.Internal, "can't marshal message to JSON - SendMessage", err)
	}

	// Publish message to RabbitMQ
	err = s.rabbitmq.Channel.Publish(
		rabbitmq.ExchangeName, // exchange
		rabbitmq.RoutingKey,   // routing key
		false,                 // mandatory
		false,                 // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        messageJSON,
		})
	if err != nil {
		return nil, helper.RespondWithErrorGRPC(ctx, codes.Internal, "can't publish message to RabbitMQ - SendMessage", err)
	}

	return &pb.CreateCommentResponse{
		Comment: &pb.Comment{
			Id:          comment.ID.String(),
			CreatedAt:   timestamppb.New(comment.CreatedAt),
			PostId:      comment.PostID.String(),
			UserId:      comment.UserID.String(),
			CommentText: comment.CommentText,
		},
	}, nil
}

func (s *server) DeleteComment(ctx context.Context, req *pb.DeleteCommentRequest) (*pb.DeleteCommentResponse, error) {
	accessToken, err := auth.GetBearerTokenFromGrpc(ctx)
	if err != nil {
		return nil, helper.RespondWithErrorGRPC(ctx, codes.Internal, "can't get bearer token: DeleteComment", err)
	}

	userID, err := auth.ValidateJWT(accessToken, s.tokenSecret)
	if err != nil {
		return nil, helper.RespondWithErrorGRPC(ctx, codes.Unauthenticated, "can't get user id from token: DeleteComment", err)
	}

	commentID, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, helper.RespondWithErrorGRPC(ctx, codes.Internal, "can't parse comment id: DeleteComment", err)
	}

	commentParams, err := s.db.GetCommentByID(ctx, commentID)
	if err != nil {
		return nil, helper.RespondWithErrorGRPC(ctx, codes.Internal, "can't get comment by id: DeleteComment", err)
	}

	if commentParams.UserID != userID {
		return nil, helper.RespondWithErrorGRPC(ctx, codes.PermissionDenied, "you are not allowed to delete this comment: DeleteComment", err)
	}

	err = s.db.DeleteComment(ctx, commentParams.ID)
	if err != nil {
		return nil, helper.RespondWithErrorGRPC(ctx, codes.Internal, "can't delete comment: DeleteComment", err)
	}

	return &pb.DeleteCommentResponse{
		Status: "comment deleted successfully",
	}, nil
}

func (s *server) ReportPost(ctx context.Context, req *pb.ReportPostRequest) (*pb.ReportPostResponse, error) {
	postID, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, helper.RespondWithErrorGRPC(ctx, codes.Internal, "can't parse post id: ReportPost", err)
	}

	accessToken, err := auth.GetBearerTokenFromGrpc(ctx)
	if err != nil {
		return nil, helper.RespondWithErrorGRPC(ctx, codes.Internal, "can't get bearer token: ReportPost", err)
	}

	userID, err := auth.ValidateJWT(accessToken, s.tokenSecret)
	if err != nil {
		return nil, helper.RespondWithErrorGRPC(ctx, codes.Internal, "can't get user id from token: ReportPost", err)
	}

	reportParams := database.ReportPostParams{
		ID:     postID,
		Reason: req.GetReason(),
	}

	report, err := s.db.ReportPost(ctx, reportParams)
	if err != nil {
		return nil, helper.RespondWithErrorGRPC(ctx, codes.Internal, "can't report post to db: ReportPost", err)
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
		return nil, helper.RespondWithErrorGRPC(ctx, codes.Internal, "can't get reports from db: GetAllReports", err)
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
		return nil, helper.RespondWithErrorGRPC(ctx, codes.Internal, "can't reset posts: ResetPosts", err)
	}

	return &pb.ResetPostsResponse{
		Status: "All posts deleted successfully",
	}, nil
}
