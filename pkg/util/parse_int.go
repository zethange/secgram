package util

import "strconv"

func ParseInt(s string, d int64) int64 {
	res, err := strconv.Atoi(s)
	if err != nil {
		return d
	}
	return int64(res)
}
