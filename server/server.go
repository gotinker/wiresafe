package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"log"
	"net"

	pb "github.com/gotinker/wiresafe"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	port = ":65145"
	salt = "lumi2022ngbd" //TODO initialize this from some safe
)

type WireSafeServer struct {
	pb.UnimplementedWireSafeServer
}

func (s *WireSafeServer) InitKey(ctx context.Context, in *pb.KeySpec) (*pb.InitReply, error) {

	log.Printf("Received %v", in)

	hash := sha256.Sum256([]byte(in.Safename + salt + in.GetSafename()))

	return &pb.InitReply{Workingkey: hex.EncodeToString(hash[:])}, nil

}

func (s *WireSafeServer) Encrypt(ctx context.Context, in *pb.EncryptReq) (*pb.EncryptResp, error) {

	log.Printf("Received %v", in)

	return nil, status.Errorf(codes.Unimplemented, "method Encrypt not implemented")
}
func (s *WireSafeServer) Decrypt(ctx context.Context, in *pb.DecryptReq) (*pb.DecryptResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Decrypt not implemented")
}

func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {

		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterWireSafeServer(s, &WireSafeServer{})
	log.Printf("Listening on %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
