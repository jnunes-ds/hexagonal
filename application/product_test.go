package application

import (
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProduct_Enable(t *testing.T) {
	product := Product{}
	product.Name = "Hello"
	product.Status = DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "The price must be greater than zero to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := Product{}
	product.Name = "Hello"
	product.Status = ENABLED
	product.Price = 10

	err := product.Disable()
	require.Equal(t, "The price must be zero to disable the product", err.Error())

	product.Price = 0
	err = product.Disable()
	require.Nil(t, err)
}

func TestProduct_IsValid(t *testing.T) {
	product := Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Hello"
	product.Status = DISABLED
	product.Price = 10

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, "The status must be enabled or disabled", err.Error())

	product.Status = ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "The price must be greater than zero", err.Error())
}
