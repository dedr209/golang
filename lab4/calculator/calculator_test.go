package calculator

import "testing"

func TestCalculateWindowPriceGo(t *testing.T) {
	total, err := CalculateWindowPriceGo(120, 140, MaterialWood, GlassDouble, true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := 120*140*3.0 + 350.0
	if total != want {
		t.Fatalf("got %.2f, want %.2f", total, want)
	}
}

func TestCalculateWindowPriceC(t *testing.T) {
	goTotal, err := CalculateWindowPriceGo(130, 150, MaterialMetalPlastic, GlassSingle, false)
	if err != nil {
		t.Fatalf("unexpected Go error: %v", err)
	}

	cTotal, err := CalculateWindowPriceC(130, 150, MaterialMetalPlastic, GlassSingle, false)
	if err != nil {
		t.Fatalf("unexpected C error: %v", err)
	}

	if cTotal != goTotal {
		t.Fatalf("cgo total %.2f does not match Go total %.2f", cTotal, goTotal)
	}
}

func TestCalculateWindowPriceValidation(t *testing.T) {
	_, err := CalculateWindowPriceGo(0, 150, MaterialWood, GlassSingle, false)
	if err == nil {
		t.Fatal("expected validation error for zero width")
	}
}
