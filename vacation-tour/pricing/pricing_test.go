package pricing

import "testing"

func TestCalculateTourPriceStandardSummer(t *testing.T) {
	total, err := CalculateTourPrice("Bulgaria", "summer", 3, "standard")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if total != 250 {
		t.Fatalf("got %.2f, want 250.00", total)
	}
}

func TestCalculateTourPriceLuxuryWinter(t *testing.T) {
	total, err := CalculateTourPrice("Germany", "winter", 2, "luxury")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if total != 340 {
		t.Fatalf("got %.2f, want 340.00", total)
	}
}

func TestCalculateTourPriceBadSeason(t *testing.T) {
	_, err := CalculateTourPrice("Poland", "spring", 4, "standard")
	if err == nil {
		t.Fatal("expected error for unknown season")
	}
}
