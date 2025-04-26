package main

import "fmt"

func soma(num1, num2 int) int {
	return num1 + num2
}

func subtracao(num1, num2 int) int {
	return num1 - num2
}

func multiplicacao(num1, num2 int) int {
	return num1 * num2
}

func divisao(num1, num2 int) int {
	return num1 / num2
}

func main() {
	var num1 int
	var num2 int

	fmt.Scan(&num1)
	fmt.Scan(&num2)

	soma := soma(num1, num2)
	subtracao := subtracao(num1, num2)
	multiplicacao := multiplicacao(num1, num2)
	divisao := divisao(num1, num2)

	fmt.Printf("%d + %d = %d\n", num1, num2, soma)
	fmt.Printf("%d - %d = %d\n", num1, num2, subtracao)
	fmt.Printf("%d * %d = %d\n", num1, num2, multiplicacao)
	fmt.Printf("%d / %d = %d\n", num1, num2, divisao)
}
