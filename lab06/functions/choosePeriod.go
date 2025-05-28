package functions

import (
	"lab06/structures"
	"time"
)

func ChoosePeriod(data []structures.Data, beginning time.Time, end time.Time) []structures.Data {
	var newData []structures.Data
	for _, d := range data {
		if !d.Date.Before(beginning) && !d.Date.After(end) {
			newData = append(newData, d)
		}
	}
	return newData
}