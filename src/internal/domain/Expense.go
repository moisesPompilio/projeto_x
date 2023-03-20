package domain

import "github.com/google/uuid"

type Expense struct {
	Id      uuid.UUID
	Value   float64
	Reasonn string
	Company uuid.UUID
	Store   uuid.UUID
}
