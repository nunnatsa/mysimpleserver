package main

import (
	"log"
	"net/http"
	"os"

	"mysimpleserver/mux"
)

func main() {
	log.Println("starting server on port 8080")
	if err := http.ListenAndServe(":8080", mux.GetMux()); err != nil {
		log.Println("something went wrong while starting the server", err.Error())
		os.Exit(1)
	}
}
