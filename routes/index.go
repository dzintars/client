package routes

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/oswee/client/models"
	"github.com/oswee/client/utils"
	pb "github.com/oswee/proto/customer/go"
	"google.golang.org/grpc"
)

func indexHandler(r *mux.Router) {
	r.HandleFunc("/favicon.ico", faviconHandler).Methods("GET")
	r.HandleFunc("/", indexGetHandler).Methods("GET")
	r.HandleFunc("/create", customerPostHandler).Methods("POST")
}

// Handlers

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "/static/img/favicon.png", nil)
}

func indexGetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header.Get("X-Forwarded-For"), time.Now())

	// customers, err := models.ListCustomers()
	// if err != nil {
	// 	return
	// }
	const (
		address = "localhost:50051"
	)

	// Set up a connection to the gRPC server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// Creates a new CustomerClient
	client := pb.NewCustomerClient(conn)

	// Filter with an empty Keyword
	filter := &pb.CustomerFilter{Keyword: ""}
	models.GetCustomers(client, filter)

	utils.ExecuteTemplate(w, "index.html", struct {
		Title string
		// Customers  []models.Customer
	}{
		Title: "Customers",
		// Customers: customers,
	})
}

func customerPostHandler(w http.ResponseWriter, r *http.Request) {
	const (
		address = "localhost:50051"
	)

	// Set up a connection to the gRPC server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// Creates a new CustomerClient
	client := pb.NewCustomerClient(conn)

	customer := &pb.CustomerRequest{
		Id:    101,
		Name:  "Shiju Varghese",
		Email: "shiju@xyz.com",
		Phone: "732-757-2923",
		Addresses: []*pb.CustomerRequest_Address{
			&pb.CustomerRequest_Address{
				Street:            "1 Mission Street",
				City:              "San Francisco",
				State:             "CA",
				Zip:               "94105",
				IsShippingAddress: false,
			},
			&pb.CustomerRequest_Address{
				Street:            "Greenfield",
				City:              "Kochi",
				State:             "KL",
				Zip:               "68356",
				IsShippingAddress: true,
			},
		},
	}
	// Create a new customer
	models.CreateCustomer(client, customer)

	customer = &pb.CustomerRequest{
		Id:    102,
		Name:  "Irene Rose",
		Email: "irene@xyz.com",
		Phone: "732-757-2924",
		Addresses: []*pb.CustomerRequest_Address{
			&pb.CustomerRequest_Address{
				Street:            "1 Mission Street",
				City:              "San Francisco",
				State:             "CA",
				Zip:               "94105",
				IsShippingAddress: true,
			},
		},
	}

	// Create a new customer
	models.CreateCustomer(client, customer)

	utils.ExecuteTemplate(w, "create.html", struct {
		Title string
	}{
		Title: "Create",
	})
}
