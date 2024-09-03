package ping

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"time"
)

func logRequestData(code int, timeElapsed time.Duration, timeForReceivedRequest time.Time, req *http.Request) {
	fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\n", 
	timeForReceivedRequest.Format(time.RFC3339),
	code, 
	fmt.Sprint(timeElapsed.Seconds() * -1) + "s",
	req.RemoteAddr,
	req.Method,
	req.URL.Path,
)
}

func PingGet(res http.ResponseWriter, req *http.Request) {
	now := time.Now()
	res.WriteHeader(http.StatusOK)
	res.Header().Add("Content-Type", "text/csv")
	res.Write([]byte("GET OK"))
	logRequestData(http.StatusOK, time.Until(now), now, req)
}

func PingPost(res http.ResponseWriter, req *http.Request) {
	now := time.Now()

	results, err := csv.NewReader(req.Body).ReadAll()
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Header().Add("Content-Type", "text/csv")
		res.Write([]byte("Bad request," + err.Error()))
		logRequestData(http.StatusBadRequest, time.Until(now), now, req)
		return
	}

	var firstRecord string
	if len(results) >= 1 {
		if len(results[0]) >= 1 {
			firstRecord = results[0][0]
		}
	}

	res.WriteHeader(http.StatusOK)
	res.Header().Add("Content-Type", "text/csv")
	res.Write([]byte("POST OK - " + firstRecord))
	logRequestData(http.StatusOK, time.Until(now), now, req)
}
