package validation

import (
	"fmt"
	"reflect"
)

func ValidateRequiredFields_Interfaces_Strings_Pointer(obj interface{}, requiredFields []string) error {
	if obj == nil {
		return fmt.Errorf("object is nil")
	}

	v := reflect.ValueOf(obj)

	for _, field := range requiredFields {
		// Verificar se o campo existe
		if f := v.FieldByName(field); !f.IsValid() {
			return fmt.Errorf("field %q does not exist", field)
		} else {
			// Verificar se o campo é vazio ou nulo
			switch f.Kind() {
			case reflect.String:
				if f.String() == "" {
					return fmt.Errorf("field %q cannot be empty", field)
				}
			case reflect.Interface, reflect.Ptr:
				if f.IsNil() {
					return fmt.Errorf("field %q cannot be nil", field)
				}
			default:
				// Nenhum tratamento específico para outros tipos de dados
			}
		}
	}

	return nil
}
