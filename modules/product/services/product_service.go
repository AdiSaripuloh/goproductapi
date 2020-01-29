package services

import (
	"errors"
	"sync"

	"github.com/AdiSaripuloh/goproductapi/modules/product/dto"
	"github.com/AdiSaripuloh/goproductapi/modules/product/models"
	"github.com/AdiSaripuloh/goproductapi/modules/product/repositories"

	"github.com/jinzhu/copier"
	uuid "github.com/satori/go.uuid"
)

var (
	productSvcLock sync.Once
	productSvc     ProductService
)

// ProductService interface
type ProductService interface {
	GetByUUID(id uuid.UUID) (*dto.Product, error)
	Create(product dto.CreateProductRequest) (*dto.Product, error)
	GetAll() ([]dto.Product, error)
}

type productService struct {
	repository repositories.ProductRepository
}

// NewProductService initialize new product service
func NewProductService(repository repositories.ProductRepository) ProductService {
	productSvcLock.Do(func() {
		productSvc = &productService{
			repository: repository,
		}
	})

	return productSvc
}

func (svc productService) GetByUUID(id uuid.UUID) (*dto.Product, error) {
	if id == uuid.Nil {
		return nil, errors.New("invalid id")
	}

	model, err := svc.repository.GetByUUID(id)
	if err != nil {
		return nil, err
	}

	var result dto.Product
	if err := copier.Copy(&result, model); err != nil {
		return nil, err
	}

	return &result, nil
}

func (svc productService) Create(product dto.CreateProductRequest) (*dto.Product, error) {
	model := models.NewProduct(product.Name)
	if err := svc.repository.Save(*model); err != nil {
		return nil, err
	}

	var result dto.Product
	if err := copier.Copy(&result, model); err != nil {
		return nil, err
	}

	return &result, nil
}

func (svc productService) GetAll() ([]dto.Product, error) {
	model, err := svc.repository.GetAll()
	if err != nil {
		return nil, err
	}

	var results []dto.Product
	if err := copier.Copy(&results, model); err != nil {
		return nil, err
	}

	return results, nil
}
