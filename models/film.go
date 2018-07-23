package models

import (
	"context"
	"log"

	"github.com/oswee/proto"
)

// ListFilms retrieve list of all films
func ListFilms() (*proto.ListFilmsResponse, error) {
	cc := gLoc()
	defer cc.Close()
	// We can now create stubs that wrap cc:
	c := proto.NewStarfriendsClient(cc)
	req := &proto.ListFilmsRequest{}
	res, err := c.ListFilms(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling ListFilms RPC: %v", err)
	}
	return res, nil
}
