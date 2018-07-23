package models

import (
	"log"

	"google.golang.org/grpc"
)

func gLoc() (cc *grpc.ClientConn) {
	// First we create the connection:
	cc, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	return
}
