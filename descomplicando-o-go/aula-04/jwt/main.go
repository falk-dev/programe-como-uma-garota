package main

import (
	"fmt"

	"github.com/falk-dev/programe-como-uma-garota/tree/main/descomplicando-o-go/aula-04/jwt/tokens"
)

func main() {
	fmt.Println("Gerando JWT...")
	tokens.Generate()
	fmt.Println("JWT gerado com sucesso!")
}
