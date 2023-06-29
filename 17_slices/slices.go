package main

import (
	"fmt"
)

func main() {
	sabores := []string{"pepperoni", "mozzarela", "abacaxi", "quatro queijos", "portuguesa"}
	fatia := sabores[:]
	fmt.Println(fatia)
	sabores = append(sabores[:2], sabores[3:]...)
	fmt.Println(sabores)

	umaslice := []int{1, 2, 3, 4}
	outraslice := []int{9, 10, 11, 12}

	// fmt.Println(umaslice)
	// umaslice = append(umaslice, 5, 6, 7, 8)
	// fmt.Println(umaslice)
	umaslice = append(umaslice, outraslice...)
	fmt.Println(umaslice)
}
