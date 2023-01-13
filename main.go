package main

import (
	"fmt"
	"net"
	"os"

	"google.golang.org/grpc"

	pb "github.com/Rosalita/quo-proto/proto"
)

func main(){
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}


func run() error {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()

	s := newServer()

	pb.RegisterQuoteServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(listener); err != nil {
		return err
	}

	return nil
}