package main

import (
	"fmt"
)

type Pessoa struct {
	nome      string
	sobrenome string
	sabores   []string
}

func main() {
	// x := struct {
	// 	nome  string
	// 	idade int
	// }{
	// 	nome:  "Mario",
	// 	idade: 50,
	// }

	pessoa01 :=
		Pessoa{
			nome:      "Lorena",
			sobrenome: "Holanda",
			sabores:   []string{"romeu e julieta", "chocolate", "creme", "napolitano"}}

	pessoa02 :=
		Pessoa{
			nome:      "Tetê",
			sobrenome: "Holanda",
			sabores:   []string{"romeu e julieta", "flocos", "creme", "napolitano"}}

	fmt.Println(pessoa01.nome, pessoa01.sobrenome)
	for _, value := range pessoa01.sabores {
		fmt.Println("\t", value)
	}

	fmt.Println(pessoa02.nome, pessoa02.sobrenome)
	for _, value := range pessoa02.sabores {
		fmt.Println("\t", value)
	}
}
