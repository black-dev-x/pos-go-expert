package tax

import "testing"

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expected := 5.0 // 5.0 will pass, 6.0 will fail

	result := CalculateTax(amount)

	if result != expected {
		t.Errorf("Expected %.2f, but got %.2f", expected, result)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount, expected float64
	}
	table := []calcTax{
		{500.0, 5.0},
		{1000.0, 10.0},
		{1500.0, 10.0},
	}
	for _, v := range table {
		result := CalculateTax(v.amount)
		if result != v.expected {
			t.Errorf("Expected %.2f, but got %.2f", v.expected, result)
		}
	}
}

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.0)
	}
}

func BenchmarkCalculateTax2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax2(500.0)
	}
}
