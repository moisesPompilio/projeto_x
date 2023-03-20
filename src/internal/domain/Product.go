package domain

import "github.com/google/uuid"

type Product struct {
	Id             uuid.UUID
	Nome           string
	Image          string
	Cost           float64
	SuggestedPrice float64
	Company        uuid.UUID
}
