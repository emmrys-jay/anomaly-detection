package router

import (
	"net/http"

	"github.com/emmrys-jay/anomaly-detection-api/pkg/handler/ping"
	"github.com/emmrys-jay/anomaly-detection-api/pkg/handler/sensors"
)

func Route() *http.ServeMux {

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", ping.PingGet)
	mux.HandleFunc("POST /", ping.PingPost)

	mux.HandleFunc("POST /log", sensors.LogSensorData)

	return mux
}
