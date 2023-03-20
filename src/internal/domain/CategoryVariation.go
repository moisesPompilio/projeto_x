package domain

import "github.com/google/uuid"

type CategoryVariation struct {
	Id              uuid.UUID
	Name            string
	Company         uuid.UUID
	IsFloatQuantity bool
}
