package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type PokemonGetResponse struct {
	Count    int64
	Next     string
	Previous string
	Results  []struct {
		Name string
		Url  string
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
	return
}

func main() {
	url := "https://pokeapi.co/api/v2/pokemon"
	fmt.Println("Acessando a API na url", url)

	resp, err := http.Get(url)
	checkErr(err)

	fmt.Println("O status code retornado pela url é", resp.StatusCode)

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	checkErr(err)

	fmt.Println("O corpo da requisição retornado é", string(body))
	fmt.Println(resp.Header)

	pokemonGetResponse := new(PokemonGetResponse)

	err = json.Unmarshal(body, pokemonGetResponse)
	checkErr(err)

	fmt.Println("O nome dos 5 primeiros pokemons retornados são:")

	for i := 0; i < 5; i++ {
		fmt.Println(pokemonGetResponse.Results[i].Name)
	}
}
