package main

import "fmt"

func main() {
	x := 2
	y := 3
	soma := add(x, y)
	fmt.Println(soma)
	sum, product := calc(1, 2)
	fmt.Println(sum)
	fmt.Println(product)
	log("hi, Lorena!")
	log("what's up!")
}

func add(first int, second int) (sum int) {
	sum = first + second
	return
}

// Multiple returns

func calc(first int, second int) (sum int, product int) {
	sum = first + second
	product = first * second
	return
}

func log(message string) {
	fmt.Println(message)
}
