package models

import (
	"context"
	"log"

	"github.com/oswee/proto"
)

func ListFilms(c proto.StarfriendsClient, req proto.ListFilmsRequest) {

	res, err := c.ListFilms(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling ListFilms RPC: %v", err)
	}
}
