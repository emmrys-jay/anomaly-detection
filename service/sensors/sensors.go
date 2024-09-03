package sensors

import (
	"fmt"
	"strconv"
	"time"

	"github.com/emmrys-jay/anomaly-detection-api/internal/model"
	mongodb "github.com/emmrys-jay/anomaly-detection-api/pkg/repository/mongo"
)

func LogSensorsData(records [][]string) error {
	var (
		data = make([]model.SensorsData, 0, len(records))
	)

	for _, row := range records {

		var datum model.SensorsData

		if model.IsHeader(row) {
			continue
		}

		for idxIn, column := range row {
			assignStructValue(idxIn, &datum, column)
		}

		data = append(data, datum)
	}

	err := mongodb.CreateSensorDataEntry(data)
	if err != nil {
		return err
	}

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
		res, err := time.Parse("2006-01-02 15:04:05", value)
		if err != nil {
			return
		}
		target.DateTime = res

	case 9:
		res, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return
		}
		target.Speed = res
		
	case 10:
		res, err := strconv.ParseInt(value, 10, 8)
		if err != nil {
			return
		}
		target.VibrationDetected = int8(res)

	case 11:
		res, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return
		}
		target.Temperature = res

	default: 
		return
	}
}

func printStruct(value []model.SensorsData) {
	for _, val := range value {
		fmt.Printf("%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v\n", val.AccX, val.AccY, val.AccZ,
			val.GyrX, val.GyrY, val.GyrZ, val.Longitude, val.Latitude, val.DateTime, 
			val.Speed, val.VibrationDetected, val.Temperature)
	}
}
