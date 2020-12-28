package bytestrings

import (
	"bytes"
	"io"
	"io/ioutil"
)

//Buffer demonstra alguns truques para inicializar bytes
//Buffers
//Esses buffers implementam uma interface io.Reader
func Buffer(rawString string) *bytes.Buffer {

	//vamos começar com uma string codificada em bytes brutos
	rawBytes := []byte(rawString)

	//existem várias maneiras de criar um buffer de
	//bytes brutos ou da string original
	var b = new(bytes.Buffer)
	b.Write(rawBytes)

	//alternativamente
	b = bytes.NewBuffer(rawBytes)

	//e evitando o array de bytes inicial completamente
	b = bytes.NewBufferString(rawString)

	return b
}

//toString é um exemplo de como pegar um io.Reader e consumir
//tudo, então retornando uma string
func toString(r io.Reader) (string, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
