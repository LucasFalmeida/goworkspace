package main

import (
	"fmt"
)

func main() {
	var resposta1, resposta2, resposta3 string

	fmt.Println("Nome: ")
	_, err := fmt.Scan(&resposta1)

	if err != nil {
		panic(err)
	}

	fmt.Println("Esporte preferido")
	_, err = fmt.Scan(&resposta2)

	if err != nil {
		panic(err)
	}
	fmt.Println("Comida preferida")
	_, err = fmt.Scan(&resposta3)

	if err != nil {
		panic(err)
	}

	fmt.Println(resposta1, resposta2, resposta3)
}
