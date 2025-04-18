// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: post.proto

package protos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	PostService_CreatePost_FullMethodName        = "/post.PostService/CreatePost"
	PostService_ChangePost_FullMethodName        = "/post.PostService/ChangePost"
	PostService_DeletePost_FullMethodName        = "/post.PostService/DeletePost"
	PostService_LikePost_FullMethodName          = "/post.PostService/LikePost"
	PostService_UnlikePost_FullMethodName        = "/post.PostService/UnlikePost"
	PostService_GetLikersFromPost_FullMethodName = "/post.PostService/GetLikersFromPost"
	PostService_GetPostByID_FullMethodName       = "/post.PostService/GetPostByID"
	PostService_GetAllPosts_FullMethodName       = "/post.PostService/GetAllPosts"
	PostService_ReportPost_FullMethodName        = "/post.PostService/ReportPost"
	PostService_GetAllReports_FullMethodName     = "/post.PostService/GetAllReports"
	PostService_CreateComment_FullMethodName     = "/post.PostService/CreateComment"
	PostService_DeleteComment_FullMethodName     = "/post.PostService/DeleteComment"
	PostService_ResetPosts_FullMethodName        = "/post.PostService/ResetPosts"
)

// PostServiceClient is the client API for PostService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostServiceClient interface {
	CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*CreatePostResponse, error)
	ChangePost(ctx context.Context, in *ChangePostRequest, opts ...grpc.CallOption) (*ChangePostResponse, error)
	DeletePost(ctx context.Context, in *DeletePostRequest, opts ...grpc.CallOption) (*DeletePostResponse, error)
	LikePost(ctx context.Context, in *LikePostRequest, opts ...grpc.CallOption) (*LikePostResponse, error)
	UnlikePost(ctx context.Context, in *UnlikePostRequest, opts ...grpc.CallOption) (*UnlikePostResponse, error)
	GetLikersFromPost(ctx context.Context, in *GetLikersFromPostRequest, opts ...grpc.CallOption) (*GetLikersFromPostResponse, error)
	GetPostByID(ctx context.Context, in *GetPostByIDRequest, opts ...grpc.CallOption) (*GetPostByIDResponse, error)
	GetAllPosts(ctx context.Context, in *GetAllPostsRequest, opts ...grpc.CallOption) (*GetAllPostsResponse, error)
	ReportPost(ctx context.Context, in *ReportPostRequest, opts ...grpc.CallOption) (*ReportPostResponse, error)
	GetAllReports(ctx context.Context, in *GetAllReportsRequest, opts ...grpc.CallOption) (*GetAllReportsResponse, error)
	CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*CreateCommentResponse, error)
	DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*DeleteCommentResponse, error)
	ResetPosts(ctx context.Context, in *ResetPostsRequest, opts ...grpc.CallOption) (*ResetPostsResponse, error)
}

type postServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPostServiceClient(cc grpc.ClientConnInterface) PostServiceClient {
	return &postServiceClient{cc}
}

