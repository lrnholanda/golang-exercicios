package main

import "fmt"

type abacaxi int

var x abacaxi

func main() {
	fmt.Printf("O valor de x: %v\n", x)
	fmt.Printf("O tipo de x: %T\n", x)
	x = 42
	fmt.Printf("O valor de x: %v\n", x)
}
