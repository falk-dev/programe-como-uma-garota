package main

import "fmt"

func soma(num int, valores ...int) (int, int) {
	for _, valor := range valores {
		num += valor
	}
	return num, len(valores)
}

func main() {
	soma_total, tamanho := soma(10, 5, 5, 5, 5, 5, 5)
	fmt.Println(soma_total)
	fmt.Println(tamanho)
}
