package util

import "math/rand"

//rand  [0,n)
func GetRandom(min int, max int) int {
	if min > max {
		temp := min
		min = max
		max = temp
	}

	distance := max - min
	return rand.Intn(distance) + min
}

func GetRandomFloat64(min float64, max float64) float64 {
	if min > max {
		temp := min
		min = max
		max = temp
	}

	rdm := rand.Float64()
	return rdm*(max-min) + min
}
