package calculator
/*
double calculate_window_price(double width, double height, int material, int chambers, int has_sill);
*/
import "C"
import "errors"
func CalculateWindowPriceC(width, height float64, material, chambers int, hasSill bool) (float64, error) {
if width <= 0 || height <= 0 {
return 0, errors.New("width and height must be positive")
}
if chambers < 1 || chambers > 2 {
return 0, errors.New("chambers must be 1 or 2")
}
if material < 0 || material > 2 {
return 0, errors.New("invalid material")
}
sillInt := 0
if hasSill {
sillInt = 1
}
total := C.calculate_window_price(
C.double(width),
C.double(height),
C.int(material),
C.int(chambers),
C.int(sillInt),
)
return float64(total), nil
}
