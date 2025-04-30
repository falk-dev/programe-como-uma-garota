package main

import "fmt"

// Struct não anônima
type Carro struct {
	Marca string
	Ano   int
}

func main() {
	// Struct anônima
	carro := Carro{
		Marca: "Fiat Palio Attractive 1.0",
		Ano:   2014,
	}

	fmt.Println(carro)
	fmt.Println(carro.Marca)
	fmt.Println(carro.Ano)
}
