package main

import (
	"fmt"
	"strconv"
)

func lotteryCoupons(n int32) int32 {
	//atribui valor dos numeros num array
	counts := map[int32]int32{}
	//variavel que conta o numero de lotes 1,2,3
	maxCount := int32(0)
	//corre por dentro do array, pegando o valor de cada index
	for i := 1; i <= int(n); i++ {
		//converte o valor para string, para pegar cada valor, e somar seu respectivo numero (ex: 11 = 1 + 1 = 2)
		m := strconv.Itoa(i)
		//varial que acumula a soma dos dois numeros
		var total int32 = 0
		//faz um range das strings do map (valores numericos tipados para strings para split)
		for _, c := range m {
			//converte C para string, para ai sim converter para int (Atoi)
			digit, _ := strconv.Atoi(string(c))
			//total recebe o valor de cada item
			total += int32(digit)

		}
		// pega o valor total dentro do map criado e adiciona 1
		counts[total] += int32(1)
		//de 0 a 9, lembrando que só pode ser 3 numeros maximos
		//o total (de 0 a 9) - 10 11 12 13 (1+0 = 1. 1+1 = 2...)

		//o valor de total for maior que maxcount (0), maxcount recebe o valor do total
		if counts[total] > maxCount {
			maxCount = counts[total]

		}

	}

	queroPassar := int32(0)
	//range dentro do map que contem o valor da soma de cada indice do bilhete, pega o valor (resultado ex= 1+1 = >>>>2<<<,), se o valor for igual ao valor contido em maxcount (1,2 ou 3( grupos de pessoas com mesmo ticket)), acrescenta esse valor na variavel quero passar (s), que retorna o grupo de maior numero de pessoas com o ticket vencedor
	for _, v := range counts {
		if v == maxCount {
			queroPassar++

		}

	}
	//retorna o int esperado lá em cima, que é o valor (1,2,3) vencedor
	return queroPassar

}

func main() {
	fmt.Println(lotteryCoupons(26))
}
