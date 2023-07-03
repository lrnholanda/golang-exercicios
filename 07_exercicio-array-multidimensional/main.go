package main

import "fmt"

func main() {
	user := [][]string{
		{
			"Lorena",
			"Holanda",
			"ser piranha",
		},
		{
			"Chris",
			"Oliveira",
			"ser padrão",
		},
		{"Italo",
			"Lima",
			"comer cu",
		},
	}
	for _, row := range user {
		for _, val := range row {
			fmt.Println(val)
		}
	}
}
