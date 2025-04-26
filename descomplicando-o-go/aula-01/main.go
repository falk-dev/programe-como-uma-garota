// A linguagem é organizada em pacotes
package main

// Pacote de entrada e saída
import "fmt"

// A convenção de Go é camelCase
// Quando o nome da função está com a primeira letra minúscula, significa que ela é privada dentro do pacote
func dummyFunction() {}

// Quando o nome da função está com a primeira letra maiúscula, significa que ela é pública dentro do código todo
func DummyFunction() {}

// Função principal
func main() {
	// Atribuição curta de variável (declaração + atribuição)
	// A presença de ':' quer dizer que é uma string
	s := "Hello World"

	// Para declarar e depois utilizar, fica assim:
	var u string
	u = "Hello World 2.0"

	// Saída
	fmt.Println(s)
	fmt.Println(u)

	// Entrada
	var entrada int
	fmt.Scan(&entrada)

	// Desvios Condicionais
	// If - Else if - Else
	if entrada == 1 {
		fmt.Println("É 1.")
	} else if entrada == 2 {
		fmt.Println("É 2.")
	} else {
		fmt.Println("Não é 1.")
	}

	// Switch - Case
	switch entrada {
	case 1:
		fmt.Println("É 1.")
	case 2:
		fmt.Println("É 2.")
	default:
		fmt.Println("Não é 1.")
	}

	// Laços de repetição
	// For
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// Em Go, não há While. Uma alternativa que funciona de forma semelhante é:
	running := true
	for running {
		// Faça algo
	}

}
