package main

import "fmt"

type Person struct {
	nome   string
	idade  int
	sexo   string
	cidade string
}

var person Person

func main() {
	fmt.Println("Seja Bem-vindos ao cadastro nacional")
	person.nome, person.idade, person.sexo, person.cidade = cadastrar()
	fmt.Println(person)
}

func cadastrar() (string, int, string, string) {
	var nome string
	var idade int
	var sexo string
	var cidade string

	fmt.Println("Digite seu nome: ")
	fmt.Scan(&nome)

	fmt.Println("Digite sua idade: ")
	fmt.Scan(&idade)

	fmt.Println("Informe seu sexo: ")
	fmt.Scan(&sexo)

	fmt.Println("Digite o nome de sua cidade: ")
	fmt.Scan(&cidade)

	return nome, idade, sexo, cidade
}
