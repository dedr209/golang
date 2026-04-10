package calculator

/*
double calculate_window_price(double width, double height, int material, int glass, int with_windowsill);
*/
import "C"

// CalculateWindowPriceC computes the total price in C and returns the result to Go.
func CalculateWindowPriceC(width, height float64, material, glass int, withWindowsill bool) (float64, error) {
	if err := validateInput(width, height, material, glass); err != nil {
		return 0, err
	}

	total := C.calculate_window_price(
		C.double(width),
		C.double(height),
		C.int(material),
		C.int(glass),
		C.int(boolToInt(withWindowsill)),
	)

	return float64(total), nil
}
