package main

import (
	"fmt"

	"github.com/yabaernesto/curso-go/functions"
	"github.com/yabaernesto/curso-go/methods"
	"github.com/yabaernesto/curso-go/structs"
)

func main() {
	structs.Estrutura()
	response := functions.Soma(10, 15)

	fmt.Println("Soma: ", response)

	multiplica := func (x int) int {
    return x * 2
  }

  resultado := multiplica(2)
  fmt.Println("Multiplicação", resultado)

	// metodo
	p1 := methods.Pessoa{Nome: "Yaba", Idade: 25}
  p2 := methods.Pessoa{Nome: "Ernesto", Idade: 25}
	p1.Apresentar()
  p2.Apresentar()
}
