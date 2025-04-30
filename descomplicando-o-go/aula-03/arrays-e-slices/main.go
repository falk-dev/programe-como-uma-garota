package main

import "fmt"

func main() {
	// Definindo um array
	caracteres := [5]string{"a", "e", "i", "o", "u"}

	fmt.Println(len(caracteres))

	for i := 0; i < len(caracteres); i++ {
		fmt.Println(caracteres[i])
	}

	// Array multidimensional
	matriz_caracteres := [2][3]string{
		{"a", "b", "c"},
		{"d", "e", "f"},
	}

	// Imprimindo array multidimensional
	for i := 0; i < len(matriz_caracteres); i++ {
		for j := 0; j < len(matriz_caracteres[i]); j++ {
			fmt.Println(matriz_caracteres[i][j])
		}
	}

	// Slice - Array variável
	caracteres_slice := []string{"a", "b", "c", "d", "e"}

	caracteres_slice = append(caracteres_slice, "f")

	fmt.Println(caracteres_slice)

	capacidade := 200
	tamanhoDoSlice := 10

	// tamanhoDoSlice é a quantidade de espaços que o Slice vai iniciar logo após a criação
	// capacidade é a quantidade máxima de elementos que o Slice pode ter
	// para adicionar elementos fora do tamanho inicial do Slice, utiliza o append
	caracteres_slice_make := make([]string, tamanhoDoSlice, capacidade)
	caracteres_slice_make = append(caracteres_slice_make, "5")
}
