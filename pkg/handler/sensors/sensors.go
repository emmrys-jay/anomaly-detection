package sensors

import (
	"encoding/csv"
	"net/http"

	"github.com/emmrys-jay/anomaly-detection-api/service/sensors"
)

func LogSensorData(res http.ResponseWriter, req *http.Request) {
	records, err := csv.NewReader(req.Body).ReadAll()
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Header().Add("Content-Type", "text/csv")
		res.Write([]byte("Bad request," + err.Error()))
		return
	}

	err = sensors.LogSensorsData(records)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Header().Add("Content-Type", "text/csv")
		res.Write([]byte("Internal Server Error," + err.Error()))
		return
	}

	res.WriteHeader(http.StatusCreated)
	res.Header().Add("Content-Type", "text/csv")
	res.Write([]byte("Success,Added Records"))
}
