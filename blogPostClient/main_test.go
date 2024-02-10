package main

import (
	"testing"

	"github.com/keshu12345/blog/MockBlogServiceClient/mocks"
	"github.com/keshu12345/blog/pb"
	"github.com/test-go/testify/mock"
)

func TestGetResponseFromServer(t *testing.T) {
	mockClient := &mocks.BlogServiceClient{}

	mockPost := &pb.Post{
		Title:   "New Post",
		Content: "This is a new post.",
		Author:  "Author A",
		Tags:    []string{"tag1", "tag2"},
	}
	mockClient.On("CreatePost", mock.Anything, mock.Anything).Return(mockPost, nil)
	mockClient.On("ReadPost", mock.Anything, mock.Anything).Return(mockPost, nil)
	mockClient.On("UpdatePost", mock.Anything, mock.Anything).Return(mockPost, nil)
	mockClient.On("DeletePost", mock.Anything, mock.Anything).Return(&pb.DeletePostResponse{Success: true}, nil)
	getResponseFromServer(mockClient)
	mockClient.AssertExpectations(t)
}
