package main

import (
	"fmt"
	"os"
)

func main() {
	// Múltiplos retornos
	file, err := os.Open("arquivo.txt")

	// Tratamento de erro
	// Go não tem exceção
	// Não existe try-catch
	if err != nil {
		fmt.Println(err)
		return
	}

	// Retorna informações do arquivo
	stat, err := file.Stat()

	// Sempre checar os erros dos métodos utilizados
	// Isso evita comportamento inesperado no programa
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Tamanho do arquivo: ", stat.Size())
}
