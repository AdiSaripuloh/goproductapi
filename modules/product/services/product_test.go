package services

import (
	"github.com/AdiSaripuloh/goproductapi/modules/product/models"
	"testing"
	"time"

	mock_repositories "github.com/AdiSaripuloh/goproductapi/test/mocks/product"
	"github.com/golang/mock/gomock"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestProductService_GetAllProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock_repositories.NewMockProductRepository(ctrl)

	mock.EXPECT().GetAll().Return([]models.Product{
		{ID: uuid.UUID{uuid.V4}, Name: "Product 1", CreatedAt: time.Now()},
		{ID: uuid.UUID{uuid.V4}, Name: "Product 2", CreatedAt: time.Now()},
		{ID: uuid.UUID{uuid.V4}, Name: "Product 3", CreatedAt: time.Now()},
		{ID: uuid.UUID{uuid.V4}, Name: "Product 4", CreatedAt: time.Now()},
	}, nil)

	service := productService{
		repository: mock,
	}

	products, _ := service.GetAll()

	assert.Equal(t, 2, len(products))
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 2", products[1].Name)
	assert.Equal(t, "Product 4", products[2].Name)
}
