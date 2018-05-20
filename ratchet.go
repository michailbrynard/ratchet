package ratchet

import (
	"fmt"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/testdata"

	//"github.com/golang/protobuf/proto"

	pb "github.com/odinsplasmarifle/ratchet/proto"
)

type server struct {
	Tls      bool
	CertFile string
	KeyFile  string
	Port     int
}

func NewServer() *server {
	s := server{
		Tls:      false,
		CertFile: "",
		KeyFile:  "",
		Port:     8080,
	}
	return &s
}

func (s *server) Serve() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", s.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	if s.Tls {
		if s.CertFile == "" {
			s.CertFile = testdata.Path("server1.pem")
		}
		if s.KeyFile == "" {
			s.KeyFile = testdata.Path("server1.key")
		}
		creds, err := credentials.NewServerTLSFromFile(s.CertFile, s.KeyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterRatchetServer(grpcServer, &server{})

	log.Print("Start ratchet")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

// Ledgers

func (s *server) CreateLedger(ctx context.Context, tx *pb.CreateLedgerRequest) (*pb.Ledger, error) {
	return &pb.Ledger{Id: "1"}, nil
}

func (s *server) ListLedgers(tx *pb.ListLedgersRequest, stream pb.Ratchet_ListLedgersServer) error {
	for i := 0; i < 10; i++ {
		tx := &pb.Ledger{Id: "1"}

		if err := stream.Send(tx); err != nil {
			return err
		}
	}

	return nil
}

func (s *server) GetLedger(ctx context.Context, in *pb.GetLedgerRequest) (*pb.Ledger, error) {
	return &pb.Ledger{Id: "1"}, nil
}

func (s *server) UpdateLedger(ctx context.Context, in *pb.UpdateLedgerRequest) (*pb.Ledger, error) {
	return &pb.Ledger{Id: "1"}, nil
}

// Assets

func (s *server) CreateAsset(ctx context.Context, tx *pb.CreateAssetRequest) (*pb.Asset, error) {
	return &pb.Asset{Id: "1"}, nil
}

func (s *server) ListAssets(tx *pb.ListAssetsRequest, stream pb.Ratchet_ListAssetsServer) error {
	for i := 0; i < 10; i++ {
		tx := &pb.Asset{Id: "1"}

		if err := stream.Send(tx); err != nil {
			return err
		}
	}

	return nil
}

func (s *server) GetAsset(ctx context.Context, in *pb.GetAssetRequest) (*pb.Asset, error) {
	return &pb.Asset{Id: "1"}, nil
}

func (s *server) UpdateAsset(ctx context.Context, in *pb.UpdateAssetRequest) (*pb.Asset, error) {
	return &pb.Asset{Id: "1"}, nil
}

// // Accounts

func (s *server) CreateAccount(ctx context.Context, tx *pb.CreateAccountRequest) (*pb.Account, error) {
	return &pb.Account{Id: "1"}, nil
}

func (s *server) ListAccounts(tx *pb.ListAccountsRequest, stream pb.Ratchet_ListAccountsServer) error {
	for i := 0; i < 10; i++ {
		tx := &pb.Account{Id: "1"}

		if err := stream.Send(tx); err != nil {
			return err
		}
	}

	return nil
}

func (s *server) GetAccount(ctx context.Context, in *pb.GetAccountRequest) (*pb.Account, error) {
	return &pb.Account{Id: "1"}, nil
}

func (s *server) UpdateAccount(ctx context.Context, in *pb.UpdateAccountRequest) (*pb.Account, error) {
	return &pb.Account{Id: "1"}, nil
}

// Transactions

func (s *server) CreateTransaction(ctx context.Context, tx *pb.CreateTransactionRequest) (*pb.Transaction, error) {
	return &pb.Transaction{Id: "1"}, nil
}

func (s *server) ListTransactions(tx *pb.ListTransactionsRequest, stream pb.Ratchet_ListTransactionsServer) error {
	for i := 0; i < 10; i++ {
		tx := &pb.Transaction{Id: "1"}

		if err := stream.Send(tx); err != nil {
			return err
		}
	}

	return nil
}

func (s *server) GetTransaction(ctx context.Context, in *pb.GetTransactionRequest) (*pb.Transaction, error) {
	return &pb.Transaction{Id: "1"}, nil
}

func (s *server) UpdateTransaction(ctx context.Context, in *pb.UpdateTransactionRequest) (*pb.Transaction, error) {
	return &pb.Transaction{Id: "1"}, nil
}
