package templates

import (
	"os"
	"strings"
	"text/template"
)

const sampleTemplate = `
	This template demonstrates printing a {{.Variable | printf "%#v"}}.
	{{if .Condition}}
	If condition is set, we'll print this
	{{else}}
	Otherwise, we'll print this instead
	{{end}}
	
	Next we'll interate over an array of strings:
	{{range $index, $item := .Items}}
	{{$index}}: {{$item}}
	{{end}}
	
	We can also easily import other functions like
	strings.Split
	then immediately used the array created as a result:
	{{range $index, $item := split .Words " , "}}
	{{$index}}: {{$item}}
	{{end}}
	
	Blocks are a way to embed templates into one another
	{{block "block_example" .}}
	No Block defined!
	{{end}}
	
	{{/*
		This is a way
		to insert a multi-line comment
		*/}}`

const secondTemplate = `
	{{ define "block_example" }}
	{{ .OtherVariable }}
	{{end}}`

	//RunTemplate inicializa um template e demonstra uma 
	//variedade de funções helpers de template
	func RunTemplate() error  {
		data := struct {
			Condition bool
			Variable string
			Items []string
			Words string
			OtherVariable string
		}{
			Condition: true,
			Variable: "variable",
			Items: []string{"item1", "item2", "item3"},
			Words: "another_item1,another_item2,another_item3",
			OtherVariable: "I'm defined in a second tempate!",
		}

		funcmap := template.FuncMap{
			"split": strings.Split,
		}

		//estes também podem ser acorrentados
		t := template.New("example")
		t = t.Funcs(funcmap)

		//Em vez disso, poderíamos usar Must para entrar em pânico 
		//com o erro template.Must (t.Parse (sampleTemplate))
		t, err := t.Parse(sampleTemplate)
		if err != nil {
			return err
		}

		//para demonstrar os blocos, criaremos outro templates, clonando 
		//o primeiro template e, em seguida, analisando um segundo
		t2, err := t.Clone()
		if err != nil {
			return err
		}

		t2, err = t2.Parse(secondTemplate)
		if err != nil {
			return err
		}

		//escreva o template para stdout e 
		//preencha-o com dados
		err = t2.Execute(os.Stdout, &data)
		if err != nil {
			return err
		}

		return nil
	}