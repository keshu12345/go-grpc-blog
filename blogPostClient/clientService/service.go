package clientService

import (
	"context"

	"github.com/keshu12345/blog/pb"
)

type ClientService struct {
}

// Create Post
func (cleintService ClientService) CreatePost(ctx context.Context,
	client pb.BlogServiceClient) (*pb.Post, error) {
	return client.CreatePost(ctx, &pb.CreatePostRequest{
		Title:   "New Post",
		Content: "This is a new post.",
		Author:  "Author A",
		Tags:    []string{"tag1", "tag2"},
	})
}

func (cleintService ClientService) ReadPost(ctx context.Context, client pb.BlogServiceClient,
	post *pb.Post) (*pb.Post, error) {
	return client.ReadPost(ctx, &pb.ReadPostRequest{PostId: post.PostId})
}

func (cleintService ClientService) UpdatePost(ctx context.Context, client pb.BlogServiceClient,
	post *pb.Post) (*pb.Post, error) {
	return client.UpdatePost(ctx, &pb.UpdatePostRequest{
		PostId:  post.PostId,
		Title:   "Updated Post",
		Content: "This is an updated post content.",
		Author:  "Author B",
		Tags:    []string{"tag3", "tag4"},
	})
}

func (cleintService ClientService) DeletePost(ctx context.Context, client pb.BlogServiceClient,
	post *pb.Post) (*pb.DeletePostResponse, error) {

	return client.DeletePost(ctx, &pb.DeletePostRequest{
		PostId: post.PostId,
	})
}
