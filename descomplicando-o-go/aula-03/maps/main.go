package main

import "fmt"

func main() {
	// Método 1 de criação de maps - Usado quando vai definir chave-valor dentro de um for
	alfabeto := make(map[string]string)
	alfabeto["vogais"] = "aeiou"
	alfabeto["consoantes"] = "bcdfghjklmnpqrstvwxyz"

	fmt.Println(alfabeto)
	fmt.Println(alfabeto["vogais"])
	fmt.Println(alfabeto["consoantes"])

	// Método 2 de criação de maps - Usado quando já sabe todos os valores inicialmente
	alfabeto = map[string]string{
		"vogais":     "aeiou",
		"consoantes": "bcdfghjklmnpqrstvwxyz",
	}

	fmt.Println(alfabeto)
	fmt.Println(alfabeto["vogais"])
	fmt.Println(alfabeto["consoantes"])
}
