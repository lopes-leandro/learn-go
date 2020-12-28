package csvformat

import (
	"bytes"
	"encoding/csv"
	"io"
	"os"
)

//Book tem um autor e um título
type Book struct {
	Author string
	Title string
}

//Books é um tipo nomeado para uma variedade de livros
type Books []Book

//ToCSV pega um conjunto de livros e grava em um io.Writer
//retorna quaisquer erros
func (books *Books) ToCSV(w io.Writer) error  {
	n := csv.NewWriter(w)
	err := n.Write([]string{"Author", "Title"})
	if err != nil {
		return err
	}

	for _, book := range *books {
		err := n.Write([] string{book.Author, book.Title})
		if err != nil {
			return err
		}
	}

	n.Flush()
	return n.Error()
}

//WriteCSVOutput inicializa um conjunto de livros
//e os grava em os.Stdout
func WriteCSVOutput() error {
	b := Books{
		Book{
			Author: "F Scott Fitzgerald",
			Title: "The Great Gatsby",
		},
		Book{
			Author: "J D Salinger",
			Title: "The Catcher in the Rye",
		},
	}
	return b.ToCSV(os.Stdout)
}

//WriteCSVBuffer retorna um buffer CSV para
//um conjunto de livros
func WriteCSVBuffer() (*bytes.Buffer, error)  {
	b := Books{
		Book{
			Author: "F Scott Fitzgerald",
			Title: "The Great Gatzby",
		},
		Book{
			Author: "J D Salinger",
			Title: "The Catcher in the Rye",
		},
	}

	w := &bytes.Buffer{}
	err := b.ToCSV(w)
	return w, err
}