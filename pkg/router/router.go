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
	// mux.HandleFunc("PATCH /label", sensors.LabelData)
	mux.HandleFunc("DELETE /delete", sensors.DeleteInvalidData)
	mux.HandleFunc("POST /move", sensors.MoveData)
	mux.HandleFunc("PATCH /label/none", sensors.LabelNoneData)

	return mux
}
