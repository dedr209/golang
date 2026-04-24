package calculator

import "testing"

func TestCalculateTourPriceVariants(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		days      int
		country   int
		season    int
		vouchers  int
		roomType  int
		withGuide bool
		want      float64
	}{
		{name: "Bulgaria_Summer_Standard_NoGuide", days: 1, country: CountryBulgaria, season: SeasonSummer, vouchers: 1, roomType: RoomStandard, withGuide: false, want: 100},
		{name: "Bulgaria_Winter_Standard_NoGuide", days: 2, country: CountryBulgaria, season: SeasonWinter, vouchers: 1, roomType: RoomStandard, withGuide: false, want: 300},
		{name: "Germany_Summer_Standard_NoGuide", days: 2, country: CountryGermany, season: SeasonSummer, vouchers: 3, roomType: RoomStandard, withGuide: false, want: 960},
		{name: "Germany_Winter_Standard_NoGuide", days: 3, country: CountryGermany, season: SeasonWinter, vouchers: 2, roomType: RoomStandard, withGuide: false, want: 1200},
		{name: "Poland_Summer_Standard_NoGuide", days: 3, country: CountryPoland, season: SeasonSummer, vouchers: 2, roomType: RoomStandard, withGuide: false, want: 720},
		{name: "Poland_Winter_Standard_NoGuide", days: 4, country: CountryPoland, season: SeasonWinter, vouchers: 1, roomType: RoomStandard, withGuide: false, want: 720},
		{name: "Bulgaria_Summer_Lux_WithGuide", days: 2, country: CountryBulgaria, season: SeasonSummer, vouchers: 2, roomType: RoomLux, withGuide: true, want: 580},
		{name: "Germany_Winter_Lux_WithGuide", days: 4, country: CountryGermany, season: SeasonWinter, vouchers: 3, roomType: RoomLux, withGuide: true, want: 3080},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			goTotal, err := CalculateTourPriceGo(tt.days, tt.country, tt.season, tt.vouchers, tt.roomType, tt.withGuide)
			if err != nil {
				t.Fatalf("unexpected Go error: %v", err)
			}

			cTotal, err := CalculateTourPriceC(tt.days, tt.country, tt.season, tt.vouchers, tt.roomType, tt.withGuide)
			if err != nil {
				t.Fatalf("unexpected C error: %v", err)
			}

			assertNearlyEqual(t, goTotal, tt.want)
			assertNearlyEqual(t, cTotal, tt.want)
			assertNearlyEqual(t, cTotal, goTotal)
		})
	}
}

func TestCalculateTourPriceValidation(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		days     int
		country  int
		season   int
		vouchers int
		roomType int
	}{
		{name: "ZeroDays", days: 0, country: CountryBulgaria, season: SeasonSummer, vouchers: 1, roomType: RoomStandard},
		{name: "NegativeDays", days: -1, country: CountryBulgaria, season: SeasonSummer, vouchers: 1, roomType: RoomStandard},
		{name: "InvalidCountryLow", days: 1, country: -1, season: SeasonSummer, vouchers: 1, roomType: RoomStandard},
		{name: "InvalidCountryHigh", days: 1, country: len(dailyRateTable), season: SeasonSummer, vouchers: 1, roomType: RoomStandard},
		{name: "InvalidSeasonLow", days: 1, country: CountryBulgaria, season: -1, vouchers: 1, roomType: RoomStandard},
		{name: "InvalidSeasonHigh", days: 1, country: CountryBulgaria, season: len(dailyRateTable[0]), vouchers: 1, roomType: RoomStandard},
		{name: "ZeroVouchers", days: 1, country: CountryBulgaria, season: SeasonSummer, vouchers: 0, roomType: RoomStandard},
		{name: "NegativeVouchers", days: 1, country: CountryBulgaria, season: SeasonSummer, vouchers: -1, roomType: RoomStandard},
		{name: "InvalidRoomTypeLow", days: 1, country: CountryBulgaria, season: SeasonSummer, vouchers: 1, roomType: -1},
		{name: "InvalidRoomTypeHigh", days: 1, country: CountryBulgaria, season: SeasonSummer, vouchers: 1, roomType: RoomLux + 1},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			goTotal, goErr := CalculateTourPriceGo(tt.days, tt.country, tt.season, tt.vouchers, tt.roomType, false)
			cTotal, cErr := CalculateTourPriceC(tt.days, tt.country, tt.season, tt.vouchers, tt.roomType, false)

			if goErr == nil {
				t.Fatal("expected Go variant to return an error")
			}
			if cErr == nil {
				t.Fatal("expected C variant to return an error")
			}
			if goErr.Error() != cErr.Error() {
				t.Fatalf("Go error %q does not match C error %q", goErr.Error(), cErr.Error())
			}
			if goTotal != 0 || cTotal != 0 {
				t.Fatalf("expected zero totals on validation failure, got Go=%.2f C=%.2f", goTotal, cTotal)
			}
		})
	}
}

func assertNearlyEqual(t *testing.T, got, want float64) {
	t.Helper()

	const epsilon = 1e-9
	if got-want > epsilon || want-got > epsilon {
		t.Fatalf("got %.10f, want %.10f", got, want)
	}
}
