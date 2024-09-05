package sensors

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
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

	// run in background
	go mongodb.CreateSensorDataEntry(data)

	return nil
}

func LabelData() error {
	// file, err := os.OpenFile("data.csv", os.O_APPEND|os.O_RDONLY|os.O_WRONLY, 0777)
	// if err != nil {
	// 	log.Println("Could not open file, error - ", err.Error())
	// 	return err
	// }

	contents, err := ioutil.ReadFile("data.csv")
	if err != nil {
		log.Println("Error getting contents, error - ", err.Error())
		return err
	}

	reader := strings.NewReader(string(contents))
	
	records, err := csv.NewReader(reader).ReadAll()
	if err != nil {
		log.Println("Error getting contents, error - ", err.Error())
		return err
	}

	filters := make([]model.LabelFilter, 0, len(records))

	for _, row := range records {
		var labelFilter model.LabelFilter

		for idx, col := range row {
			if idx == 0 {
				t, err := time.Parse("2006-01-02 15:04:05", col)
				if err != nil {
					log.Fatalln("Could not parse time, error - ", err.Error())
					return err
				}

				labelFilter.StartTime = t.Add(-1 * time.Hour)
			}

			if idx == 1 {
				labelFilter.Anomaly = col
			}

			if idx == 2 {
				t, err := time.Parse("2006-01-02 15:04:05", col)
				if err != nil {
					log.Fatalln("Could not parse time, error - ", err.Error())
					return err
				}

				labelFilter.EndTime = t.Add(-1 * time.Hour)
			}
		}

		filters = append(filters, labelFilter)
	}

	err = mongodb.LabelData(filters)
	if err != nil {
		log.Fatalln("Could not label data, error - ", err.Error())
		return err
	}

	return nil
}

func DeleteInvalidData() error {
	err := mongodb.DeleteInvalidData()
	return err
}

func MoveData() error {
	err := mongodb.MoveData()
	return err
}

func LabelNoneData() error {
	err := mongodb.LabelNoneData()
	return err
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
