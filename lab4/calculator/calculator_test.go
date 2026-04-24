package calculator

import "testing"

// --- Tour Calculator Tests ---
func TestCalculateTourPriceGo(t *testing.T) {
	total, err := CalculateTourPriceGo(10, CountryBulgaria, SeasonSummer, 2, RoomStandard, false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := 10.0 * 2.0 * 100.0 // 2000
	if total != want {
		t.Fatalf("got %.2f, want %.2f", total, want)
	}
}
func TestCalculateTourPriceC(t *testing.T) {
	goTotal, _ := CalculateTourPriceGo(7, CountryGermany, SeasonWinter, 1, RoomLux, true)
	cTotal, err := CalculateTourPriceC(7, CountryGermany, SeasonWinter, 1, RoomLux, true)
	if err != nil {
		t.Fatalf("unexpected C error: %v", err)
	}
	if cTotal != goTotal {
		t.Fatalf("cgo total %.2f does not match Go total %.2f", cTotal, goTotal)
	}
}

// --- Window Calculator Tests ---
func TestCalculateWindowPriceC(t *testing.T) {
	// Test: 130x90, Wood (0), 1-chamber (1), with sill
	// Area = 11700. Rate = 2.5. Base = 29250. + 350 = 29600
	total, err := CalculateWindowPriceC(130, 90, 0, 1, true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := 29600.0
	if total != want {
		t.Fatalf("got %.2f, want %.2f", total, want)
	}
}
