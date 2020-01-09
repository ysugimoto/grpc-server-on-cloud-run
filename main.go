package main

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/ysugimoto/grpc-server-on-cloud-run/spec"
	"google.golang.org/grpc"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context, req *spec.HelloRequest) (*spec.HelloReply, error) {
	name := req.GetName()
	return &spec.HelloReply{
		Greeting: fmt.Sprintf("Hello, %s!", name),
	}, nil
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "50001"
	}
	conn, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	server := grpc.NewServer()
	spec.RegisterGreeterServer(server, &Server{})
	server.Serve(conn)
}
