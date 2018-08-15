package routes

import (
	"github.com/gorilla/mux"
)

//NewRouter is main routing entry point
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	//There we are calling handlers from package routes
	fileServer(r)   // Fileserver to serve static files
	indexHandler(r) // Root level handler
	dmsRouter(r)    // Delivery Management Software

	return r
}
