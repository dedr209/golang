package main

import "math"

type LCG struct {
	a, c, m, current int
}

func NewLCG(seed int) *LCG {
	return &LCG{
		a:       1103515245,
		c:       12345,
		m:       1 << 31, // 2^31
		current: seed,
	}
}

func (lcg *LCG) Next() int {
	lcg.current = (lcg.a*lcg.current + lcg.c) % lcg.m
	return lcg.current
}

func (lcg *LCG) NextIntn(n int) int {
	return lcg.Next() % n
}

func (lcg *LCG) NextFloat() float64 {
	return float64(lcg.Next()) / float64(lcg.m)
}

func CalculateStats(frequencies map[int]int, k int) (float64, float64, float64) {
	var expectedValue float64
	var variance float64

	probabilities := make(map[int]float64)

	for x, freq := range frequencies {
		p := float64(freq) / float64(k) // P = L / K
		probabilities[x] = p
		expectedValue += float64(x) * p // M(X) = Sum(x_i * p_i)
	}

	for x, p := range probabilities {
		diff := float64(x) - expectedValue
		variance += (diff * diff) * p // D(X) = Sum((x_i - M(X))^2 * p_i)
	}

	stdDev := math.Sqrt(variance) // Sigma = sqrt(D(X))

	return expectedValue, variance, stdDev
}

func main() {
}
