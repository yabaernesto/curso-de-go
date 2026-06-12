package methods

import "fmt"

type Pessoa struct {
	Nome string
	Idade int
}

// metodo. Pessoa recebe o metodo Apresentar, o "p" pode ser qualquer coisa
func (p Pessoa) Apresentar() {
	fmt.Printf("Olá, meu nome é %s e tenho %d anos.\n", p.Nome, p.Idade)
}
