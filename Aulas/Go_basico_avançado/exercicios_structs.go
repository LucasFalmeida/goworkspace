package main

import "fmt"

var c int

func main() {
	objeto := []int{1, 3, 4, 5, 6, 7, 8}
	fmt.Println(objeto)
	c := falar(objeto...)
	fmt.Println(c)

}

func falar(a ...int) int {
	c := 0
	for _, v := range a {
		c += v
	}
	return c

}
