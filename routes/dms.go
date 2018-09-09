package routes

import (
	"github.com/gorilla/mux"
	h "github.com/oswee/client/handlers"
)

func dmsRouter(r *mux.Router) {
	r.HandleFunc("/", h.CreateDeliveryOrderPostHandler).Methods("POST")
	r.HandleFunc("/delete", h.DeleteDeliveryOrderPostHandler).Methods("POST")
	r.HandleFunc("/update", h.UpdateDeliveryOrderPostHandler).Methods("POST")
	r.HandleFunc("/geocode", h.GoogleGeoCode).Methods("POST")
	r.HandleFunc("/batchgeocode", h.BatchGeoCode).Methods("POST")
	r.HandleFunc("/routes", h.CreateRoutePost).Methods("POST")
}
