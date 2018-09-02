package models

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func gLoc() (cc *grpc.ClientConn) {

	// Create the client TLS credentials
	creds, err := credentials.NewClientTLSFromFile("cert/server.crt", "")
	if err != nil {
		log.Fatalf("could not load tls cert: %s", err)
	}
	// Initiate a connection with the server
	cc, err = grpc.Dial("localhost:8080", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	return
}

func routeService() (cc *grpc.ClientConn) {

	// Create the client TLS credentials
	creds, err := credentials.NewClientTLSFromFile("cert/server.crt", "")
	if err != nil {
		log.Fatalf("could not load tls cert: %s", err)
	}
	// Initiate a connection with the server
	cc, err = grpc.Dial("localhost:8082", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	return
}
