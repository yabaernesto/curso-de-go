package structs

import (
	"fmt"
)

type Client struct {
	Nome    string
	Idade   int
	// struct dentro de struct
	Endreco Endereco
	Email   string
}

type Endereco struct {
	Rua    string
	Numero int
	Cidade string
}

func Estrutura() {
	client1 := Client{
		Nome:    "Yaba",
		Idade:   25,
		Endreco: Endereco {
			Rua:    "Rua A",
			Numero: 123,
			Cidade: "Rangel",
		},
		Email:   "yaba@email.com",
	}

	client2 := Client{
		Nome:    "Maria",
		Idade:   25,
		Endreco: Endereco {
			Rua:    "Rua B",
			Numero: 456,
			Cidade: "Rio de Janeiro",
		},
		Email:   "maria@gmail.com",
	}

	client2.Email = "maria@outlook.com"

	fmt.Println(client1)
	fmt.Println(client2)
	fmt.Println(client2.Endreco.Cidade)
}
