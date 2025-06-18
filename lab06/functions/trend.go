package functions

import (
	"math"
	"lab06/structures"
)

func EMA(data []structures.Data) float64 {
	n := len(data)

	if n == 0 {
		return 0.0
	}
	
	alfa := 2.0 / float64(n + 1)

	var numerator float64
	var denominator float64

	for i := range n {
		numerator += data[n-i-1].Last*math.Pow(1-alfa, float64(i))
		denominator += + math.Pow(1-alfa, float64(i))
	}

	ema := numerator / denominator

	return ema

}
