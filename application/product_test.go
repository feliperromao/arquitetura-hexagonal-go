package application_test

import (
	"arquitetura-hexagonal-go/application"
	"testing"

	"github.com/stretchr/testify/require"
	uuidv4 "github.com/satori/go.uuid"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Test Product"
	product.Price = 10.0
	product.Status = application.DISABLED

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0.0
	err = product.Enable()
	require.NotNil(t, err)
	require.Equal(t, "product price must be greater than 0 to enable", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "Test Product 2"
	product.Price = 10.0
	product.Status = application.ENABLED

	err := product.Disable()
	require.Equal(t, "product price must be 0 to disable", err.Error())

	product.Price = 0.0
	err = product.Disable()
	require.Nil(t, err)
	require.Equal(t, application.DISABLED, product.Status)
	require.Equal(t, 0.0, product.Price)
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.ID = uuidv4.NewV4().String()
	product.Name = "Test Product 3"
	product.Price = 10.0
	product.Status = application.DISABLED

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID_STATUS"
	_, err = product.IsValid()
	require.NotNil(t, err)
	require.Equal(t, "invalid product status, must be ENABLED or DISABLED", err.Error())

	product.Status = application.DISABLED
	product.Price = -10.0
	_, err = product.IsValid()
	require.NotNil(t, err)
	require.Equal(t, "product price must be greater than or equal to 0", err.Error())
}