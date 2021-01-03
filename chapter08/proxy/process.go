package proxy

import(
	"bytes"
	"net/http"
	"net/url"
)

//ProcessRequest modifica a requisição
// de acordo com as configurações de proxy
func (p *Proxy) ProcessRequest(req *http.Request) error {
	
	proxyURLRaw := p.BaseURL + req.URL.String()

	proxyURL, err := url.Parse(proxyURLRaw)
	if err != nil {
		return err
	}

	req.URL = proxyURL
	req.Host = proxyURL.Host
	req.RequestURI = ""
	return nil
}

//CopyResponse pega a resposta do cliente
// e grava tudo no ResponseWriter no handler original
func CopyRespose(w http.ResponseWriter, res *http.Response)  {
	
	var out bytes.Buffer
	out.ReadFrom(res.Body)

	for key, values := range res.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	w.WriteHeader(res.StatusCode)
	w.Write(out.Bytes())
	
}