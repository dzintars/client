package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oswee/client/models"
	"github.com/oswee/client/utils"
	"github.com/oswee/proto"
	app "github.com/oswee/proto/application/go"
)

func indexHandler(r *mux.Router) {
	r.HandleFunc("/", indexGetHandler).Methods("GET")
}

func indexGetHandler(w http.ResponseWriter, r *http.Request) {

	films, err := models.ListFilms()
	if err != nil {
		log.Fatalf("Error while calling ListFilms RPC: %v", err)
	}

	apps, err := models.ListApplications(100)
	if err != nil {
		log.Fatalf("Error while calling ListApplications model: %v", err)
	}

	abc, err := models.GetApplication(10)
	if err != nil {
		log.Fatalf("Error while calling GetApplication model: %v", err)
	}

	utils.ExecuteTemplate(w, "index.html", struct {
		Films []*proto.Film
		Apps  []*app.Application
		Appx  *app.Application
	}{
		Films: films.Films,
		Apps:  apps.Applications,
		Appx:  abc.Application,
	})
}
