package main

import "fmt"

func main() {
	x := 33

	for x <= 122{
		x++
		fmt.Println("%V\t%#U\n", x, x)

	}
}

