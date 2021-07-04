package parse

import (
	"strconv"
)

func StringToFloat(value string) (float64, error) {
	return strconv.ParseFloat(value, 64)
}

func MustGetFloat(value string) float64 {
	f, _ := StringToFloat(value)
	return f
}

func IntAsString(value interface{}) string {
	return strconv.FormatInt(value.(int64), 10)
}

func BytesAsString(value interface{}) string {
	return string(value.([]byte))
}

func MustGetInt(value string) int64 {
	i, _ := strconv.ParseInt(value, 10, 64)
	return i
}
