package negociate

import (
	"io"
	"github.com/unrolled/render"
)

//Respond ativa o tipo de conteúdo
// para determinar a resposta
func (n *Negotiator) Respond(w io.Writer, status int, v interface{})  {
	
	switch n.ContentType {
	case render.ContentJSON:
		n.Render.JSON(w, status, v)
	case render.ContentXML:
		n.Render.XML(w, status, v)
	default:
		n.Render.JSON(w, status, v)
		
	}
}