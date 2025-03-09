package tax

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateTax(t *testing.T) {
	tax, error := CalculateTax(1000)
	assert.Equal(t, 10.0, tax)
	assert.Nil(t, error)

	tax, error = CalculateTax(0)
	assert.Equal(t, 0.0, tax)
	assert.Error(t, error, "amount should be greater than 0")
}
