package calculator

// CalculateTourPriceGo computes the order total entirely in Go.
func CalculateTourPriceGo(days int, country, season int, vouchers int, roomType int, withGuide bool) (float64, error) {
	if err := validateInput(days, country, season, vouchers, roomType); err != nil {
		return 0, err
	}

	baseTotal := float64(days*vouchers) * dailyRateTable[country][season]
	if roomType == RoomLux {
		baseTotal *= 1 + luxMarkup
	}

	total := baseTotal
	if withGuide {
		total += float64(days) * guideDailyFee
	}

	return total, nil
}
