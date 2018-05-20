package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lrsmith/golang-fibonacci/handlers"
	"github.com/lrsmith/golang-fibonacci/middleware"
)

func main() {

	amw := middleware.AuthenticationMiddleware{}
	log.Println("Populating authentication tables")
	amw.Populate()

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/status", handlers.Status)
	router.HandleFunc("/v1/fibseq", handlers.FibSeq)

	router.Use(amw.AuthenticationMiddleware)
	router.Use(middleware.LoggingMiddleware)

	log.Println("Starting golang-fibonacci")

	log.Fatal(http.ListenAndServeTLS(":8443", "./config/server.crt", "./config/server.key", router))

}
