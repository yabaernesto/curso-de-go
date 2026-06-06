# Curso de Go

Go é uma linguagem criada para resolver problemas de concorrência.

### Comandos

- Inicializar um módulo go: go mod init nomeDoModulo
- Rodar main: go run main.go
- Bundle do pacote: go bundle -o nome do bundle

---

Inicializar o módulo Go com uma tag github e ome do projeto é um padrão de comunidade (também serve para identificar o módulo).
É recomendando criar um Go módulo (go.mod) com github nome do usuario e depois nome do projeto.

## Os pacotes (packages) servem como namespace para utilizar funções em outros lugares (funções, métodos, variáveis)

Em Go, tudo que começa com a letra maiúscula é exportável (e é possível acessar fora do seu pacote).
Se a letra for minúscula, ela não é exportável é apenas visível ao seu módulo.

## Variáveis e constantes

Em Go existem duas maneiras de criar variáveis, e essas duas maneiras têm variações.
A primeira forma é colocar a palavra reservada "var" o nome da variável e depois o tipo e a sua expressão:

```go
var info string
var texto string = "ola"
var idade int = 0
```

Quando o tipo não é passado, ela sumi o tipo do valor da variável:

```go
var texto = "Yaba"
var idade = 25
```

Go tem um mecanismo de inicialização de variável quando ás variáveis não são inicializadas logo no inicio.

```go
var info string // info vai receber string vazia ""
var idade int // idade vai receber 0
var casado bool // casado vai receber false
```

Declaração de variável curta/short:

```go
info := "alguma coisa informativa"
idade := 25
casado := false
```

### Constantes

Valores não modificados após a declaração

```go
package main

import "fmt"

const nome string = "Yaba"

func main() {
  fmt.Println(nome)
}
```

---
