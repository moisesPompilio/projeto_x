package testunit_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	validation "github.com/moisesPompilio/projeto_x/src/pkg/helps/Validation"
)

type TestStructure struct {
	Name     string
	Age      int
	Email    string
	Address  *Address
	Language interface{}
}

type Address struct {
	Street  string
	Number  int
	ZipCode string
	Country string
}

func TestValidateRequiredFields_Valid(t *testing.T) {
	AddressTest := &Address{
		Street:  "Street",
		Number:  1,
		ZipCode: "ZipCode",
		Country: "Country",
	}

	objTest := TestStructure{Email: "test@example.com", Name: "John Doe", Language: "English", Address: AddressTest}

	requiredFields := []string{"Email", "Name", "Language", "Address"}

	err := validation.ValidateRequiredFields_Interfaces_Strings_Pointer(objTest, requiredFields)

	assert.NoError(t, err)
}

func TestValidateRequiredFields_Invalid_MissingRequiredField(t *testing.T) {

	objTest := TestStructure{Email: "", Name: "John Doe", Language: "English"}
	requiredFields := []string{"Email", "Name", "Password", "NickName", "Object"}

	err := validation.ValidateRequiredFields_Interfaces_Strings_Pointer(objTest, requiredFields)

	assert.Error(t, err)
	assert.EqualError(t, err, "field \"Email\" cannot be empty")
}

func TestValidateRequiredFields_Invalid_ObjectIsNil(t *testing.T) {
	requiredFields := []string{"Email", "Name", "Password", "NickName", "Object"}
	err := validation.ValidateRequiredFields_Interfaces_Strings_Pointer(nil, requiredFields)

	assert.Error(t, err)
	assert.EqualError(t, err, "object is nil")
}

func TestValidateRequiredFields_Invalid_FieldTypeIsInterfaceAndValueIsNil(t *testing.T) {
	objTest := TestStructure{Email: "test@example.com", Name: "", Language: nil}
	requiredFields := []string{"Email", "Name", "Language"}

	err := validation.ValidateRequiredFields_Interfaces_Strings_Pointer(objTest, requiredFields)

	assert.Error(t, err)
	assert.EqualError(t, err, "field \"Name\" cannot be empty")
}

func TestValidateRequiredFields_Invalid_FieldTypePointIsNil(t *testing.T) {
	objTest := TestStructure{Email: "test@example.com", Name: "1234", Language: "English", Address: nil}
	requiredFields := []string{"Email", "Name", "Language", "Address"}

	err := validation.ValidateRequiredFields_Interfaces_Strings_Pointer(objTest, requiredFields)

	assert.Error(t, err)
	assert.EqualError(t, err, "field \"Address\" cannot be nil")
}

func TestValidateRequiredFields_Invalid_FieldDosNotExist(t *testing.T) {
	objTest := TestStructure{Email: "test@example.com", Name: "1234", Language: "English"}
	requiredFields := []string{"Email", "Name", "Language", "Object"}

	err := validation.ValidateRequiredFields_Interfaces_Strings_Pointer(objTest, requiredFields)

	assert.Error(t, err)
	assert.EqualError(t, err, "field \"Object\" does not exist")
}
