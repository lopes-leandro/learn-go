package validation

//Controller mantém nossas funções de validação
type Controller struct{
	ValidadePayload func (p *Payload) error
}

//New inicializa a controller com nossa
// validação local, ele pode ser substituído
func New() *Controller  {
	return &Controller{
		ValidadePayload: ValidatePayload,
	}
}