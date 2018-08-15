package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oswee/client/models"
	"github.com/oswee/client/utils"
	app "github.com/oswee/proto/application/go"
	"github.com/oswee/proto/shipping/go"
)

func indexHandler(r *mux.Router) {
	r.HandleFunc("/", indexGetHandler).Methods("GET")
	r.HandleFunc("/shipping", shippingGetHandler).Methods("GET")
}

func indexGetHandler(w http.ResponseWriter, r *http.Request) {

	_, err := models.CreatePageView(r)
	if err != nil {
		log.Fatalf("Error while calling GetApplication model: %v", err)
	}

	deliveryOrders, err := models.ListDeliveryOrders(1, 100)
	if err != nil {
		log.Fatalf("Error while calling GetApplication model: %v", err)
	}

	application, err := models.GetApplication(10)
	if err != nil {
		log.Fatalf("Error while calling GetApplication model: %v", err)
	}

	utils.ExecuteTemplate(w, "index.html", struct {
		DeliveryOrders []*shipping.DeliveryOrder
		Application    *app.Application
	}{
		DeliveryOrders: deliveryOrders.DeliveryOrders,
		Application:    application.Application,
	})
}

func shippingGetHandler(w http.ResponseWriter, r *http.Request) {

	_, err := models.CreatePageView(r)
	if err != nil {
		log.Fatalf("Error while calling GetApplication model: %v", err)
	}

	applications, err := models.ListApplications(100)
	if err != nil {
		log.Fatalf("Error while calling ListApplications model: %v", err)
	}

	application, err := models.GetApplication(19)
	if err != nil {
		log.Fatalf("Error while calling GetApplication model: %v", err)
	}

	utils.ExecuteTemplate(w, "shipping.html", struct {
		Applications []*app.Application
		Application  *app.Application
	}{
		Applications: applications.Applications,
		Application:  application.Application,
	})
	fmt.Println("Shipping delivered")
}
