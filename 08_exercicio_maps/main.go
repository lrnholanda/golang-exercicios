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
		"bhargara_aditya": {
			"falar sobre algoritmos e big O",
			"escrever livros",
			"pintar quadros bonitos",
		},
	}
	user["loureiro_kiko"] = []string{"lavar os cabelos com loreal", "tocar fogo na guitarra"}
	//delete(user, "oliveira_chris")
	for k, value := range user {
		fmt.Println(k)
		for i, h := range value {
			fmt.Println("\t", i, " - ", h)
		}
	}
}
