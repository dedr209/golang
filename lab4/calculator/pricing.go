package calculator

import "fmt"

const (
	MaterialWood = iota
	MaterialMetal
	MaterialMetalPlastic
)

const (
	GlassSingle = iota
	GlassDouble
)

const windowsillFee = 350.0

var priceTable = [3][2]float64{
	{2.5, 3.0},
	{0.5, 1.0},
	{1.5, 2.0},
}

func validateInput(width, height float64, material, glass int) error {
	if width <= 0 || height <= 0 {
		return fmt.Errorf("width and height must be positive")
	}

	if material < 0 || material >= len(priceTable) {
		return fmt.Errorf("unknown material index: %d", material)
	}

	if glass < 0 || glass >= len(priceTable[0]) {
		return fmt.Errorf("unknown glass index: %d", glass)
	}

	return nil
}

func boolToInt(value bool) int {
	if value {
		return 1
	}
	return 0
}
