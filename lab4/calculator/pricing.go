package calculator

import "fmt"

const (
	CountryBulgaria = iota
	CountryGermany
	CountryPoland
)

const (
	SeasonSummer = iota
	SeasonWinter
)

const (
	RoomStandard = iota
	RoomLux
)

const (
	guideDailyFee = 50.0
	luxMarkup     = 0.20
)

var dailyRateTable = [3][2]float64{
	{100, 150},
	{160, 200},
	{120, 180},
}

func validateInput(days int, country, season int, vouchers int, roomType int) error {
	if days <= 0 {
		return fmt.Errorf("days must be positive")
	}

	if country < 0 || country >= len(dailyRateTable) {
		return fmt.Errorf("unknown country index: %d", country)
	}

	if season < 0 || season >= len(dailyRateTable[0]) {
		return fmt.Errorf("unknown season index: %d", season)
	}

	if vouchers <= 0 {
		return fmt.Errorf("vouchers must be positive")
	}

	if roomType != RoomStandard && roomType != RoomLux {
		return fmt.Errorf("unknown room type index: %d", roomType)
	}

	return nil
}

func boolToInt(value bool) int {
	if value {
		return 1
	}
	return 0
}
