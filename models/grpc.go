package models

import (
	"log"

	"google.golang.org/grpc"
)

func GLoc() (cc *grpc.ClientConn) {

	// First we create the connection:
	cc, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return cc
}
