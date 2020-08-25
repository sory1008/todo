package main

import (
	"context"
	"protos/todo"

	//context golang.org/x/net/context

	//"dataportal/protos/todo"

	"fmt"
	"log"
	"net"

	//todo "../todo"

	//"todo"

	grpc "google.golang.org/grpc"
)

type taskServer struct {
}

func (s taskServer) List(ctx context.Context, void *todo.Void) (*todo.TaskList, error) {
	return nil, fmt.Errorf("not implemented")
}

func main() {
	srv := grpc.NewServer()
	var tasks taskServer

	todo.RegisterTasksServer(srv, tasks)

	l, err := net.Listen("tcp", "8888")
	if err != nil {
		log.Fatalf("could not listen to :8888 - %v", err)
	}
	log.Fatal(srv.Serve(l))

}
