package model

var csvHeaders = []string{
	"AccX",
	"AccY",
	"AccZ",
	"GyrX",
	"GyrY",
	"GyrZ",
	"Latitude",
	"Longitude",
	"Vibration Detected",
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
	VibrationDetected int8
}
