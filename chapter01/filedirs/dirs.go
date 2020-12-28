package filedirs

import (
	"errors"
	"io"
	"os"
)

//Operate manipulador de arquivos e diretórios
func Operate() error {
	//este 0755 é semelhante ao que você veria com Chown
	//em uma linha de comando, isso criará um diretor /tmp/example, 
	//você também pode usar um caminho absoluto
	//em vez de um relativo
	if err := os.Mkdir("example_dir", os.FileMode(0755));
	err != nil {
		return err
	}

	//vá para o diretório / tmp
	if err := os.Chdir("example_dir"); err != nil {
		return err
	}

	//f é um objeto de arquivo genérico
	//ele também implementa várias interfaces
	//e pode ser usado como reader ou writer
	//se os bits corretos são definidos ao abrir
	f, err := os.Create("test.txt")
	if err != nil {
		return err
	}

	//escrevemos um valor de comprimento conhecido para o arquivo e
	//valide que escreveu corretamente
	value := []byte("hello\n")
	count, err :=f.Write(value)
	if err != nil {
		return err
	}
	if count != len(value) {
		return errors.New("incorrect length returned from write")
	}
	if err := f.Close(); err != nil {
		return err
	}

	//leia o arquivo
	f, err = os.Open("test.txt")
	if err != nil {
		return err
	}

	io.Copy(os.Stdout, f)

	if err := f.Close(); err != nil {
		return err
	}

	// vá para o diretório /tmp
	if err := os.Chdir(".."); err != nil {
		return err
	}

	//limpeza, os.RemoveAll pode ser perigoso se você
	//apontar para o diretório errado, usar a entrada do usuário,
	//e especialmente se você executar como root
	if err := os.RemoveAll("example_dir"); err != nil {
		return err
	}

	return nil
}