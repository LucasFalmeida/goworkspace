package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

type calc struct {
	a int
	b int
}

// create a FuncMap to register functions.
// "uc" is what the func will be called in the template
// "uc" is the ToUpper func from package strings
// "ft" is a func I declared
// "ft" slices a string, returning the first three characters
var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
	"m":  math,
}

func math(a int, b int) int {
	c := a * b
	return c
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("template_func.html"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 3 {
		s = s[:3]
	}
	return s
}

func main() {
	a := calc{
		a: 2, b: 3,
	}

	b := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	g := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	m := sage{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases with hatred but with love alone is healed.",
	}

	f := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        2,
	}

	c := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}
	fmt.Println(a)
	sages := []sage{b, g, m}
	cars := []car{f, c}
	numeros := []calc{a, aa}
	//struct anonimo
	data := struct {
		Wisdom      []sage
		Transport   []car
		Calculadora []calc
	}{
		sages,
		cars,
		numeros,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "template_func.html", data)
	if err != nil {
		log.Fatalln(err)
	}
}
