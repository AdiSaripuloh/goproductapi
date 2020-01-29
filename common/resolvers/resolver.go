package resolvers

import (
	inmemoryProductRepositories "github.com/AdiSaripuloh/goproductapi/modules/product/repositories/inmemory"
	productServices "github.com/AdiSaripuloh/goproductapi/modules/product/services"
)

// Resolver struct
type Resolver struct {
	ProductService productServices.ProductService
}

// NewResolver initialize new resolvers object
func NewProductResolver() *Resolver {
	// storage
	productStorage := inmemoryProductRepositories.NewProductStorage()

	// repositories
	productRepository := inmemoryProductRepositories.NewProductRepositoryInMemory(productStorage)

	// service
	productService := productServices.NewProductService(productRepository)

	return &Resolver{
		ProductService: productService,
	}
}
