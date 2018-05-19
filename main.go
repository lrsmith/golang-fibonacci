package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lrsmith/golang-fibonacci/handlers"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/status", handlers.Status)

	log.Fatal(http.ListenAndServeTLS(":8443", "./config/server.crt", "./config/server.key", router))

}
