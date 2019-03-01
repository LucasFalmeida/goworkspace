package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Create("nomes.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	r := strings.NewReader("Coé")

	io.Copy(f, r)
}
