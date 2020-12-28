package csvformat

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
)

// Movie conterá nosso CSV analisado
type Movie struct {
	Title    string
	Director string
	Year     int
}

// ReadCSV mostra alguns exemplos de processamento de CSV
// que é passado como um io.Reader
func ReadCSV(b io.Reader) ([]Movie, error) {

	r := csv.NewReader(b)

	//Estas são algumas opções de configuração opcionais
	r.Comma = ';'
	r.Comment = '-'

	var movies []Movie

	//pegue e ignore o cabeçalho por agora
	//também podemos querer usar isso para uma chave de dicionário ou
	//alguma outra forma de pesquisa
	_, err := r.Read()
	if err != nil && err != io.EOF {
		return nil, err
	}

	//loop até que esteja tudo processado
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		year, err := strconv.ParseInt(record[2], 10, 64)
		if err != nil {
			return nil, err
		}

		m := Movie{record[0], record[1], int(year)}
		movies = append(movies, m)
	}
	return movies, nil
}

//AddMoviesFromText usa o analisador CSV com uma string
func AddMoviesFromText() error {
	//este é um exemplo de como pegamos uma string, convertendo
	//em um buffer, e lê-lo
	//com o pacote csv
	in := `
- first our headers
movie title;director;year released
- then some data
Guardians of the Galaxy Vol. 2;James Gunn;2017
Star Wars: Episode VIII;Rian Johnson;2017
`
	b := bytes.NewBufferString(in)
	m, err := ReadCSV(b)
	if err != nil {
		return err
	}

	fmt.Printf("%#v\n", m)
	return nil
}
