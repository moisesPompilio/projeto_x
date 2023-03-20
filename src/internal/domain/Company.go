package domain

import (
	"github.com/google/uuid"

	"github.com/moisesPompilio/projeto_x/src/internal/interfaces/input"
	validation "github.com/moisesPompilio/projeto_x/src/pkg/helps/Validation"
)

type Company struct {
	Id   uuid.UUID
	Name string
	CNPJ string
}

func NewCompany(input input.CreateCompanyDTO) (company Company, err error) {

	company = Company{
		Id:   uuid.New(),
		Name: input.Nome,
		CNPJ: input.CNPJ,
	}
	err = company.Valid()
	if err != nil {
		return Company{}, err
	}

	return
}

func (company *Company) Valid() error {
	requiredFields := []string{"Name", "CNPJ"}
	err := validation.ValidateRequiredFields_Interfaces_Strings_Pointer(*company, requiredFields)
	return err
}
