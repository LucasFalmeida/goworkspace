package main

import "fmt"

func main() {

	fmt.Println(`
		1-> Soma		
		2-> Subtracao
		3-> Multiplicacao
		4-> Divisao
		
		`)
	fmt.Println("Qual operação deseja realizar?:")
	var ope int
	//adiciona o valor ao ponteiro de ope
	fmt.Scanln(&ope)
	var num int
	var num1 int

	//Procedimiento para la suma

	if ope == 1 {
		fmt.Println("Entre com o primeiro numero:")
		fmt.Scanln(&num)
		fmt.Println("Entre com o segundo numero:")
		fmt.Scanln(&num1)
		resultado := num + num1
		fmt.Println("O resultado é: ", resultado)

	}
	//subtracao
	if ope == 2 {
		fmt.Println("Entre com o primeiro numero:")
		fmt.Scanln(&num)
		fmt.Println("Entre com o segundo numero:")
		fmt.Scanln(&num1)
		resultado := num - num1
		fmt.Println("O resultado é: ", resultado)

	}
	if ope == 3 {
		fmt.Println("Entre com o primeiro numero:")
		fmt.Scanln(&num)
		fmt.Println("Entre com o segundo numero:")
		fmt.Scanln(&num1)
		resultado := num * num1
		fmt.Println("O resultado é: ", resultado)

	}
	if ope == 4 {
		fmt.Println("Entre com o primeiro numero:")
		fmt.Scanln(&num)
		fmt.Println("Entre com o segundo numero:")
		fmt.Scanln(&num1)
		resultado := num / num1
		fmt.Println("O Resultado é: ", resultado)

	}

	fmt.Println("Pressione enter para sair \n")
	var espera string
	fmt.Scanln(&espera)

}
