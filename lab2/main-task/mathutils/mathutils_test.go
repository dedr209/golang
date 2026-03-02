package mathutils

import (
	"math"
	"testing"
)

func floatEquals(a, b float64) bool {
	const eps = 1e-9
	return math.Abs(a-b) < eps
}

func TestMinOfTree(t *testing.T) {
	tests := []struct {
		a, b, c float64
		want    float64
	}{
		{1, 2, 3.3, 1},
		{3, 2, 1, 1},
		{2, 1, 3, 1},
		{-1, 0, 1, -1},
		{0, 0, 0, 0},
	}
	for _, tt := range tests {
		got := MinOfTree(tt.a, tt.b, tt.c)
		if !floatEquals(got, tt.want) {
			t.Errorf("MinOfTree(%v, %v, %v) = %v; want %v", tt.a, tt.b, tt.c, got, tt.want)
		}
	}
}

func TestAverageOfTree(t *testing.T) {
	got := AverageOfTree(1, 2, 3)
	want := 2.0
	if !floatEquals(got, want) {
		t.Errorf("AverageOfTree(1,2,3) = %v; want %v", got, want)
	}
}

func TestSolveLinear(t *testing.T) {
	x, err := SolveLinear(2, -4)
	if err != nil || !floatEquals(x, 2) {
		t.Errorf("SolveLinear(2, -4) = %v, %v; want 2, nil", x, err)
	}
	_, err = SolveLinear(0, 0)
	if err == nil {
		t.Error("SolveLinear(0, 0) should return error for infinite solutions")
	}
	_, err = SolveLinear(0, 1)
	if err == nil {
		t.Error("SolveLinear(0, 1) should return error for no solution")
	}
}
