package inmemory

import "github.com/goproductapi/modules/product/models"

// ProductStorage to store in store product in memory
type ProductStorage struct {
	Products []models.Product
}

// NewProductStorage initialize new product storage object
func NewProductStorage() *ProductStorage {
	return &ProductStorage{
		Products: []models.Product{},
	}
}
