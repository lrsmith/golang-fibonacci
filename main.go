package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	statsd "github.com/smira/go-statsd"
	"github.com/tkanos/gonfig"

	"github.com/lrsmith/golang-fibonacci/handlers"
	"github.com/lrsmith/golang-fibonacci/middleware"
)

// Configuration Struct for storing configuration information
type Configuration struct {
	Port         int
	SSLKey       string
	SSLCert      string
	StatsdEnable bool
	StatsdPrefix string
	StatsdServer string
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

	// Read in application configuration
	log.Println("Reading configs")
	appConfigs := getConfigs()

	// Initialze Authentication
	amw := middleware.AuthenticationMiddleware{}
	log.Println("Populating authentication tables")
	amw.Populate()

	// Initialize Statsd client.
	if appConfigs.StatsdEnable {
		log.Println("Enabling and configuring statsd client")
		middleware.StatsdClient = statsd.NewClient(appConfigs.StatsdServer,
			statsd.MaxPacketSize(1400),
			statsd.MetricPrefix(appConfigs.StatsdPrefix))
	}

	// Setup URL router
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/status", handlers.Status)
	router.HandleFunc("/v1/fibseq", handlers.FibSeq)

	// Enable middleware
	router.Use(amw.AuthenticationMiddleware)
	router.Use(middleware.LoggingMiddleware)
	if appConfigs.StatsdEnable {
		router.Use(middleware.StatsdMiddleware)
	}

	// Start service
	log.Println("Starting golang-fibonacci")
	log.Fatal(http.ListenAndServeTLS(":"+strconv.Itoa(appConfigs.Port), appConfigs.SSLCert, appConfigs.SSLKey, router))

	if appConfigs.StatsdEnable {
		middleware.StatsdClient.Close()
	}
}
