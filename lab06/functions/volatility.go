package functions

import (
	"math"
	"lab06/structures"
)

func ATR(data []structures.Data) float64 {
	n := len(data)
	if n == 0 {
		return 0.0
	}
	TR := make([]float64, n)
	for i := 1; i < n; i++ {
		highLow := data[i].High - data[i].Low
		highPrevClose := math.Abs(data[i].High - data[i-1].Last)
		lowPrevClose := math.Abs(data[i].Low - data[i-1].Last)

		TR[i] = math.Max(highLow, math.Max(highPrevClose, lowPrevClose))
	}

	m := len(TR)
	var atr float64

	for i := range m {
		if i == 0 {
			for j := 1; j < m; j++ {
				atr += TR[j]
			}
			atr /= float64(m - 1)
		}
		if i > 0 {
			atr = (atr*(float64(m-1)) + TR[i]) / float64(m)
		}
	}
	return atr
}