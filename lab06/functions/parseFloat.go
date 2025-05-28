package functions

import (
	"strconv"
	"strings"
)

func ParseFloat64(s string) (float64, error) {
	cleanedString := strings.ReplaceAll(s, "$", "")
	val, err := strconv.ParseFloat(cleanedString, 64)
	if err != nil {
		return 0, err
	}
	return float64(val), nil
}