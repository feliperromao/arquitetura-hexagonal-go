package cli_test

import (
	"arquitetura-hexagonal-go/adapters/cli"
	"arquitetura-hexagonal-go/application"
	mock_application "arquitetura-hexagonal-go/application/mocks"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productId := "abc-123"
	productName := "any_name"
	productPrice := 29.19
	productStatus := application.DISABLED

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetID().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	serviceMock := mock_application.NewMockProductServiceInterface(ctrl)
	serviceMock.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	serviceMock.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	serviceMock.EXPECT().Enable(productId).Return(productMock, nil).AnyTimes()
	serviceMock.EXPECT().Disable(productId).Return(productMock, nil).AnyTimes()

	createdResultExpected := fmt.Sprintf("product (%s) was created", productName)
	createdResult, err := cli.Run(serviceMock, "create", "", productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, createdResultExpected, createdResult)
	
	enableResultExpected := fmt.Sprintf("product (%s) was enabled", productId)
	enabledResult, err := cli.Run(serviceMock, "enable", productId, productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, enableResultExpected, enabledResult)

	disabledResultExpected := fmt.Sprintf("product (%s) was disabled", productId)
	disableResult, err := cli.Run(serviceMock, "disable", productId, productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, disabledResultExpected, disableResult)
}