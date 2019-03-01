package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tmath.html"))
}

var fm = template.FuncMap{
	"multiplicar": math,
}

func math(a int, b int) int {
	c := a * b
	return c
}

func main() {
	a := 2
	b := 3
	dados := []int{a, b}
	err := tpl.ExecuteTemplate(os.Stdout, "tmath.html", dados)
	if err != nil {
		log.Fatalln(err)
	}
}
