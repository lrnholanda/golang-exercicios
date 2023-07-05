package main

import "fmt"

func main() {
	total, quantos := soma(10, 10, 1, 2, 3, 5)
	fmt.Println(total, quantos)
}

func soma(x ...int) (int, int) {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma, len(x)
}
