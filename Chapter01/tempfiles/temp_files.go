package tempfiles

import (
	"fmt"
	"io/ioutil"
	"os"
)

//WorkWithTemp dará alguns padrões básicos para trabalhar
//com arquivos e diretórios temporários
func WorkWithTemp() error  {
	//Se você precisar de um local temporário para armazenar arquivos 
	//com o mesmo nome, ou seja. template1-10.html um diretório 
	//temporário é uma boa maneira de abordá-lo, o primeiro argumento 
	//estando em branco significa que ele usará para criar o diretório 
	//no local retornado por os.TempDir()
	t, err := ioutil.TempDir("", "tmp")
	if err != nil {
		return err
	}

	//Isso excluirá tudo dentro do arquivo temporário quando esta 
	//função for encerrada se você quiser fazer isso mais tarde, 
	//certifique-se de retornar o nome do diretório para a f
	//unção de chamada
	defer os.RemoveAll(t)

	//o diretório deve existir para criar o arquivo temporário criado.
	// t é um objeto *os.File.
	tf, err := ioutil.TempFile(t, "tmp")
	if err != nil {
		return err
	}

	fmt.Println(tf.Name())

	//normalmente excluiríamos o arquivo temporário aqui, 
	//mas como o estamos colocando em um diretório temporário, 
	//ele é limpo pelo adiamento anterior
	return nil
}
