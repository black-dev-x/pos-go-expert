package tax

import (
	"errors"
	"time"
)

func CalculateTax(amount float64) (float64, error) {
	if amount <= 0 {
		return 0, errors.New("amount should be greater than 0")
	}
	if amount < 1000 {
		return 5, nil
	}
	if amount < 20000 {
		return 10, nil
	}
	return 20, nil
}

func CalculateTax2(amount float64) float64 {
	time.Sleep(time.Millisecond)
	if amount >= 1000.0 {
		return 10.0
	}
	return 5.0
}
