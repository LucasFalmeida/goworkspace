package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

type calcular struct {
	Valor int
}

func (c calcular) Retorna() int {
	return c.Valor
}
func (c calcular) Calcular(x int) int {
	return x * 2
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func main() {

	c := calcular{
		Valor: 3,
	}
	fmt.Println(c)
	err := tpl.Execute(os.Stdout, c)
	if err != nil {
		log.Fatalln(err)
	}
}

// Generally speaking, best practice:
// call functions in templates for formatting only; not processing or data access.

// The main reasons you don't want to do any data processing in your template:
// (1) separation of concerns
// (2) if you're using a function more than once in a template,
// the server needs to do the processing more than once.
// (though the standard library might cache processing -
// I've yet to dig into this - if you find the answer, let me know)
