package main

import "fmt"

type abacaxi int

var x abacaxi
var y int

func main() {
	fmt.Printf("O valor de x: %v\n", x)
	fmt.Printf("O tipo de x: %T\n", x)
	x = 42
	fmt.Printf("O valor de x: %v\n", x)

	y = int(x)
	fmt.Printf("O valor de y: %v\n", y)
	fmt.Printf("O tipo de y: %T\n", y)
}
