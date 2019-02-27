package main

import "fmt"

func main() {

	defer foo()
	bar()
	test()

}

func foo() {
	fmt.Println("Foo")
}
func bar() {
	fmt.Println("bar")
}
func test() {
	fmt.Println("test")
}
