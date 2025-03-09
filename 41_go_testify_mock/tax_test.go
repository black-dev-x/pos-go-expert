package tax

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	repository := &TaxRepositoryMock{}
	repository.On("SaveTax", 10.0).Return(nil)
	repository.On("SaveTax", 0.0).Return(errors.New("amount should be greater than 0"))
	error := CalculateTaxAndSave(1000.0, repository)
	assert.Nil(t, error)

	error = CalculateTaxAndSave(0.0, repository)
	assert.Error(t, error, "amount should be greater than 0")
	repository.AssertExpectations(t)
}
