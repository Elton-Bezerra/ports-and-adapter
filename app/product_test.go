package app_test

import (
	"testing"

	"github.com/Elton-Bezerra/ports-and-adapter/app"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := app.Product{}
	product.Name = "Hello"
	product.Status = app.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()

	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := app.Product{}
	product.Name = "Hello"
	product.Status = app.DISABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "the price must be zero in order to have the product disabled", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := app.Product{}
	product.ID = uuid.NewString()
	product.Name = "Hello"
	product.Status = app.DISABLED
	product.Price = 10

	_, err := product.IsValid()

	require.Nil(t, err)

	product.Status = "abrobrinha"
	_, err = product.IsValid()
	require.Equal(t, "status must be enabled or disabled", err.Error())

	product.Status = app.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "the price must be greater or equal to 0", err.Error())

	product.Price = 10
	product.Name = ""
	product.Status = ""
	_, err = product.IsValid()
	require.Equal(t, "Name: non zero value required", err.Error())

}
