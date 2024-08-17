package sensors

import (
	"fmt"
	"log"
	"strconv"

	"github.com/emmrys-jay/anomaly-detection-api/internal/model"
)

func LogSensorsData(records [][]string) error {
	var (
		data = make([]model.SensorsData, 0, len(records))
	)

	for _, row := range records {

		var datum model.SensorsData

		for idxIn, column := range row {
			assignStructValue(idxIn, &datum, column)
		}

		log.Printf("Values received are: %+v\n", datum)
		data = append(data, datum)
	}

	printStruct(data)
	// Save to database

	return nil
}

func assignStructValue(idx int, target *model.SensorsData, value string) {
	switch idx {

	case 0:
		res, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return
		}
		target.AccX = res

	case 1:
		res, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return
		}
		target.AccY = res

	case 2:
		res, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return
		}
		target.AccZ = res

	case 3:
		res, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return
		}
		target.GyrX = res

	case 4:
		res, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return
		}
		target.GyrY = res

	case 5:
		res, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return
		}
		target.GyrZ = res

	case 6:
		res, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return
		}
		target.Latitude = res

	case 7:
		res, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return
		}
		target.Longitude = res

	case 8:
		res, err := strconv.ParseInt(value, 10, 8)
		if err != nil {
			return
		}
		target.VibrationDetected = int8(res)
	}
}

func printStruct(value []model.SensorsData) {
	for _, val := range value {
		fmt.Printf("%v,%v,%v,%v,%v,%v,%v,%v,%v\n", val.AccX, val.AccY, val.AccZ,
			val.GyrX, val.GyrY, val.GyrZ, val.Longitude, val.Latitude, val.VibrationDetected)
	}
}
