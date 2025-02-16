package main

import "fmt"

type Cliente struct {
	nome string
}

func (c Cliente) exibeNome() {
	c.nome = "Maria"
	fmt.Printf("O nome do cliente é %s\n", c.nome)
}

func main() {
	cliente := Cliente{
		nome: "João",
	}

	cliente.exibeNome()   // O nome do cliente é Maria
	println(cliente.nome) // João
}
