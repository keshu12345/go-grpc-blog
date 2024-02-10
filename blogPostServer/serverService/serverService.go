package serverService

import (
	"context"
	"fmt"
	"sync"

	"github.com/keshu12345/blog/pb"
)

type Server struct {
	pb.UnimplementedBlogServiceServer
	mu    sync.Mutex
	Posts map[string]*pb.Post
}

func (s *Server) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.Post, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	postID := fmt.Sprintf("POST_%d", len(s.Posts)+1)
	post := &pb.Post{
		PostId:          postID,
		Title:           req.Title,
		Content:         req.Content,
		Author:          req.Author,
		PublicationDate: req.PublicationDate,
		Tags:            req.Tags,
	}

	s.Posts[postID] = post

	return post, nil
}
func (s *Server) ReadPost(ctx context.Context, req *pb.ReadPostRequest) (*pb.Post, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	post, exists := s.Posts[req.PostId]
	if !exists {
		return nil, fmt.Errorf("post with ID %s not found", req.PostId)
	}

	return post, nil
}

func (s *Server) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.Post, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	post, exists := s.Posts[req.PostId]
	if !exists {
		return nil, fmt.Errorf("post with ID %s not found", req.PostId)
	}

	// Update the post fields
	post.Title = req.Title
	post.Content = req.Content
	post.Author = req.Author
	post.Tags = req.Tags

	s.Posts[req.PostId] = post

	return post, nil
}

func (s *Server) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.Posts[req.PostId]
	if !exists {
		return nil, fmt.Errorf("post with ID %s not found", req.PostId)
	}

	delete(s.Posts, req.PostId)

	return &pb.DeletePostResponse{
		Success: true,
	}, nil
}
