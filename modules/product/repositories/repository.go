package repositories

import (
	"github.com/AdiSaripuloh/goproductapi/modules/product/models"

	uuid "github.com/satori/go.uuid"
)

// ProductRepository interface
type ProductRepository interface {
	GetByUUID(id uuid.UUID) (*models.Product, error)
	GetAll() ([]models.Product, error)
	Save(product models.Product) error
}
