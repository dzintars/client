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
}

func createDeliveryOrderPostHandler(w http.ResponseWriter, r *http.Request) {
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