func (c *postServiceClient) CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*CreatePostResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreatePostResponse)
	err := c.cc.Invoke(ctx, PostService_CreatePost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) ChangePost(ctx context.Context, in *ChangePostRequest, opts ...grpc.CallOption) (*ChangePostResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChangePostResponse)
	err := c.cc.Invoke(ctx, PostService_ChangePost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) DeletePost(ctx context.Context, in *DeletePostRequest, opts ...grpc.CallOption) (*DeletePostResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeletePostResponse)
	err := c.cc.Invoke(ctx, PostService_DeletePost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) LikePost(ctx context.Context, in *LikePostRequest, opts ...grpc.CallOption) (*LikePostResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LikePostResponse)
	err := c.cc.Invoke(ctx, PostService_LikePost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) UnlikePost(ctx context.Context, in *UnlikePostRequest, opts ...grpc.CallOption) (*UnlikePostResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UnlikePostResponse)
	err := c.cc.Invoke(ctx, PostService_UnlikePost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetLikersFromPost(ctx context.Context, in *GetLikersFromPostRequest, opts ...grpc.CallOption) (*GetLikersFromPostResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetLikersFromPostResponse)
	err := c.cc.Invoke(ctx, PostService_GetLikersFromPost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetPostByID(ctx context.Context, in *GetPostByIDRequest, opts ...grpc.CallOption) (*GetPostByIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPostByIDResponse)
	err := c.cc.Invoke(ctx, PostService_GetPostByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetAllPosts(ctx context.Context, in *GetAllPostsRequest, opts ...grpc.CallOption) (*GetAllPostsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllPostsResponse)
	err := c.cc.Invoke(ctx, PostService_GetAllPosts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) ReportPost(ctx context.Context, in *ReportPostRequest, opts ...grpc.CallOption) (*ReportPostResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReportPostResponse)
	err := c.cc.Invoke(ctx, PostService_ReportPost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetAllReports(ctx context.Context, in *GetAllReportsRequest, opts ...grpc.CallOption) (*GetAllReportsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllReportsResponse)
	err := c.cc.Invoke(ctx, PostService_GetAllReports_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*CreateCommentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateCommentResponse)
	err := c.cc.Invoke(ctx, PostService_CreateComment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*DeleteCommentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteCommentResponse)
	err := c.cc.Invoke(ctx, PostService_DeleteComment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) ResetPosts(ctx context.Context, in *ResetPostsRequest, opts ...grpc.CallOption) (*ResetPostsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResetPostsResponse)
	err := c.cc.Invoke(ctx, PostService_ResetPosts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostServiceServer is the server API for PostService service.
// All implementations must embed UnimplementedPostServiceServer
// for forward compatibility.
type PostServiceServer interface {
	CreatePost(context.Context, *CreatePostRequest) (*CreatePostResponse, error)
	ChangePost(context.Context, *ChangePostRequest) (*ChangePostResponse, error)
	DeletePost(context.Context, *DeletePostRequest) (*DeletePostResponse, error)
	LikePost(context.Context, *LikePostRequest) (*LikePostResponse, error)
	UnlikePost(context.Context, *UnlikePostRequest) (*UnlikePostResponse, error)
	GetLikersFromPost(context.Context, *GetLikersFromPostRequest) (*GetLikersFromPostResponse, error)
	GetPostByID(context.Context, *GetPostByIDRequest) (*GetPostByIDResponse, error)
	GetAllPosts(context.Context, *GetAllPostsRequest) (*GetAllPostsResponse, error)
	ReportPost(context.Context, *ReportPostRequest) (*ReportPostResponse, error)
	GetAllReports(context.Context, *GetAllReportsRequest) (*GetAllReportsResponse, error)
	CreateComment(context.Context, *CreateCommentRequest) (*CreateCommentResponse, error)
	DeleteComment(context.Context, *DeleteCommentRequest) (*DeleteCommentResponse, error)
	ResetPosts(context.Context, *ResetPostsRequest) (*ResetPostsResponse, error)
	mustEmbedUnimplementedPostServiceServer()
}

// UnimplementedPostServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPostServiceServer struct{}

func (UnimplementedPostServiceServer) CreatePost(context.Context, *CreatePostRequest) (*CreatePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePost not implemented")
}
func (UnimplementedPostServiceServer) ChangePost(context.Context, *ChangePostRequest) (*ChangePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePost not implemented")
}
func (UnimplementedPostServiceServer) DeletePost(context.Context, *DeletePostRequest) (*DeletePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePost not implemented")
}
func (UnimplementedPostServiceServer) LikePost(context.Context, *LikePostRequest) (*LikePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LikePost not implemented")
}
func (UnimplementedPostServiceServer) UnlikePost(context.Context, *UnlikePostRequest) (*UnlikePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnlikePost not implemented")
}
func (UnimplementedPostServiceServer) GetLikersFromPost(context.Context, *GetLikersFromPostRequest) (*GetLikersFromPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLikersFromPost not implemented")
}
func (UnimplementedPostServiceServer) GetPostByID(context.Context, *GetPostByIDRequest) (*GetPostByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPostByID not implemented")
}
func (UnimplementedPostServiceServer) GetAllPosts(context.Context, *GetAllPostsRequest) (*GetAllPostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPosts not implemented")
}
func (UnimplementedPostServiceServer) ReportPost(context.Context, *ReportPostRequest) (*ReportPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportPost not implemented")
}
func (UnimplementedPostServiceServer) GetAllReports(context.Context, *GetAllReportsRequest) (*GetAllReportsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllReports not implemented")
}
func (UnimplementedPostServiceServer) CreateComment(context.Context, *CreateCommentRequest) (*CreateCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateComment not implemented")
}
func (UnimplementedPostServiceServer) DeleteComment(context.Context, *DeleteCommentRequest) (*DeleteCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComment not implemented")
}
func (UnimplementedPostServiceServer) ResetPosts(context.Context, *ResetPostsRequest) (*ResetPostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPosts not implemented")
}
func (UnimplementedPostServiceServer) mustEmbedUnimplementedPostServiceServer() {}
func (UnimplementedPostServiceServer) testEmbeddedByValue()                     {}

// UnsafePostServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostServiceServer will
// result in compilation errors.
type UnsafePostServiceServer interface {
	mustEmbedUnimplementedPostServiceServer()
}

func RegisterPostServiceServer(s grpc.ServiceRegistrar, srv PostServiceServer) {
	// If the following call pancis, it indicates UnimplementedPostServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PostService_ServiceDesc, srv)
}

func _PostService_CreatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).CreatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_CreatePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).CreatePost(ctx, req.(*CreatePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_ChangePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).ChangePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_ChangePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).ChangePost(ctx, req.(*ChangePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_DeletePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).DeletePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_DeletePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).DeletePost(ctx, req.(*DeletePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_LikePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).LikePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_LikePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).LikePost(ctx, req.(*LikePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_UnlikePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnlikePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).UnlikePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_UnlikePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).UnlikePost(ctx, req.(*UnlikePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetLikersFromPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLikersFromPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetLikersFromPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_GetLikersFromPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetLikersFromPost(ctx, req.(*GetLikersFromPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetPostByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetPostByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_GetPostByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetPostByID(ctx, req.(*GetPostByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetAllPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllPostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetAllPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_GetAllPosts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetAllPosts(ctx, req.(*GetAllPostsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_ReportPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).ReportPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_ReportPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).ReportPost(ctx, req.(*ReportPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetAllReports_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllReportsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetAllReports(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_GetAllReports_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetAllReports(ctx, req.(*GetAllReportsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_CreateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).CreateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_CreateComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).CreateComment(ctx, req.(*CreateCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_DeleteComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).DeleteComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_DeleteComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).DeleteComment(ctx, req.(*DeleteCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_ResetPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetPostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).ResetPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_ResetPosts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).ResetPosts(ctx, req.(*ResetPostsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PostService_ServiceDesc is the grpc.ServiceDesc for PostService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PostService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "post.PostService",
	HandlerType: (*PostServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePost",
			Handler:    _PostService_CreatePost_Handler,
		},
		{
			MethodName: "ChangePost",
			Handler:    _PostService_ChangePost_Handler,
		},
		{
			MethodName: "DeletePost",
			Handler:    _PostService_DeletePost_Handler,
		},
		{
			MethodName: "LikePost",
			Handler:    _PostService_LikePost_Handler,
		},
		{
			MethodName: "UnlikePost",
			Handler:    _PostService_UnlikePost_Handler,
		},
		{
			MethodName: "GetLikersFromPost",
			Handler:    _PostService_GetLikersFromPost_Handler,
		},
		{
			MethodName: "GetPostByID",
			Handler:    _PostService_GetPostByID_Handler,
		},
		{
			MethodName: "GetAllPosts",
			Handler:    _PostService_GetAllPosts_Handler,
		},
		{
			MethodName: "ReportPost",
			Handler:    _PostService_ReportPost_Handler,
		},
		{
			MethodName: "GetAllReports",
			Handler:    _PostService_GetAllReports_Handler,
		},
		{
			MethodName: "CreateComment",
			Handler:    _PostService_CreateComment_Handler,
		},
		{
			MethodName: "DeleteComment",
			Handler:    _PostService_DeleteComment_Handler,
		},
		{
			MethodName: "ResetPosts",
			Handler:    _PostService_ResetPosts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "post.proto",
}
