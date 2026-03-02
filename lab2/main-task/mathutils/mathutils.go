package mathutils

import "fmt"

func MinOfTree(a, b, c float64) float64 {
	if a < b && a < c {
		return a
	}
	if b < a && b < c {
		return b
	}
	return c
}

func AverageOfTree(a, b, c float64) float64 {
	return (a + b + c) / 3
}

func SolveLinear(a, b float64) (float64, error) {
	if a == 0 {
		if b == 0 {
			return 0, fmt.Errorf("Infinite solutions")
		}
		return 0, fmt.Errorf("No solution, a cannot be zero")
	}
	return -b / a, nil
}
