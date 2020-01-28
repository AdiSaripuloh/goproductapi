package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Product data model
type Product struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
}

// NewProduct initialize new instance of Product
func NewProduct(name string) *Product {
	return &Product{
		ID:        uuid.NewV4(),
		Name:      name,
		CreatedAt: time.Now().UTC(),
	}
}
