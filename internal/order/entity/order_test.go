package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyID_WhenCreateANewOrder_ThenShouldRecieveAnError(t *testing.T) {
	order := Order{}
	assert.Error(t, order.IsValid(), "invalid id")
}

func TestGivenAnEmptyPrice_WhenCreateANewOrder_ThenShouldRecieveAnError(t *testing.T) {
	order := Order{ID: "123"}
	assert.Error(t, order.IsValid(), "invalid price")
}

func TestGivenAnEmptyTax_WhenCreateANewOrder_ThenShouldRecieveAnError(t *testing.T) {
	order := Order{ID: "123", Price: 10}
	assert.Error(t, order.IsValid(), "invalid tax")
}

func TestGivenAValidParams_WhenICallNewOrder_ThenIShouldReceiveCreateOrderWithAllParams(t *testing.T) {
	order := Order{ID: "123", Price: 10, Tax: 2}
	assert.Equal(t, order.ID, "123")
	assert.Equal(t, order.Price, 10.0)
	assert.Equal(t, order.Tax, 2.0)
	assert.Nil(t, order.IsValid())
}

func TestGivenAValidParams_WhenICallNewOrderFunc_ThenIShouldReceiveCreateOrderWithAllParams(t *testing.T) {
	order, err := NewOrder("123", 10, 2)
	assert.Nil(t, err)
	assert.Equal(t, order.ID, "123")
	assert.Equal(t, order.Price, 10.0)
	assert.Equal(t, order.Tax, 2.0)
}

func TestGivenAPriceAndTax_WhenICallCalculateTax_WhenIShowldSetFinalPrice(t *testing.T) {
	order, err := NewOrder("123", 10, 2)
	assert.Nil(t, err)
	assert.Nil(t, order.CalculateFinalPrice())
	assert.Equal(t, order.FinalPrice, 12.0)
}
