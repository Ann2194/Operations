package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	api "operationProject/pkg/api/github.com/example/path/gen"
	"operationProject/pkg/operations"
)

func main() {
	s := grpc.NewServer()
	srv := &operations.GRPCServer{}
	api.RegisterOperationServer(s, srv)

	l, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(fmt.Errorf("listen error: %w", err))
	}
	if err := s.Serve(l); err != nil {
		panic(fmt.Errorf("failed response: %w", err))
	}
}
