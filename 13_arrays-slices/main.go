package main

import "fmt"

func main() {
	// Arrays
	cities := [5]string{"NY", "LA"}
	ids := [...]int{1, 2, 3, 4}

	// fmt.Println(len(cities))
	// fmt.Println(cap(cities))
	fmt.Println("-------- Arrays --------")

	for _, citie := range cities {
		fmt.Println(citie)
	}

	for _, id := range ids {
		fmt.Println(id)
	}

	fmt.Println("-------- Slices --------")
	// Slices
	var numbers []int
	numbers = append(numbers, 1)
	numbers = append(numbers, 5, 6)
	fmt.Println(numbers)

	letters := []string{"A", "B", "C", "D", "E"}
	remove := 2 // remove index
	// 0 - remove index, remove +1 to end
	letters = append(letters[:remove], letters[remove+1:]...)

	// Create a slice with make()

	slice := make([]int, 2, 5)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	// Copy elements

	arr := [3]int{1, 2, 3}
	dest := make([]int, 5)
	copy(dest, arr[0:2]) // copies slice {1,2} into dest
	fmt.Println(dest)

}
