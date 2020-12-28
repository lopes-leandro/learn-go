package bytestrings

import (
	"fmt"
	"io"
	"os"
	"strings"
)

//SearchString mostra vários métodos
//para pesquisar uma string
func SearchString()  {
	
	s := "this is a test"

	//retorna verdadeiro porque s contém
	//a palavra this
	fmt.Println(strings.Contains(s, "this"))

	//retorna verdadeiro porque s contém a letra a
	//também corresponderia se contivesse b ou c
	fmt.Println(strings.ContainsAny(s, "abc"))

	//retorna verdadeiro porque s começa com this
	fmt.Println(strings.HasPrefix(s, "this"))

	//retorna verdadeiro porque s termina com this
	fmt.Println(strings.HasSuffix(s, "test"))
}

//ModifyString modifica uma string de várias maneiras
func ModifyString()  {
	
	s := "simple string"

	//imprime [string simples]
	fmt.Println(strings.Split(s, " "))

	//imprime [Simple String]
	fmt.Println(strings.Title(s))

	//imprime "simple string"; todos arrastando e
	//o espaço em branco à esquerda é removido
	s = "  simple string "
	fmt.Println(strings.TrimSpace(s))
}

//StringReader demonstra como criar
//uma interface io.Reader rapidamente com uma string
func StringReader()  {
	s := "simple string\n"
	r := strings.NewReader(s)

	//imprime s na Stdout
	io.Copy(os.Stdout, r)
}