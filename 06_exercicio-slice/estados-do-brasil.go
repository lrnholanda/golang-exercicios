package main

import "fmt"

func main() {
	// brasil := make([]string, 26, 30)
	// estados := []string{"Acre",
	// 	"Alagoas",
	// 	"Amapá",
	// 	"Amazonas",
	// 	"Bahia",
	// 	"Ceará",
	// 	"Distrito Federal",
	// 	"Espírito Santo",
	// 	"Goiás",
	// 	"Maranhão",
	// 	"Mato Grosso",
	// 	"Mato Grosso do Sul",
	// 	"Minas Gerais",
	// 	"Pará",
	// 	"Paraíba",
	// 	"Paraná",
	// 	"Pernambuco",
	// 	"Piauí",
	// 	"Rio de Janeiro",
	// 	"Rio Grande do Norte",
	// 	"Rio Grande do Sul",
	// 	"Rondônia",
	// 	"Roraima",
	// 	"Santa Catarina",
	// 	"São Paulo",
	// 	"Sergipe",
	// 	"Tocantins"}
	// //estados = append(estados)

	// for _, estado := range estados {
	// 	brasil = append(brasil, estado)
	// }

	// fmt.Println(brasil)
	// fmt.Println(len(brasil))
	// fmt.Println(cap(brasil))

	brasil := make([]string, 26, 26)
	brasil = []string{
		"Acre",
		"Alagoas",
		"Amapá",
		"Amazonas",
		"Bahia",
		"Ceará",
		"Distrito Federal",
		"Espírito Santo",
		"Goiás",
		"Maranhão",
		"Mato Grosso",
		"Mato Grosso do Sul",
		"Minas Gerais",
		"Pará",
		"Paraíba",
		"Paraná",
		"Pernambuco",
		"Piauí",
		"Rio de Janeiro",
		"Rio Grande do Norte",
		"Rio Grande do Sul",
		"Rondônia",
		"Roraima",
		"Santa Catarina",
		"São Paulo",
		"Sergipe",
		"Tocantins",
	}

	fmt.Println(len(brasil), cap(brasil))

	for i := 0; i < len(brasil); i++ {
		fmt.Println(brasil[i])
	}

}
