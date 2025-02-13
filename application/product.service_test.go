package application_test

import (
	"testing"

	"github.com/GuilhermeBeneti1990/go-hexagonal/application"
	mock_application "github.com/GuilhermeBeneti1990/go-hexagonal/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestProductService_Get(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockIProduct(ctrl)
	persistence := mock_application.NewMockIProductPersistence(ctrl)
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Get("123")
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockIProduct(ctrl)
	persistence := mock_application.NewMockIProductPersistence(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}
	result, err := service.Create("Product 01", 10)
	require.Nil(t, err)
	require.Equal(t, product, result)

}

func TestProductService_EnableAndDisable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockIProduct(ctrl)
	product.EXPECT().Enable().Return(nil)
	product.EXPECT().Disable().Return(nil)

	persistence := mock_application.NewMockIProductPersistence(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Enable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)

	result, err = service.Disable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)

}
