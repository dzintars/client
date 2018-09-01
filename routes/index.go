package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oswee/client/models"
	"github.com/oswee/client/utils"
	app "github.com/oswee/stubs/app/v1"
	dms "github.com/oswee/stubs/dms/v1"
)

func indexHandler(r *mux.Router) {
	r.HandleFunc("/", indexGetHandler).Methods("GET")
	r.HandleFunc("/shipping", shippingGetHandler).Methods("GET")
	r.HandleFunc("/signin", signinGetHandler).Methods("GET")
	r.HandleFunc("/signup", signupGetHandler).Methods("GET")
}

// func setupResponse(w *http.ResponseWriter, req *http.Request) {
// 	(*w).Header().Set("Access-Control-Allow-Origin", "*")
// 	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
// 	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
// }

func indexGetHandler(w http.ResponseWriter, r *http.Request) {

	// setupResponse(&w, r)
	// if (*r).Method == "OPTIONS" {
	// 	return
	// }

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
		DeliveryOrders []*dms.DeliveryOrder
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
}

func signinGetHandler(w http.ResponseWriter, r *http.Request) {
	// ToDO: Log handler execution time
	_, err := models.CreatePageView(r)
	if err != nil {
		log.Fatalf("Error while calling GetApplication model: %v", err)
	}

	application, err := models.GetApplication(11)
	if err != nil {
		log.Fatalf("Error while calling GetApplication model: %v", err)
	}

	utils.ExecuteTemplate(w, "signin.html", struct {
		Application *app.Application
	}{
		Application: application.Application,
	})
}
