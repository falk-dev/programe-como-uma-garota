package main

import (
	"github.com/falk-dev/programe-como-uma-garota/tree/main/descomplicando-o-go/aula-04/variaveis-de-ambiente/envloader"
	"github.com/joho/godotenv" // Biblioteca para carregar variáveis do .env automaticamente
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	err := godotenv.Load(".env")
	checkErr(err)
	envloader.Load()
}
