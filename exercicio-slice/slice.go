package main

import "fmt"

func main() {
	slice := []int{20, 45, 67, 98, 23, 2, 4, 5, 84, 24}
	for _, value := range slice {
		fmt.Printf("O item %d, e seu tipo é %T\n", value, value)
	}
}
