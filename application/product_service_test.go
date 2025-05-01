package application_test

import (
	"arquitetura-hexagonal-go/application"
	mock_application "arquitetura-hexagonal-go/application/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	repository := mock_application.NewMockProductRepositoryInterface(ctrl)
	repository.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		ProductRepository: repository,
	}

	result, err := service.Get("123")
	require.Nil(t, err)
	require.Equal(t, product, result)
}