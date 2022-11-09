package entity

import "testing"

func TestGivenAnEmptyID_WhenCreateANewOrder_ThenShouldRecieveAnError(t *testing.T) {
	order := Order()
	assert.Error(t, order.)
}