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
	if err != nil {
		require.Equal(t, "The price must be greater than zero to enable the product", err.Error())
	}
}

func TestProduct_Disable(t *testing.T) {
	product := Product{}
	product.Name = "Hello"
	product.Status = ENABLED
	product.Price = 10

	err := product.Disable()
	if err != nil {
		require.Equal(t, "The price must be zero to disable the product", err.Error())
	}

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

func TestProduct_GetId(t *testing.T) {
	product := Product{}
	newUUID := uuid.NewV4().String()
	product.ID = newUUID
	require.NotEmpty(t, product.GetId())
	require.Equal(t, newUUID, product.GetId())
}

func TestProduct_GetName(t *testing.T) {
	product := Product{}
	product.Name = "Hello"
	require.NotEmpty(t, product.GetName())
	require.Equal(t, "Hello", product.GetName())
}

func TestProduct_GetStatus(t *testing.T) {
	product := Product{}
	product.Status = ENABLED
	require.NotEmpty(t, product.GetStatus())
	require.Equal(t, ENABLED, product.GetStatus())

	product.Status = DISABLED
	require.Equal(t, DISABLED, product.GetStatus())
}

func TestProduct_GetPrice(t *testing.T) {
	product := Product{}
	product.Price = 10
	require.NotEmpty(t, product.GetPrice())
	require.Equal(t, 10.0, product.GetPrice())
}
