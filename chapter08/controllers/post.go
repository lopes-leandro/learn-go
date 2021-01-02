package controllers

import (
	"encoding/json"
	"net/http"
)

//SetValue modifica o armazenamento
// subjacente do objeto controller
func (c *Controller) SetValue(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	if err := req.ParseForm(); err != nil {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	value := req.FormValue("value")
	c.storage.Put(value)
	w.WriteHeader(http.StatusOK)

	p := Payload{Value: value}
	if payload, err := json.Marshal(p); err == nil {
		w.Write(payload)
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

}
