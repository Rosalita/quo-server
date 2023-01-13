package main

import (
	"context"
	"log"

	pb "github.com/Rosalita/quo-proto/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type server struct {
	pb.UnimplementedQuoteServiceServer
}

func newServer() server {
	return server{}
}

func (s *server) CreateQuote(ctx context.Context, req *pb.CreateQuoteRequest) (*pb.CreateQuoteResponse, error) {
	log.Printf("Received request %s", req.String())

	now := timestamppb.Now()

	return &pb.CreateQuoteResponse{
		QuoteName: req.QuoteName,
		QuoteText: req.QuoteText,
		CreatedAt: now,
	}, nil
}
