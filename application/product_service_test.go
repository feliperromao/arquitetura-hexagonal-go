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

func TestProductService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	repository := mock_application.NewMockProductRepositoryInterface(ctrl)
	repository.EXPECT().Create(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		ProductRepository: repository,
	}

	result, err := service.Create("Product Name", 100.0)
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductService_Enable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	product.EXPECT().Enable().Return(nil).AnyTimes()

	repository := mock_application.NewMockProductRepositoryInterface(ctrl)
	repository.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()
	repository.EXPECT().Update(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		ProductRepository: repository,
	}

	result, err := service.Enable("123")
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductService_Disable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	product.EXPECT().Disable().Return(nil).AnyTimes()

	repository := mock_application.NewMockProductRepositoryInterface(ctrl)
	repository.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()
	repository.EXPECT().Update(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		ProductRepository: repository,
	}

	result, err := service.Disable("123")
	require.Nil(t, err)
	require.Equal(t, product, result)
}