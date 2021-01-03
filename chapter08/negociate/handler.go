package negociate

import (
	"encoding/xml"
	"net/http"
)

//Payload define seu layout em xml e json
type Payload struct{
	XMLName xml.Name `xml:"payload" json:"-"`
	Status string `xml:"status" json:"status"`
}

//Handler obtém um negociador usando a
// requisição e, em seguida, renderiza um payload
func Handler(w http.ResponseWriter, req *http.Request)  {
	
	n := GetNegotiator(req)
	n.Respond(w, http.StatusOK, &Payload{Status: "Successful!"})
	
}