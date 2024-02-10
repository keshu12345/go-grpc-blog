package main

import (
	"flag"
	"fmt"
	"net"

	"github.com/keshu12345/blog/blogPostServer/serverService"
	"github.com/keshu12345/blog/config"
	"github.com/keshu12345/blog/pb"

	logger "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var configDirPath = flag.String("config", "", "path for config dir")

func main() {

	flag.Parse()
	config := config.GetConfig(*configDirPath)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Server.Port))
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	server := &serverService.Server{Posts: make(map[string]*pb.Post)}
	pb.RegisterBlogServiceServer(s, server)

	logger.Info("server listening at", lis.Addr())
	if err := s.Serve(lis); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}
}
