package domain

import "github.com/google/uuid"

type ProductVariation struct {
	Id                uuid.UUID
	ProductId         uuid.UUID
	CategoryVariation uuid.UUID
	Nome              string
}
