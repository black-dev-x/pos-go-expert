package tax

type Repository interface {
	SaveTax(tax float64) error
}

func CalculateTaxAndSave(amount float64, repository Repository) error {
	tax := CalculateTax(amount)
	err := repository.SaveTax(tax)
	return err
}

func CalculateTax(amount float64) float64 {
	if amount <= 0 {
		return 0
	}
	if amount < 1000 {
		return 5
	}
	if amount < 20000 {
		return 10
	}
	return 20
}
