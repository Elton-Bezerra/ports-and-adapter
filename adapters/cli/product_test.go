package cli_test

import (
	"fmt"
	"testing"

	"github.com/Elton-Bezerra/ports-and-adapter/adapters/cli"
	mock_app "github.com/Elton-Bezerra/ports-and-adapter/app/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productName := "Product Test"
	productPrice := 25.99
	productStatus := "disabled"
	productId := "abc"

	// mocks product
	productMock := mock_app.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetID().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()

	// mocks service behavior
	service := mock_app.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	resultExpected := fmt.Sprintf(
		"Product ID %s with the name %s has been created with price %f and status %s",
		productId,
		productName,
		productPrice,
		productStatus,
	)

	result, err := cli.Run(service, "create", "", productName, productPrice)

	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product %s has been enabled", productName)

	result, err = cli.Run(service, "enable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product %s has been disabled", productName)

	result, err = cli.Run(service, "disable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product %s has been enabled", productName)

	result, err = cli.Run(service, "enable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product ID: %s\n Product Name: %s\n Product Price: %f\n Product Status: %s",
		productId,
		productName,
		productPrice,
		productStatus,
	)

	result, err = cli.Run(service, "get", "abc", "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)
}
