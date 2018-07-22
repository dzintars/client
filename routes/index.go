package routes

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oswee/client/utils"
	"github.com/oswee/proto"
	"google.golang.org/grpc"
)

func indexHandler(r *mux.Router) {
	r.HandleFunc("/", indexGetHandler).Methods("GET")
}

func indexGetHandler(w http.ResponseWriter, r *http.Request) {

	cc, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()
	// We can now create stubs that wrap conn:
	c := proto.NewStarfriendsClient(cc)

	data, err := doUnary(c)
	if err != nil {
		log.Fatalf("Error while calling ListFilms RPC: %v", err)
	}
	fmt.Println(data)

	utils.ExecuteTemplate(w, "index.html", data)
}

func doUnary(c proto.StarfriendsClient) (*proto.ListFilmsResponse, error) {
	req := &proto.ListFilmsRequest{}
	res, err := c.ListFilms(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling ListFilms RPC: %v", err)
	}
	return res, nil
}
