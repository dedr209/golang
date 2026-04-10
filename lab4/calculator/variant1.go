package calculator

// CalculateWindowPriceGo computes the total price entirely in Go.
func CalculateWindowPriceGo(width, height float64, material, glass int, withWindowsill bool) (float64, error) {
	if err := validateInput(width, height, material, glass); err != nil {
		return 0, err
	}

	total := width * height * priceTable[material][glass]
	if withWindowsill {
		total += windowsillFee
	}

	return total, nil
}
