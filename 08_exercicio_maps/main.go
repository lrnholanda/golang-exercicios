package main

import "fmt"

func main() {
	user := map[string][]string{
		"holanda_lorena": {
			"ser piranha",
			"estudar golang",
			"ser trouxa",
		},
		"oliveira_chris": {
			"ser padrão",
			"ouvir taylor s",
			"se interessar por low profiles",
		},
	}

	for k, value := range user {
		fmt.Println(k)
		for i, h := range value {
			fmt.Println("\t", i, " - ", h)
		}
	}
}
