package main

import (
	"net/http"

	"github.com/oswee/client/routes"
	"github.com/oswee/client/utils"
)

func main() {
	utils.LoadTemplates("templates/*.html")
	//Load routes package
	r := routes.NewRouter()
	//Pass routes package to Handler
	http.Handle("/", r)
	http.ListenAndServe(":8001", nil)
}
