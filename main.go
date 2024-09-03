package main

import (
	"log"
	"net/http"

	"github.com/emmrys-jay/anomaly-detection-api/internal/config"
	mongodb "github.com/emmrys-jay/anomaly-detection-api/pkg/repository/mongo"
	"github.com/emmrys-jay/anomaly-detection-api/pkg/router"
)

func init() {
	// setup mongoDB
	mongodb.ConnectDB()
}

func main() {
	app := router.Route()

	if config.PORT == "" {
		config.PORT = "8080"
	}

	log.Printf("Server is starting at 127.0.0.1:%s \n", config.PORT)
	http.ListenAndServe(":"+config.PORT, app)
}
