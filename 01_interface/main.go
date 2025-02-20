package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Pessoa interface {
	Desativar()
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c *Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado com sucesso", c.Nome)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	cliente := Cliente{
		Nome:  "João",
		Idade: 30,
		Ativo: true,
		Endereco: Endereco{
			Logradouro: "Rua dos Bobos",
			Numero:     0,
			Cidade:     "São Paulo",
			Estado:     "SP",
		},
	}

	fmt.Println(cliente)
	cliente.Desativar()
	fmt.Println(cliente.Ativo)
}
