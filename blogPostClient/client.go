package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/keshu12345/blog/blogPostClient/clientService"
	"github.com/keshu12345/blog/config"
	"github.com/keshu12345/blog/pb"
	logger "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var configDirPath = flag.String("config", "", "path for config dir")

func main() {

	flag.Parse()

	config := config.GetConfig(*configDirPath)
	conn, err := grpc.Dial(fmt.Sprintf("%v", config.Server.Host)+fmt.Sprintf(":%d", config.Server.Port),
		grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		logger.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewBlogServiceClient(conn)
	getResponseFromServer(client)
}

func getResponseFromServer(client pb.BlogServiceClient) {

	clietService := clientService.ClientService{}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//post created
	newPost, err := clietService.CreatePost(ctx, client)
	if err != nil {
		logger.Errorf("Unable to create newPost %v", err)
	}
	logger.Info("Post Created :", newPost)

	// read post
	getPost, err := clietService.ReadPost(ctx, client, newPost)
	if err != nil {
		logger.Fatalf("Could not read post: %v", err)
	}
	logger.Info("Post read: ", getPost)

	// Update Post
	updatedPost, err := clietService.UpdatePost(ctx, client, newPost)
	if err != nil {
		logger.Fatalf("Could not update post: %v", err)
	}
	logger.Info("Post updated: ", updatedPost)

	// Delete Post
	delResponse, err := clietService.DeletePost(ctx, client, newPost)
	if err != nil {
		logger.Fatalf("Could not delete post: %v", err)
	}
	logger.Info("Post deleted: ", delResponse.Success)
}
