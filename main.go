package main

import (
	"log"
	"net"

	pb "github.com/odinsplasmarifle/ratchet/ratchet"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (s *server) Signal(ctx context.Context, in *pb.SignalRequest) (*pb.SignalReply, error) {
	return &pb.SignalReply{Key: "123"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterNodeServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
