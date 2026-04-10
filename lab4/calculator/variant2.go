package calculator

/*
double calculate_tour_price(int days, int country, int season, int vouchers, int room_type, int with_guide);
*/
import "C"

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
