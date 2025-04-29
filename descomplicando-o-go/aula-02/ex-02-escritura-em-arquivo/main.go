package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {
	txt := []byte("Esse eh o meu novo arquivo")

	// 0666 é a permissão do arquivo para leitura e escrita
	// O arquivo será criado se não existir
	// Se o arquivo já existir, ele será sobrescrito
	perm := fs.FileMode(0666)
	os.WriteFile("arquivo-novo.txt", txt, perm)

	// O método Create cria um novo arquivo
	// Método mais tradicional
	file, err := os.Create(("arquivo-novo.txt"))

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	// Métodos de sobrescrita de arquivo
	file.Write(txt)
	file.WriteString("Minha string")
	file.WriteAt(txt, 40)

	// Método de escrita de arquivo, adicionando o conteúdo ao final do arquivo
	// Permissão de escrita apenas
	// O arquivo será criado se não existir
	file, err = os.OpenFile("arquivo-novo.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)

	// Já que adicionamos as restrições acima (de adicionar ao fim do arquivo e somente escrita),
	// podemos utilizar os mesmos métodos de escrita
	file.Write(txt)
	file.WriteString("Minha string")
	file.WriteAt(txt, 40)
}
