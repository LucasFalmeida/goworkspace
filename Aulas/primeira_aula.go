package main

import "fmt"

type doce struct {
	nome  string
	valor int
}

func chocolate(nome string, valor int) {
	nome_chocolate := nome
	valor_chocolate := valor
	doce := doce{
		nome_chocolate,
		valor_chocolate,
	}

	fmt.Println("O primeiro produto é:", doce)
	fmt.Println("O nome do chocolate é:", nome_chocolate, "e o preço é:", valor_chocolate)

}

func (b doce) falar() {
	fmt.Println("Eu sou o chocolate!, ", b.nome)

}

func main() {
	fmt.Println("Testando o poder da linguagem GO")

	chocolate("Cacau", 55)
	//instancia de um novo struc
	teste := doce{
		"amendoas", 40,
	}
	//chamei o metodo falar pois a variavel teste é do tipo doce, logo é permitido
	teste.falar()
}
