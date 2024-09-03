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
	"Speed",
	"Vibration Detected",
	"Temp",
}

var DatabaseName = "anomaly_detection"
var CollectionName = "sensor_data"


func IsHeader(row []string) bool {
	for idx, val := range row {
		if !strings.EqualFold(val, csvHeaders[idx]) {
			return false
		}
	}

	return true
}

type SensorsData struct {
	AccX              float64   `bson:"Accel_X"`
	AccY              float64   `bson:"Accel_Y"`
	AccZ              float64   `bson:"Accel_Z"`
	GyrX              float64   `bson:"Gyro_X"`
	GyrY              float64   `bson:"Gyro_Y"`
	GyrZ              float64   `bson:"Gyro_Z"`
	Latitude          float64   `bson:"Latitude"`
	Longitude         float64   `bson:"Longitude"`
	DateTime          time.Time `bson:"Time"`
	Speed             float64   `bson:"Speed"`
	VibrationDetected int8      `bson:"Vibration"`
	Temperature       float64   `bson:"Temperature"`
	CreatedAt time.Time 		`bson:"CreatedAt"`
}
