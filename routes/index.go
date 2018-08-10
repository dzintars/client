package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oswee/client/models"
	"github.com/oswee/client/utils"
	app "github.com/oswee/proto/application/go"
)

func indexHandler(r *mux.Router) {
	r.HandleFunc("/", indexGetHandler).Methods("GET")
}

func indexGetHandler(w http.ResponseWriter, r *http.Request) {

	applications, err := models.ListApplications(100)
	if err != nil {
		log.Fatalf("Error while calling ListApplications model: %v", err)
	}

	application, err := models.GetApplication(10)
	if err != nil {
		log.Fatalf("Error while calling GetApplication model: %v", err)
	}

	utils.ExecuteTemplate(w, "index.html", struct {
		Applications []*app.Application
		Application  *app.Application
	}{
		Applications: applications.Applications,
		Application:  application.Application,
	})
}
