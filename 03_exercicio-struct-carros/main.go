package main

import "fmt"

type veiculo struct {
	cor    string
	portas int
}

type caminhonete struct {
	veiculo
	traçãoNasQuatro bool
}

type sedan struct {
	veiculo
	modeloLuxo bool
}

func main() {
	carro01 := caminhonete{veiculo{cor: "azul", portas: 4}, true}
	carro02 := sedan{veiculo{cor: "prata", portas: 4}, true}

	fmt.Println(carro01, carro02)
	fmt.Printf("Caminhonete cor: %s\t Sedan cor: %s\t Sedan modelo de luxo: %t\n", carro01.cor, carro02.cor, carro02.modeloLuxo)

}
