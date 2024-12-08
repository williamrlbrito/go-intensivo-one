package entity_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/williamrlbrito/go-intensivo-one/internal/order/entity"
)

func TestGivenAnEmptyIdWhenCreateOrderThenShouldReceiveError(t *testing.T) {
	order := entity.Order{}
	assert.Error(t, order.IsValid(), "invalid ID")
}

func TestGivenAnEmptyPriceWhenCreateOrderThenShouldReceiveError(t *testing.T) {
	order := entity.Order{ID: 1}
	assert.Error(t, order.IsValid(), "invalid Price")
}

func TestGivenAnEmotyTaxWhenCreateOrderThenShouldReceiveError(t *testing.T) {
	order := entity.Order{ID: 1, Price: 10}
	assert.Error(t, order.IsValid(), "invalid Tax")
}

func TestGivenAValidParamsWhenCallNewOrderThenShouldReceiveOrderWithAllParams(t *testing.T) {
	order, err := entity.NewOrder(1, 10, 1)
	assert.NoError(t, err)
	assert.Equal(t, order.ID, 1)
	assert.Equal(t, order.Price, 10.0)
	assert.Equal(t, order.Tax, 1.0)
}

func TestGivenAValidParamsWhenCallCalculateFinalPriceThenShouldCalculateFinalPriceAndSetItOnFinalPriceProperty(t *testing.T) {
	order, _ := entity.NewOrder(1, 10, 2)
	order.CalculateFinalPrice()
	assert.Equal(t, order.FinalPrice, 12.0)
}
