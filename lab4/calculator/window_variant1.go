package calculator

import "errors"

func CalculateWindowPriceGo(width, height float64, material, chambers int, hasSill bool) (float64, error) {
	if width <= 0 || height <= 0 {
		return 0, errors.New("width and height must be positive")
	}
	if chambers < 1 || chambers > 2 {
		return 0, errors.New("chambers must be 1 or 2")
	}
	rates := [][]float64{
		{2.5, 3.0}, // Wood
		{0.5, 1.0}, // Metal
		{1.5, 2.0}, // Metal-plastic
	}
	if material < 0 || material >= len(rates) {
		return 0, errors.New("invalid material")
	}
	area := width * height
	total := area * rates[material][chambers-1]
	if hasSill {
		total += 350.0
	}
	return total, nil
}
