package main

import "fmt"

func main() {
	n := 5
	inc(&n)
	fmt.Println(n)
}

func inc(x *int) {
	*x++
}

func incWithReturn() *int {
	n := 1
	return &n 
}
