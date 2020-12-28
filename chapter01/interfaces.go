package interfaces

import (
	"fmt"
	"io"
	"os"
)

// Copy: copia os dados de entrada para saída primeiro diretamente e, 
// em seguida, usando um buffer. Ele também escreve para stdout
func Copy(in io.ReadSeeker, out io.Writer) error {
	//escrevemos para fora, mas também Stdout
	w := io.MultiWriter(out, os.Stdout)

	//uma cópia padrão, isso pode ser perigoso 
	//se houver muitos dados em
	if _, err := io.Copy(w, in); err != nil {
		return err
	}

	in.Seek(0,0)

	//gravação em buffer usando blocos de 64 bytes
	buff := make([]byte, 64)
	if _, err := io.CopyBuffer(w, in, buff); err != nil {
		return err
	}

	//vamos imprimir uma nova linha
	fmt.Println()

	return nil
}