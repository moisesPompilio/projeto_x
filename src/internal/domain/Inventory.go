package domain

import "github.com/google/uuid"

type Inventory struct {
	Store            uuid.UUID
	ProductVariation uuid.UUID
	Quantity         float64
}
