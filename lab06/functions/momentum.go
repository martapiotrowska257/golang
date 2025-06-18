package functions

import (
	"lab06/structures"
)

func RSI(data []structures.Data) float64 {
	n := len(data)

	if n == 0 {
		return 0.0
	}

	rise := make([]float64, n)
	fall := make([]float64, n)
	for i := range n {
		if i == 0 {
			rise[i] = 0
			fall[i] = 0
		} else {
			diff := data[i].Last - data[i-1].Last
			if diff > 0 {
				rise[i] = diff
				fall[i] = 0
			} else {
				rise[i] = 0
				fall[i] = -diff
			}
		}
	}

	avgRise := 0.0
	for _, r := range rise {
		avgRise += r
	}
	avgRise /= float64(n)

	avgFall := 0.0
	for _, f := range fall {
		avgFall += f
	}
	avgFall /= float64(n)

	if avgFall == 0 {
		return 100.0
	}

	rs := avgRise / avgFall
	rsi := 100 - (100 / (1 + rs))

	return rsi
}
