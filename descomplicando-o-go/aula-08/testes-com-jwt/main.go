package main

import (
	"fmt"

	"github.com/falk-dev/programe-como-uma-garota/tree/main/descomplicando-o-go/aula-04/jwt/tokens"
)

func main() {
	fmt.Println("Gerando JWT...")
	token := tokens.Generate()
	fmt.Println(token)
	fmt.Println("JWT gerado com sucesso!")
}
