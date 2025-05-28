package structures

import "time"

type Data struct {
	Date   time.Time
	Last   float64
	Volume int32
	Open   float64
	High   float64
	Low    float64
}