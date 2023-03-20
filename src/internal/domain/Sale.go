package domain

import "github.com/google/uuid"

type Sale struct {
	Id               uuid.UUID
	Store            uuid.UUID
	Customer         uuid.UUID
	ProductVariation uuid.UUID
	Quantity         float64
	Value            float64
}
