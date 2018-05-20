package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tkanos/gonfig"

	"github.com/lrsmith/golang-fibonacci/handlers"
	"github.com/lrsmith/golang-fibonacci/middleware"
)

// Configuration Struct for storing configuration information
type Configuration struct {
	Port    int
	SSLKey  string
	SSLCert string
}

func getConfigs() Configuration {

	configuration := Configuration{}
	err := gonfig.GetConf("./config/app.config.json", &configuration)
	if err != nil {
		log.Fatal("Failed to read configuration file")
	}

	return configuration
}

func main() {

	log.Println("Reading configs")
	appConfigs := getConfigs()

	amw := middleware.AuthenticationMiddleware{}
	log.Println("Populating authentication tables")
	amw.Populate()

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/status", handlers.Status)
	router.HandleFunc("/v1/fibseq", handlers.FibSeq)

	router.Use(amw.AuthenticationMiddleware)
	router.Use(middleware.LoggingMiddleware)

	log.Println("Starting golang-fibonacci")

	log.Fatal(http.ListenAndServeTLS(":"+strconv.Itoa(appConfigs.Port), appConfigs.SSLCert, appConfigs.SSLKey, router))

}
