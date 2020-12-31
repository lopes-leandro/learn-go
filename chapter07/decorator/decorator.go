package decorator

import (
	"net/http"
)

//TransportFunc implementa a interface RountTripper
type TransportFunc func (*http.Request) (*http.Response, error)

//RoundTrip apenas chama a função original
func (tf TransportFunc) RoundTrip(r * http.Request) (*http.Response, error)  {
	return tf(r)
}

//Decorator é uma função conveniente para representar
// nossa função interna de middleware
type Decorator func (http.RoundTripper) http.RoundTripper

//Decorate é um ajudante para embrulhar todo o middleware
func Decorate(t http.RoundTripper, rts ...Decorator) http.RoundTripper {
	decorated := t
	for _, rt := range rts {
		decorated = rt(decorated)
	}
	return decorated
}