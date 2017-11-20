package main

import (
	"log"
	"net"

	pb "github.com/odinsplasmarifle/ratchet/ratchet"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Node struct {
	config   *Config
	messages chan Message
}

type Config struct {
	Port string
}

type Message struct {
	Type string
}

func (node *Node) Run() {
	node.messages = make(chan Message)

	lis, err := net.Listen("tcp", node.config.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	n := grpc.NewServer()
	pb.RegisterNodeServer(n, node)
	reflection.Register(n)
	if err := n.Serve(lis); err != nil {
		log.Fatalf("failed to run node: %v", err)
	}
}

func (node *Node) Signal(ctx context.Context, in *pb.SignalRequest) (*pb.SignalReply, error) {
	return &pb.SignalReply{Key: "hello"}, nil
}
