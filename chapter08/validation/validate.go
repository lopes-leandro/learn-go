package validation

import (
	"errors"
)

//Verror é um erro que ocorre durante
// a validação, podemos devolvê-lo a
// um usuário
type Verror struct{
	error
}

//Payload é o valor que processamos
type Payload struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

//ValidatePayload é uma implementação
// do closure em nossa controller
func ValidatePayload(p *Payload) error {
	if p.Name == "" {
		return Verror{errors.New("name is required")}
	}
	if p.Age <= 0 || p.Age >= 120 {
		return Verror{errors.New("age is required and must be a value greater than 0 and less than 120")}
	}
	return nil
}