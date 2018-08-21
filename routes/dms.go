package routes

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/oswee/client/models"
)

func dmsRouter(r *mux.Router) {
	r.HandleFunc("/", createDeliveryOrderPostHandler).Methods("POST")
	r.HandleFunc("/delete", deleteDeliveryOrderPostHandler).Methods("POST")
	r.HandleFunc("/update", updateDeliveryOrderPostHandler).Methods("POST")
	r.HandleFunc("/geocode", googleGeoCode).Methods("POST")
	r.HandleFunc("/batchgeocode", batchGeoCode).Methods("POST")
}

func createDeliveryOrderPostHandler(w http.ResponseWriter, r *http.Request) {

	models.CreatePageView(r)

	err := r.ParseForm()
	if err != nil {
		log.Println(err.Error)
	}
	ref := r.FormValue("reference")
	da := r.FormValue("address")
	dz := r.FormValue("zip")
	lat, _ := strconv.ParseFloat(r.FormValue("lat"), 64)
	lng, _ := strconv.ParseFloat(r.FormValue("lng"), 64)
	tw, _ := strconv.ParseFloat(r.FormValue("weight"), 64)
	rs, _ := strconv.ParseInt(r.FormValue("sequence"), 10, 64)

	_, err = models.CreateDeliveryOrder(ref, da, dz, lat, lng, tw, rs)
	if err != nil {
		log.Println(err.Error)
	}

	http.Redirect(w, r, "/", 302)
}

func updateDeliveryOrderPostHandler(w http.ResponseWriter, r *http.Request) {

	models.CreatePageView(r)

	err := r.ParseForm()
	if err != nil {
		log.Println(err.Error)
	}
	id, _ := strconv.ParseInt(r.FormValue("order_id"), 10, 64)
	ref := r.FormValue("reference")
	da := r.FormValue("address")
	dz := r.FormValue("zip")
	lat, _ := strconv.ParseFloat(r.FormValue("lat"), 64)
	lng, _ := strconv.ParseFloat(r.FormValue("lng"), 64)
	tw, _ := strconv.ParseFloat(r.FormValue("weight"), 64)
	rs, _ := strconv.ParseInt(r.FormValue("sequence"), 10, 64)

	_, err = models.UpdateDeliveryOrder(ref, da, dz, lat, lng, tw, rs, id)
	if err != nil {
		log.Println(err.Error)
	}

	http.Redirect(w, r, "/", 302)
}

// deleteDeliveryOrderPostHandler ...
func deleteDeliveryOrderPostHandler(w http.ResponseWriter, r *http.Request) {

	models.CreatePageView(r)

	err := r.ParseForm()
	if err != nil {
		log.Println(err.Error)
	}
	id, _ := strconv.ParseInt(r.FormValue("id"), 10, 64)

	_, err = models.DeleteDeliveryOrder(id)
	if err != nil {
		log.Println(err.Error)
	}

	http.Redirect(w, r, "/", 302)
}

//models.GeoCode("Brīvības iela 55, Rīga")

func googleGeoCode(w http.ResponseWriter, r *http.Request) {

	models.CreatePageView(r)

	err := r.ParseForm()
	if err != nil {
		log.Println(err.Error)
	}
	id, _ := strconv.ParseInt(r.FormValue("order_id"), 10, 64)
	stakeholder, _ := strconv.ParseInt(r.FormValue("1"), 10, 64)
	da := r.FormValue("address")

	_, err = models.GeoCodeDeliveryOrder(id, stakeholder, da)
	if err != nil {
		log.Println(err.Error)
	}

	http.Redirect(w, r, "/", 302)
}

func batchGeoCode(w http.ResponseWriter, r *http.Request) {
	models.CreatePageView(r)
	err := r.ParseForm()
	if err != nil {
		log.Println(err.Error)
	}
	sID, _ := strconv.ParseInt(r.FormValue("stakeholder_id"), 10, 64)
	models.BatchGeocode(sID)
	http.Redirect(w, r, "/", 302)
}
