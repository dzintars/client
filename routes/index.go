package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oswee/client/models"
	"github.com/oswee/client/utils"
)

func indexHandler(r *mux.Router) {
	r.HandleFunc("/", indexGetHandler).Methods("GET")
}

func indexGetHandler(w http.ResponseWriter, r *http.Request) {

	data, err := models.ListFilms()
	if err != nil {
		log.Fatalf("Error while calling ListFilms RPC: %v", err)
	}

	utils.ExecuteTemplate(w, "index.html", data)
}
