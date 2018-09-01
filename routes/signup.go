package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oswee/client/models"
	"github.com/oswee/client/utils"
	app "github.com/oswee/stubs/app/v1"
)

func signupRouter(r *mux.Router) {
	r.HandleFunc("/signup", signupGetHandler).Methods("GET")
	r.HandleFunc("/signup", signupPostHandler).Methods("POST")
}

func signupGetHandler(w http.ResponseWriter, r *http.Request) {
	// ToDO: Log handler execution time
	_, err := models.CreatePageView(r)
	if err != nil {
		log.Fatalf("Error while calling GetApplication model: %v", err)
	}

	application, err := models.GetApplication(12)
	if err != nil {
		log.Fatalf("Error while calling GetApplication model: %v", err)
	}

	utils.ExecuteTemplate(w, "signup.html", struct {
		Application *app.Application
	}{
		Application: application.Application,
	})
}

func signupPostHandler(w http.ResponseWriter, r *http.Request) {

	models.CreatePageView(r)

	err := r.ParseForm()
	if err != nil {
		log.Println(err.Error)
	}
	fname := r.FormValue("fname")
	lname := r.FormValue("lname")
	email := r.FormValue("email")
	uname := r.FormValue("uname")
	password := r.FormValue("password")

	_, err = models.CreateSignup(fname, lname, email, uname, password)
	if err != nil {
		log.Println(err.Error)
	}

	http.Redirect(w, r, "/signin", 302)
}
