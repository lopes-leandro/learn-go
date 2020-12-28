package filedirs

import (
	"bytes"
	"io"
	"os"
	"strings"
)

// Capitalizer abre um arquivo, lê o conteúdo,
// em seguida, grava esse conteúdo em um segundo arquivo
func Capitalizer(f1 *os.File, f2 *os.File) error {
	if _, err := f1.Seek(0, io.SeekStart); err != nil {
		return err
	}

	var tmp = new(bytes.Buffer)

	if _, err := io.Copy(tmp, f1); err != nil {
		return err
	}

	s :=strings.ToUpper(tmp.String())

	if _, err := io.Copy(f2, strings.NewReader(s)); err != nil {
		return err
	}

	return nil
}

// CapitalizerExample cria dois arquivos, grava em um
// então chama Capitalizer() em ambos
func CapitalizerExample() error  {
	f1, err := os.Create("file1.txt")
	if err != nil {
		return err
	}

	if _, err := f1.Write([]byte("this file contains a number of words and new lines")); err != nil {
		return err
	}

	f2, err := os.Create("file2.txt")
	if err != nil {
		return err
	}

	if err := Capitalizer(f1, f2); err != nil {
		return err
	}

	if err := os.Remove("file1.txt"); err != nil {
		return err
	}

	if err := os.Remove("file2.txt"); err != nil {
		return err
	}

	return nil
}