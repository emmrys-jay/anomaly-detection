package model

import (
	"strings"
	"time"
)

var csvHeaders = []string{
	"AcX",
	"AcY",
	"AcZ",
	"GyX",
	"GyY",
	"GyZ",
	"Latitude",
	"Longitude",
	"Time",
	"Vibration Detected",
	"Temp",
}

func IsHeader(row []string) bool {
	for idx, val := range row {
		if !strings.EqualFold(val, csvHeaders[idx]) {
			return false
		}
	}

	return true
}

type SensorsData struct {
	AccX              float64
	AccY              float64
	AccZ              float64
	GyrX              float64
	GyrY              float64
	GyrZ              float64
	Latitude          float64
	Longitude         float64
	DateTime time.Time
	VibrationDetected int8
	Temperature float64
}
