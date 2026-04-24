package calculator

/*
#include "pricing.h"
*/
import "C"
import "fmt"

func CalculateWindowPriceC(width, height float64, material, chambers int, hasSill bool) (float64, error) {
	// Додайте перевірку вхідних даних, як у варіанті на Go
	if width <= 0 || height <= 0 {
		return 0, fmt.Errorf("width and height must be positive")
	}

	result := C.calculate_window_price(
		C.double(width),
		C.double(height),
		C.int(material),
		C.int(chambers),
		C.int(boolToInt(hasSill)),
	)
	return float64(result), nil
}

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
