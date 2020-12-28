package interfaces

import (
	"io"
	"os"
)

//PipeExample: ajuda a dar mais alguns exemplos de uso de interfaces io
func PipeExample() error  {
	//o leitor de pipe e o escritor de pipe implementam 
	//io.Reader e io.Writer
	r, w := io.Pipe()

	//isso precisa ser executado em uma rotina do go separada, 
	//pois irá bloquear a espera pelo 
	//fechamento do leitor no final para limpeza
	go func ()  {
		//por enquanto vamos escrever algo básico, 
		//isso também pode ser usado para codificar 
		//codificação json base64, etc.
		w.Write([]byte("test\n"))
		w.Close()
	}()

	if _, err := io.Copy(os.Stdout, r); err != nil {
		return err
	}

	return nil
}