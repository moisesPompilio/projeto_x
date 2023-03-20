package domain

import "github.com/google/uuid"

type Store struct {
	Id        uuid.UUID
	Nome      string
	CNPJ      string
	Inventory uuid.UUID
	Company   uuid.UUID
}
