package validation

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//Process é um handler que valida um post de Payload
func (c *Controller) Process(w http.ResponseWriter, req *http.Request) {
	
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(req.Body)

	defer req.Body.Close()

	var p Payload

	if err := decoder.Decode(&p); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := c.ValidadePayload(&p); err != nil {
		switch err.(type) {
		case Verror:
			w.WriteHeader(http.StatusBadRequest)
			//passe o Verror adiante
			w.Write([]byte(err.Error()))
			return
		default:
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}