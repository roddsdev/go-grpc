package model_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"

	"example/grpc/domain/model"

	"github.com/stretchr/testify/require"
)

func TestModel_NewProduct(t *testing.T) {
	name := "Teclado Razr"
	description := "Teclado Game Razr cheio de RGB"
	price := 200
	product, err := model.NewProduct(name, description, float32(price))

	require.Nil(t, err)
	require.NotEmpty(t, uuid.FromStringOrNil(product.ID))
	require.Equal(t, product.Name, name)
	require.Equal(t, product.Description, description)
	require.Equal(t, product.Price, price)
}
