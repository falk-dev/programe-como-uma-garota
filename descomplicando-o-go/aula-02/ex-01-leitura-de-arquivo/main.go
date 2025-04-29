package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// https://dados.gov.br/dados/conjuntos-dados/preco-de-medicamentos-no-brasil-consumidor

	file, err := os.Open("TA_PRECO_MEDICAMENTO.csv")

	if err != nil {
		fmt.Println(err)
		return
	}

	// Defer é usado para controle de fluxo
	// Neste caso, o comando abaixo é executado quando a função main termina
	// Importante utilizar o defer após checar se o arquivo foi aberto corretamente
	// Assim, o arquivo só será fechado se não houver erro ao abri-lo
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println(err)
		return
	}

}
