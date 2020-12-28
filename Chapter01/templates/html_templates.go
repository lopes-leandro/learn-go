package templates

import (
	"fmt"
	"html/template"
	"os"
)

//HTMLDifferences destaca algumas das 
//diferenças entre html/template e text/template
func HTMLDifferences() error  {
	
	t := template.New("html")
	t, err := t.Parse("<h1>Hello! {{.Name}}</h1>")
	if err != nil {
		return err
	}

	//html/template evita operações inseguras como a
	// injeção de javascript; isto é contextualmente ciente
	// e irá se comportar de forma diferente dependendo
	// de onde uma variável é renderizada
	err = t.Execute(os.Stdout, map[string]string{"Name": "<script>alert('Can you see me?)</script>"})
	if err != nil {
		return err
	}

	//você também pode chamar manualmente os escapers
	fmt.Println(template.JSEscaper(`example <example@example.com>`))
	fmt.Println(template.HTMLEscaper(`example <example@example.com>`))
	fmt.Println(template.URLQueryEscaper(`example <example@example.com>`))

	return nil
}