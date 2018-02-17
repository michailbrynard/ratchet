package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/testdata"

	//"github.com/golang/protobuf/proto"

	pb "github.com/odinsplasmarifle/ratchet/ratchet"
)

var (
	tls      = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile = flag.String("cert_file", "", "The TLS cert file")
	keyFile  = flag.String("key_file", "", "The TLS key file")
	port     = flag.Int("port", 10000, "The server port")
)

type server struct{}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	if *tls {
		if *certFile == "" {
			*certFile = testdata.Path("server1.pem")
		}
		if *keyFile == "" {
			*keyFile = testdata.Path("server1.key")
		}
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterRatchetServer(grpcServer, &server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *server) CreateTransaction(ctx context.Context, tx *pb.CreateTransactionRequest) (*pb.Transaction, error) {
	return &pb.Transaction{Type: "debit", Amount: 1}, nil
}

func (s *server) ListTransactions(tx *pb.ListTransactionsRequest, stream pb.Ratchet_ListTransactionsServer) error {
	for i := 0; i < 10; i++ {
		tx := &pb.Transaction{Type: "debit", Amount: 1}

		if err := stream.Send(tx); err != nil {
			return err
		}
	}

	return nil
}

func (s *server) GetTransaction(ctx context.Context, in *pb.GetTransactionRequest) (*pb.Transaction, error) {
	return &pb.Transaction{Type: "debit", Amount: 1}, nil
}

func (s *server) UpdateTransaction(ctx context.Context, in *pb.UpdateTransactionRequest) (*pb.Transaction, error) {
	return &pb.Transaction{Type: "debit", Amount: 1}, nil
}
