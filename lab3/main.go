package main

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

func main() {
}
