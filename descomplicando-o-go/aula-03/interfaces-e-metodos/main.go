package main

import (
	"encoding/json"
	"fmt"
)

// Interfaces são definições de métodos que uma struct pode ter
type Veiculo interface {
	Buzinar()
}

type Carro struct {
	Marca  string
	Modelo string
}

type Moto struct {
	Modelo string
	Ano    int
}

func (c *Carro) Buzinar() {
	fmt.Println("beep")
}

func (c *Moto) Buzinar() {
	fmt.Println("boop")
}

func buzina(v Veiculo) {
	v.Buzinar()
}

func main() {

	carro := Carro{
		Marca:  "Fiat",
		Modelo: "Argo",
	}

	moto := Moto{
		Modelo: "CG 160",
		Ano:    2016,
	}

	// Para passar as variáveis 'moto' e 'carro' para a função 'buzina', é preciso enviar o endereço dessas variáveis
	// Porque a 'interface' nada mais é do que a referência a um endereço de memória
	buzina(&carro)
	buzina(&moto)

	veiculos := make([]Veiculo, 2)
	veiculos[0] = &carro
	veiculos[1] = &moto

	for i := 0; i < len(veiculos); i++ {
		buzina(veiculos[i])
	}

	// JSON
	// 'json.Marshal' é usado para converter estruturas de dados em Go para o formato JSON
	// Recebe como entrada uma interface e retorna um slice de bytes, []byte.
	data, err := json.Marshal(veiculos)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Transformado para string, para ficar codificado em formato JSON
	fmt.Println(string(data))

	err = json.Unmarshal([]byte(`{"Modelo": "CG 160"}`), &moto)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf(string(moto.Modelo))
	fmt.Println(moto.Ano)
}
