package inmemory

import (
	"errors"
	"sync"

	"github.com/goproductapi/modules/product/models"
	"github.com/goproductapi/modules/product/repositories"

	uuid "github.com/satori/go.uuid"
)

var (
	productRepoLock sync.Once
	productRepo     repositories.ProductRepository
)

type productRepository struct {
	storage *ProductStorage
}

// NewProductRepositoryInMemory initialize new product repository object
func NewProductRepositoryInMemory(storage *ProductStorage) repositories.ProductRepository {
	productRepoLock.Do(func() {
		productRepo = &productRepository{
			storage: storage,
		}
	})

	return productRepo
}

func (repo *productRepository) GetByUUID(id uuid.UUID) (*models.Product, error) {
	product := models.Product{}
	for _, item := range repo.storage.Products {
		if item.ID == id {
			product = item
			break
		}
	}

	if product == (models.Product{}) {
		return nil, errors.New("product not found")
	}

	return &product, nil
}

func (repo *productRepository) GetAll() ([]models.Product, error) {
	products := []models.Product{}
	for _, item := range repo.storage.Products {
		products = append(products, item)
	}

	return products, nil
}

func (repo *productRepository) Save(product models.Product) error {
	found := false
	for idx, item := range repo.storage.Products {
		if item.ID == product.ID {
			found = true
			repo.storage.Products[idx] = item
			break
		}
	}

	if !found {
		repo.storage.Products = append(repo.storage.Products, product)
	}

	return nil
}
