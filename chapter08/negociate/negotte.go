package negociate

import (
	"net/http"

	"github.com/unrolled/render"
)

//Negotiator encerra a renderização e
// faz algumas mudanças no ContentType
type Negotiator struct {
	ContentType string
	*render.Render
}

//GetNegotiator pega uma requisição e
// descobre o tipo de conteúdo no
// cabeçalho Content-Type
func GetNegotiator(req *http.Request) *Negotiator {

	contentType := req.Header.Get("Content-Type")
	return &Negotiator{
		ContentType: contentType,
		Render:      render.New(),
	}

}
