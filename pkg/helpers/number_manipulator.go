package helpers

import "math"

// Round float generator
// example here: https://go.dev/play/p/VmTGHhV2_WI
func RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

// handle not return nil
func CheckArrayIntNil(data []int) []int {
	if len(data) > 0 {
		return data
	}
	return []int{}
}

// Return 0 if data is nil, otherwise return the value of data
func CheckIntValue(data *int) int {
	if data == nil {
		return 0
	}
	return *data
}
