package main

import (
	"context"
	"encoding/hex"
	"log"
	"time"

	pb "github.com/gotinker/wiresafe"
	"google.golang.org/grpc"
)

const (
	address = "localhost:65145"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("Failed to connect to %v", address)
	}
	defer conn.Close()
	c := pb.NewWireSafeClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*300)
	defer cancel()

	var safekey = pb.KeySpec{}
	safekey.Safename = "platform"
	safekey.Keyname = "account"

	r, err := c.InitKey(ctx, &safekey)

	if err != nil {
		log.Fatalf("Could not initialize KeySpec: %v\n", err)
	}

	workingkey := r.Workingkey

	decoded, err := hex.DecodeString(workingkey)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Working key is of type %T of length %v", decoded, len(decoded))
	log.Printf("Hex encoded generated working key is: %s", workingkey)

}
