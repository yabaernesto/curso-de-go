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
