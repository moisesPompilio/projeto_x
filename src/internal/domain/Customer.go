package domain

import "github.com/google/uuid"

type Customer struct {
	Id             uuid.UUID
	Name           string
	Neighborhood   string
	ReferralSource string
	CompanyId      uuid.UUID
}
