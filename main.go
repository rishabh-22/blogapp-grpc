package main

import (
	"github.com/rishabh-22/blogapp-grpc/blogpost"
	"github.com/rishabh-22/blogapp-grpc/pkg/api"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {

	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}
	serverRegistrar := grpc.NewServer()
	service := &api.MyBlogPostServer{}
	blogpost.RegisterBlogServiceServer(serverRegistrar, service)
	err = serverRegistrar.Serve(lis)
	if err != nil {
		log.Fatalf("impossible to serve: %s", err)
	}
}
