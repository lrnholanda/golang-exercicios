// Em GO a representação de uma string difere da maioria das linguagens como JAVA. Uma string é na verdade um slice de bytes somente leitura e não uma array de caracteres. Se utilizarmos o Println ele irá exibir seu valor decimal e se verificarmos o tipo %T vamos ver que se trata de um uint8

package main

import "fmt"

func main() {
	// w := "Hello World"
	// fmt.Println(w[0])
	// fmt.Printf("%T", w[0])
	// for _, c := range w {
	// 	fmt.Printf("%v %x %T\n", c, c, c)
	var keepGoing = true
	answer := ""
	for keepGoing {
		fmt.Println("Type command: ")
		fmt.Scan(&answer)
		if answer == "quit" {
			keepGoing = false
		}
	}
	fmt.Println("program exit")

}
