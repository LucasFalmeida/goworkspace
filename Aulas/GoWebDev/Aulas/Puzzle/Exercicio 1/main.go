package main

import "fmt"

//falta somar valor total dos livros unicos com desconto + total de livros, pra isso, subtrai o numero de livros unicos pelo total de livors (3 livros unicos e são 5 livros, logo tem 2 livros repetidos nao tem desconto, 2*42)
func main() {

	//quantidade de livros 1, 2 ,3 ,4 e 5
	livros := []int{2, 2, 2, 1, 1}
	var livros_unicos int
	var total int
	var total_livros int

	total_livros = len(livros)
	fmt.Println(total_livros)
	//passa pelo valor de cada indice do slice, e ve se o livro foi comprado mais de 2x, se não foi, aplica uma soma
	for _, v := range livros {
		if v <= 1 {
			livros_unicos = v + livros_unicos
		}
	}
	//livros_unicos porque é numero de livros unicos
	if livros_unicos == 2 {
		a := (5 / 100) * total
		fmt.Println(a)
	} else if livros_unicos == 3 {
		a := total * (10 / 100)
		fmt.Println(a)
	} else if livros_unicos == 4 {
		a := (15 / 100) * total
		fmt.Println(a)
	} else if livros_unicos == 5 {
		a := total * (10 / 100)
		fmt.Println(a)
	} else {
		fmt.Println(total)
	}

	for _, t := range livros {
		total = (t + total)
	}
	total = total * 42

	fmt.Println("total dos livros deu:", total, "reais")

}

/*


Este problema foi utilizado em 118 Dojo(s).

Uma livraria contém 7 títulos distintos e possui um esquema de descontos que leva em consideração se é o mesmo título ou não. O problema consiste em calcular o melhor desconto para o cliente.

Preço de um título qualquer: R$ 42,00

Descontos:
        2 livros - 5%
        3 livros - 10%
        4 livros - 15%
        5 livros - 20%

Quanto custa?
        2 cópias do primeiro livro
        2 cópias do segundo livro
        2 cópias do terceiro livro
        1 cópia do quarto livro
        1 cópia do quinto livro

Resposta: R$ 268,80
        Note que o preço deve ser o menor valor obtido combinando os livros.
		2 conjuntos de 4 livros é melhor que 1 de 5 livros e 1 de 3 livros.

*/
