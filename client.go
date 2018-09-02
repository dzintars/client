package main

import (
	"log"
	"net/http"
	"time"

	"github.com/oswee/client/routes"
	"github.com/oswee/client/utils"
)

const (
	port = ":80"
)

func main() {
	utils.LoadTemplates("templates/*.html")
	//Load routes package
	r := routes.NewRouter()
	//Pass routes package to Handler
	http.Handle("/", r)

	// https://www.youtube.com/watch?v=wxkEQxvxs3w&t=769s

	srv := &http.Server{
		Addr:         port,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		//TLSConfig: tlsConfig,
		Handler: nil,
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Client filed to start: %v", err)
	}
}
