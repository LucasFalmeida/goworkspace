package main

import "fmt"

type person struct {
	nome      string
	sobrenome string
}

type agenteSecreto struct {
	person
}

func (p agenteSecreto) falar() {
	fmt.Println("Olá, eu sou o ", p.nome)
}

func main() {
	agente := agenteSecreto{
		person: person{
			"Lucas",
			"Feijo",
		},
	}

	agente.falar()

}
