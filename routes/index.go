package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oswee/client/models"
	"github.com/oswee/client/utils"
	"github.com/oswee/proto"
)

func indexHandler(r *mux.Router) {
	r.HandleFunc("/", indexGetHandler).Methods("GET")
}

func indexGetHandler(w http.ResponseWriter, r *http.Request) {

	films, err := models.ListFilms()
	if err != nil {
		log.Fatalf("Error while calling ListFilms RPC: %v", err)
	}

	utils.ExecuteTemplate(w, "index.html", struct {
		Films []*proto.Film
	}{
		Films: films.Films,
	})
}
