package controllers

//Controller passa o estado para nossos handlers
type Controller struct{
	storage Storage
}

//New é um 'construtor' da controller
func New(storage Storage) *Controller {
	return &Controller{
		storage: storage,
	}
}

//Payload é nossa resposta comum
type Payload struct{
	Value string `json:"value"`
}