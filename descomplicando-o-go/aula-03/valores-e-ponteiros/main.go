package main

import "fmt"

func main() {
	var x int
	x = 10
	var b bool
	b = false

	pointer0fx := &x
	pointer0fb := &b

	fmt.Println(x)
	fmt.Println(b)
	fmt.Println(pointer0fx)
	fmt.Println(pointer0fb)
}
