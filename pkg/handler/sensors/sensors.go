package sensors

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"time"

	"github.com/emmrys-jay/anomaly-detection-api/service/sensors"
)

func logRequestData(code, timeElapsed time.Duration, timeForReceivedRequest time.Time, req *http.Request) {
	fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\n", 
	timeForReceivedRequest.Format(time.RFC3339),
	code, 
	timeElapsed.Seconds() * -1,
	req.RemoteAddr,
	req.Method,
	req.URL.Path,
)
}

func LogSensorData(res http.ResponseWriter, req *http.Request) {
	now := time.Now()

	records, err := csv.NewReader(req.Body).ReadAll()
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Header().Add("Content-Type", "text/csv")
		res.Write([]byte("Bad request," + err.Error()))
		logRequestData(http.StatusBadRequest, time.Until(now), now, req)
		return
	}

	err = sensors.LogSensorsData(records)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Header().Add("Content-Type", "text/csv")
		res.Write([]byte("Internal Server Error," + err.Error()))
		logRequestData(http.StatusInternalServerError, time.Until(now), now, req)
		return
	}

	res.WriteHeader(http.StatusCreated)
	res.Header().Add("Content-Type", "text/csv")
	res.Write([]byte("Success,Added Records"))
	logRequestData(http.StatusCreated, time.Until(now), now, req)
}
