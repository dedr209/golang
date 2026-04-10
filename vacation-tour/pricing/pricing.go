package pricing

import (
	"fmt"
	"strings"
)

const guideFeePerDay = 50.0
const luxuryMultiplier = 1.20

type seasonPrice struct {
	summer float64
	winter float64
}

var destinationPrices = map[string]seasonPrice{
	"bulgaria": {summer: 100, winter: 150},
	"germany":  {summer: 160, winter: 200},
	"poland":   {summer: 120, winter: 180},
}

func destinationSeasonPrice(destination, season string) (float64, error) {
	prices, ok := destinationPrices[strings.ToLower(strings.TrimSpace(destination))]
	if !ok {
		return 0, fmt.Errorf("unknown destination: %s", destination)
	}

	switch strings.ToLower(strings.TrimSpace(season)) {
	case "summer":
		return prices.summer, nil
	case "winter":
		return prices.winter, nil
	default:
		return 0, fmt.Errorf("unknown season: %s", season)
	}
}

func CalculateTourPrice(destination, season string, days int, roomType string) (float64, error) {
	if days <= 0 {
		return 0, fmt.Errorf("days must be positive")
	}

	basePrice, err := destinationSeasonPrice(destination, season)
	if err != nil {
		return 0, err
	}

	if strings.EqualFold(strings.TrimSpace(roomType), "luxury") {
		basePrice *= luxuryMultiplier
	}

	guideTotal := float64(days) * guideFeePerDay
	return basePrice + guideTotal, nil
}
