package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)

	go fanOutIn(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

func populate(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func fanOutIn(c1, c2 chan int) {
	//criou um grupo
	var wg sync.WaitGroup
	//entra o proximo valor
	for v := range c1 {
		//adicionou um na fila
		wg.Add(1)
		go func(v2 int) {
			//canal 2 executa a função timeConsumingWork e passa como parametro esperado (int) o valor de v2, que é o valor de V, que é o valor entre 0 e 100
			c2 <- timeConsumingWork(v2)
			//finalizou
			wg.Done()
			//entrou o proximo valor de v, que é o valor de 0 á 100
		}(v)
	}
	//espera antes de matar
	wg.Wait()
	//após finalizar tudo, ele fecha o channel 2
	close(c2)
}

func timeConsumingWork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}
