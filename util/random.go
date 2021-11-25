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

// 随机,M中取N,N<=M
func RandomNOfM(n int, m int) (idxs []int) {
	if n > m {
		return nil
	}

	idx := make([]int, m)
	idx[0] = -1 //索引0特殊处理
	for i := 0; i < n; i++ {
		r := i + rand.Intn(m-i)
		if idx[i] == 0 {
			idx[i] = i
		}
		if idx[r] == 0 {
			idx[r] = r
		}
		idx[i], idx[r] = idx[r], idx[i]
		if idx[i] == -1 {
			idx[i] = 0
		}
	}

	return idx[:n]
}

//洗牌
func Shuffle(s []int) {
	l := len(s)
	if l < 2 {
		return
	}

	for i := 0; i < l; i++ {
		r := rand.Intn(l)
		s[i], s[r] = s[r], s[i]
	}
}
