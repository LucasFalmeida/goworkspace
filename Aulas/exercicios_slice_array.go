package main

import "fmt"

func main()  {
 x:= make([]int, 10, 100)

 x[0] = 1
 x[1] = 2
 x[2] = 3
 x[3] = 4

 fmt.Println(x)
}