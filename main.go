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
	amw.Populate()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/status", handlers.Status)
	router.HandleFunc("/v1/fibseq", handlers.FibSeq)

	router.Use(amw.Middleware)

	log.Fatal(http.ListenAndServeTLS(":8443", "./config/server.crt", "./config/server.key", router))

}
