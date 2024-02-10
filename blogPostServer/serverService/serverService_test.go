package serverService

import (
	"context"
	"sync"
	"testing"

	"github.com/keshu12345/blog/pb"
	"github.com/test-go/testify/assert"
	"github.com/test-go/testify/require"
)

func TestServer_CreateReadPost(t *testing.T) {
	s := &Server{
		Posts: make(map[string]*pb.Post),
		mu:    sync.Mutex{},
	}

	ctx := context.Background()
	newPost, err := s.CreatePost(ctx, &pb.CreatePostRequest{
		Title:   "Test Post",
		Content: "This is a test post content.",
		Author:  "Test Author",
		Tags:    []string{"test1", "test2"},
	})
	if err != nil {
		t.Fatalf("Failed to create post: %v", err)
	}

	assert.NotNil(t, newPost.PostId, "The post ID should not be nil")
	readPost, err := s.ReadPost(ctx, &pb.ReadPostRequest{PostId: newPost.PostId})
	if err != nil {
		t.Fatalf("Failed to read post: %v", err)
	}

	assert.Equal(t, newPost, readPost, "The created post and read post should be the same")
	assert.Equal(t, newPost.Title, readPost.Title)
	assert.Equal(t, newPost.Author, readPost.Author)
	assert.Equal(t, newPost.Content, readPost.Content)
	assert.Equal(t, newPost.Tags, readPost.Tags)
}

func TestServer_ReadPost(t *testing.T) {

	s := &Server{
		Posts: map[string]*pb.Post{
			"POST_1": {
				PostId:  "POST_1",
				Title:   "Existing Post",
				Content: "This is an existing post content.",
				Author:  "Existing Author",
				Tags:    []string{"existing1", "existing2"},
			},
		},
		mu: sync.Mutex{},
	}

	ctx := context.Background()

	tests := []struct {
		name    string
		postID  string
		wantErr bool
	}{
		{
			name:    "Existing Post",
			postID:  "POST_1",
			wantErr: false,
		},
		{
			name:    "Non-existing Post",
			postID:  "POST_2",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			post, err := s.ReadPost(ctx, &pb.ReadPostRequest{PostId: tt.postID})

			if tt.wantErr {
				require.Error(t, err, "ReadPost should return an error for non-existing post")
				require.Nil(t, post, "Post should be nil when an error is returned")
			} else {
				require.NoError(t, err, "ReadPost should not return an error for existing post")
				require.NotNil(t, post, "Post should not be nil when no error is returned")
				require.Equal(t, tt.postID, post.PostId, "ReadPost should return the correct post ID")
			}
		})
	}
}

func TestServer_UpdatePost(t *testing.T) {
	s := &Server{
		Posts: map[string]*pb.Post{
			"POST_1": {
				PostId:  "POST_1",
				Title:   "Original Title",
				Content: "Original content.",
				Author:  "Original Author",
				Tags:    []string{"tag1", "tag2"},
			},
		},
		mu: sync.Mutex{},
	}

	ctx := context.Background()
	updatedPost, err := s.UpdatePost(ctx, &pb.UpdatePostRequest{
		PostId:  "POST_1",
		Title:   "Updated Title",
		Content: "Updated content.",
		Author:  "Updated Author",
		Tags:    []string{"tag3", "tag4"},
	})
	require.NoError(t, err, "Updating an existing post should not result in an error")
	require.NotNil(t, updatedPost, "Updated post should not be nil")
	require.Equal(t, "Updated Title", updatedPost.Title, "Post title should be updated")
	require.Equal(t, "Updated content.", updatedPost.Content, "Post content should be updated")
	require.Equal(t, "Updated Author", updatedPost.Author, "Post author should be updated")

	_, err = s.UpdatePost(ctx, &pb.UpdatePostRequest{
		PostId: "POST_999",
		Title:  "Non-Existent Post Title",
	})
	require.Error(t, err, "Updating a non-existing post should result in an error")
}

func TestServer_DeletePost(t *testing.T) {
	s := &Server{
		Posts: map[string]*pb.Post{
			"POST_1": {
				PostId:  "POST_1",
				Title:   "Post One",
				Content: "Content for post one.",
				Author:  "Author One",
			},
			"POST_2": {
				PostId:  "POST_2",
				Title:   "Post Two",
				Content: "Content for post two.",
				Author:  "Author Two",
			},
		},
		mu: sync.Mutex{},
	}

	ctx := context.Background()

	resp, err := s.DeletePost(ctx, &pb.DeletePostRequest{PostId: "POST_1"})
	assert.NoError(t, err, "Deleting an existing post should not result in an error")
	assert.True(t, resp.Success, "Response success should be true after deleting an existing post")

	_, exists := s.Posts["POST_1"]
	assert.False(t, exists, "Post should not exist in the map after being deleted")

	_, err = s.DeletePost(ctx, &pb.DeletePostRequest{PostId: "POST_999"})
	assert.Error(t, err, "Deleting a non-existing post should result in an error")
}
