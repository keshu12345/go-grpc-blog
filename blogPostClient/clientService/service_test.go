package clientService

import (
	"context"
	"testing"

	"github.com/keshu12345/blog/MockBlogServiceClient/mocks"
	"github.com/keshu12345/blog/pb"
	"github.com/stretchr/testify/assert"
	"github.com/test-go/testify/mock"
)

func TestCreatePost(t *testing.T) {

	mockClient := &mocks.BlogServiceClient{}
	mockPost := &pb.Post{
		Title:   "New Post",
		Content: "This is a new post.",
		Author:  "Author A",
		Tags:    []string{"tag1", "tag2"},
	}

	mockClient.On("CreatePost", mock.Anything, &pb.CreatePostRequest{
		Title:   "New Post",
		Content: "This is a new post.",
		Author:  "Author A",
		Tags:    []string{"tag1", "tag2"},
	}).Return(mockPost, nil)

	clientService := ClientService{}
	resp, err := clientService.CreatePost(context.Background(), mockClient)

	assert.Nil(t, err)
	assert.Equal(t, mockPost, resp)
	mockClient.AssertExpectations(t)
}

func TestClientService_CreatePost(t *testing.T) {
	mockClient := &mocks.BlogServiceClient{}
	expectedPost := &pb.Post{
		PostId:  "123",
		Title:   "Test Post",
		Content: "Test content",
		Author:  "Test Author",
		Tags:    []string{"test1", "test2"},
	}
	mockClient.On("ReadPost", mock.Anything, &pb.ReadPostRequest{PostId: "123"}).Return(expectedPost, nil)
	clientService := ClientService{}
	post, err := clientService.ReadPost(context.Background(), mockClient, &pb.Post{PostId: "123"})
	assert.NoError(t, err)
	assert.Equal(t, expectedPost, post)
	mockClient.AssertExpectations(t)
}

func TestClientService_ReadPost(t *testing.T) {
	mockClient := &mocks.BlogServiceClient{}
	expectedPost := &pb.Post{
		PostId:  "123",
		Title:   "Test Post",
		Content: "Test content",
		Author:  "Test Author",
		Tags:    []string{"test1", "test2"},
	}
	mockClient.On("ReadPost", mock.Anything, &pb.ReadPostRequest{PostId: "123"}).Return(expectedPost, nil)
	clientService := ClientService{}
	post, err := clientService.ReadPost(context.Background(), mockClient, &pb.Post{PostId: "123"})
	assert.NoError(t, err)
	assert.Equal(t, expectedPost, post)
	mockClient.AssertExpectations(t)
}

func TestClientService_UpdatePost(t *testing.T) {
	mockClient := &mocks.BlogServiceClient{}
	originalPost := &pb.Post{
		PostId:  "post123",
		Title:   "Original Post",
		Content: "Original content",
		Author:  "Author A",
		Tags:    []string{"tag1", "tag2"},
	}
	updatedPost := &pb.Post{
		PostId:  "post123",
		Title:   "Updated Post",
		Content: "This is an updated post content.",
		Author:  "Author B",
		Tags:    []string{"tag3", "tag4"},
	}
	mockClient.On("UpdatePost", mock.Anything, &pb.UpdatePostRequest{
		PostId:  originalPost.PostId,
		Title:   updatedPost.Title,
		Content: updatedPost.Content,
		Author:  updatedPost.Author,
		Tags:    updatedPost.Tags,
	}).Return(updatedPost, nil)

	clientService := ClientService{}
	resultPost, err := clientService.UpdatePost(context.Background(), mockClient, originalPost)

	assert.NoError(t, err)
	assert.Equal(t, updatedPost, resultPost)
	mockClient.AssertExpectations(t) // Verify that the expectations were met
}

func TestClientService_DeletePost(t *testing.T) {
	mockClient := &mocks.BlogServiceClient{}
	postToDelete := &pb.Post{PostId: "post123"}
	expectedResponse := &pb.DeletePostResponse{Success: true}

	// Setting up the expected call on the mock with the specific request
	mockClient.On("DeletePost", mock.Anything, &pb.DeletePostRequest{PostId: postToDelete.PostId}).Return(expectedResponse, nil)

	// Creating an instance of the ClientService
	clientService := ClientService{}

	// Calling the DeletePost method
	response, err := clientService.DeletePost(context.Background(), mockClient, postToDelete)

	// Assertions to verify the outcomes
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, response)
	mockClient.AssertExpectations(t) // Ensure the mock's expectations were met
}
