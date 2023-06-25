package main

import "fmt"

var x interface{}

func main() {
	// x = true

	// switch x.(type) {
	// case int:
	// 	fmt.Println("é um int")
	// case bool:
	// 	fmt.Println("é um bool")
	// case string:
	// 	fmt.Println("é um string")
	// case float64:
	// 	fmt.Println("é um float")
	// }
	x := 9

	if !(x%2 == 0) && x%3 == 0 {
		fmt.Println("é múltiplo de três e dois")
	}

	if x%2 == 0 || x%3 == 0 {
		fmt.Println("é múltiplo de três ou dois")
	}
}
