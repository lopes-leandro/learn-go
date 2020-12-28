package bytestrings

import (
	"bufio"
	"bytes"
	"fmt"
)

//WorkWithBuffer fará uso do buffer criado pela
//função buffer
func WorkWithBuffer() error {
	rawString := "it's easy to encode unicode into a byte array"

	b := Buffer(rawString)

	//podemos converter rapidamente um buffer em byes com
	//b.Bytes () ou uma string com b.String()
	fmt.Println(b.String())

	//porque este é um leitor io, podemos fazer uso de
	//funções genéricas do leitor io, como
	s, err := toString(b)
	if err != nil {
		return err
	}
	fmt.Println(s)

	//também podemos pegar nossos bytes e criar um leitor de bytes
	//esses leitores implementam io.Reader, io.ReaderAt,
	//io.WriterTo, io.Seeker, io.ByteScanner e
	//Interfaces io.RuneScanner
	reader := bytes.NewReader([]byte(rawString))

	//também podemos conectá-lo a um scanner que permite
	//leitura em buffer e tokenização
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	//iterar sobre todos os eventos de verificação
	for scanner.Scan() {
		fmt.Print(scanner.Text())
	}

	return nil
}
