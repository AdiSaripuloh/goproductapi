package dto

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Product dto
type Product struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createAt"`
}

// CreateProductRequest dto
type CreateProductRequest struct {
	Name string `json:"name"`
}
