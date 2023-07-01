package main

import "fmt"

func main() {
	array := [5]int{4, 7, 8, 5, 6}

	for _, value := range array {
		fmt.Printf("O item %d, e seu tipo é %T\n", value, value)
	}
}
