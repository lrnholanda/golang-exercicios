package main

import "fmt"

func main() {
	slice := []int{20, 45, 67, 98, 23, 2, 4, 5, 84, 24}
	// 	for _, value := range slice {
	// 		fmt.Printf("O item %d, e seu tipo é %T\n", value, value)
	// 	}

	fmt.Println(slice[:3])
	fmt.Println(slice[4:])
	fmt.Println(slice[1:7])
	fmt.Println(slice[2:9])
	fmt.Println(slice[2 : len(slice)-1])
}
