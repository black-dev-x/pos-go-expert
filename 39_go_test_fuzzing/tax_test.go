package tax

import "testing"

func FuzzCalculateTax(f *testing.F) {
	seed := []float64{500.0, 1000.0, 1500.0}
	for _, s := range seed {
		f.Add(s)
	}
	f.Fuzz(func(t *testing.T, amount float64) {
		CalculateTax(amount)
	})
}
