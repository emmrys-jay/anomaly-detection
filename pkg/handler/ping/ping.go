package ping

import (
	"encoding/csv"
	"net/http"
)

func PingGet(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	res.Header().Add("Content-Type", "text/csv")
	res.Write([]byte("GET OK"))
}

func PingPost(res http.ResponseWriter, req *http.Request) {
	results, err := csv.NewReader(req.Body).ReadAll()
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Header().Add("Content-Type", "text/csv")
		res.Write([]byte("Bad request," + err.Error()))
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
}
