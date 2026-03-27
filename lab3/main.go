package main

import (
	"fmt"
	"math"
)

type LCG struct {
	a, c, m, current int
}

func NewLCG(seed int) *LCG {
	return &LCG{
		a:       1664525,
		c:       1013904223,
		m:       1 << 32, // 2^31
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
	n := 150
	k := 20000
	seed := 1

	lcg := NewLCG(seed)
	frequencies := make(map[int]int)

	// Task 1: Generate Integer Sequence & Frequencies
	for i := 0; i < k; i++ {
		val := lcg.NextIntn(n)
		frequencies[val]++
	}

	expectedValue, variance, stdDev := CalculateStats(frequencies, k)

	fmt.Printf("--- Integer Sequence Statistics (K=%d, Range=[0, %d)) ---\n", k, n)
	fmt.Printf("Mathematical Expectation M(X): %.4f\n", expectedValue)
	fmt.Printf("Variance D(X):                 %.4f\n", variance)
	fmt.Printf("Standard Deviation σ(X):       %.4f\n\n", stdDev)

	// Task 2: Generate Float Sequence Sample
	lcgFloat := NewLCG(seed)
	fmt.Printf("--- Float Sequence Sample ---\n")
	for i := 0; i < 5; i++ {
		fmt.Printf("Float %d: %.6f\n", i+1, lcgFloat.NextFloat())
	}
}
