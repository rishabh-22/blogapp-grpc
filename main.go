package main

import (
	"context"
	"github.com/rishabh-22/blogapp-grpc/blogpost"
	"google.golang.org/grpc"
	"log"
	"net"
)

type myBlogPostServer struct {
	blogpost.UnimplementedBlogServiceServer
}

func (s myBlogPostServer) Create(ctx context.Context, req *blogpost.CreatePost) (*blogpost.Response, error) {
	return &blogpost.Response{
		PostID:          100,
		Title:           "test",
		Content:         "test",
		Author:          "test",
		PublicationDate: "test",
		Tags:            "test",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}
	serverRegistrar := grpc.NewServer()
	service := &myBlogPostServer{}
	blogpost.RegisterBlogServiceServer(serverRegistrar, service)
	err = serverRegistrar.Serve(lis)
	if err != nil {
		log.Fatalf("impossible to serve: %s", err)
	}
}
