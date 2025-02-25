package tax

import "testing"

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expected := 6.0 // 5.0 will pass, 6.0 will fail

	result := CalculateTax(amount)

	if result != expected {
		t.Errorf("Expected %.2f, but got %.2f", expected, result)
	}
}
