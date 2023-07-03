package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

type profissao struct {
	pessoa
	titulo  string
	salario float64
}

func main() {
	pessoa1 := pessoa{
		nome:  "Alfredo",
		idade: 30,
	}

	pessoa2 := profissao{
		pessoa: pessoa{
			nome:  "Maricota",
			idade: 31,
		},
		titulo:  "pizzaiola",
		salario: 10000,
	}

	fmt.Println(pessoa1.nome)
	fmt.Println(pessoa2.salario)

}
