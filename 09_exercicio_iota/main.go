package main

import "fmt"

const (
	a = iota + 2012
	b
	c
	_
)

func main() {
	x := 31337
	fmt.Println(a, b, c)
	fmt.Printf("%#x", x)

}
