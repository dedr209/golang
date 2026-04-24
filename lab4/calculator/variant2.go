package calculator

/*
#include "pricing.c"
*/

import "C"

func calculateWindowPriceViaC(width, height float64, material, chambers, hasSill int) float64 {
	result := C.calculate_window_price(
		C.double(width),
		C.double(height),
		C.int(material),
		C.int(chambers),
		C.int(hasSill),
	)
	return float64(result)
}

// CalculateTourPriceC computes the order total in C and returns the result to Go.
func CalculateTourPriceC(days int, country, season int, vouchers int, roomType int, withGuide bool) (float64, error) {
	if err := validateInput(days, country, season, vouchers, roomType); err != nil {
		return 0, err
	}
	total := C.calculate_tour_price(
		C.int(days),
		C.int(country),
		C.int(season),
		C.int(vouchers),
		C.int(roomType),
		C.int(boolToInt(withGuide)),
	)
	return float64(total), nil
}
