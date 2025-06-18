package functions

import "strconv"

func ParseInt32(s string) (int32, error) {
	val, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(val), nil
}
