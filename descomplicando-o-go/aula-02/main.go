package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// Retorna uma tupla com dois valores
	// A função Open retorna um ponteiro para o arquivo e um erro
	arquivo, err := os.Open("arquivo.txt")

	// Forma padrão de tratamento de erro em Go
	// Go não tem exceptions, então não existe try-catch
	if err != nil {
		fmt.Println(err)
		return
	}

	// Retorna informações do arquivo
	stat, err := arquivo.Stat()

	// Sempre checar os erros dos métodos utilizados
	// Isso evita comportamento inesperado no programa
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Tamanho do arquivo: ", stat.Size())

	// 1. Leitura do arquivo completo
	// Retorna um slice de bytes
	// Indicado para arquivos pequenos, pois lê tudo de uma vez só
	// Se o arquivo for muito grande, pode causar problemas de memória
	file, err := os.ReadFile("arquivo.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	// Converte o slice de bytes para string
	// e imprime o conteúdo do arquivo
	fmt.Println("Conteúdo do arquivo: ", string(file))

	// 2. Leitura do arquivo linha por linha
	// Indicado para arquivos grandes, pois lê linha por linha
	arquivo, err = os.Open("arquivo.txt")

	if err != nil {
		fmt.Println(err)
	}

	// Atribui o arquivo a um scanner
	scanner := bufio.NewScanner(arquivo)

	// O método Scan retorna true enquanto houver linhas para ler
	for scanner.Scan() {
		// Imprime o conteúdo da linha
		// O método Text retorna o conteúdo da linha lida
		fmt.Println(scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 3. Leitura do arquivo byte a byte
	arquivo, err = os.Open("arquivo.txt")

	// Criando um leitor de dados do arquivo e convertendo para um buffer
	reader := bufio.NewReader(arquivo)

	for {
		// O método ReadByte lê um byte do arquivo
		// Retorna o byte lido e um erro
		b, err := reader.ReadByte()

		if err != nil {
			if err != io.EOF {
				// Se o erro não for EOF, imprime o erro
				fmt.Println(err.Error())
			}
			break
		}

		// Imprime o byte lido
		// O byte lido é um número inteiro, então convertemos para string
		fmt.Println(string(b))
	}
}
