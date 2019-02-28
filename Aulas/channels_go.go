package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	//envia
	go send(c)
	//recebe
	receive(c)

	fmt.Println("acabou")
}

// envia channel
func send(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

// recebe channel
func receive(c <-chan int) {
	for v := range c {
		fmt.Println("Valor recebido no channel:", v)
	}
}
