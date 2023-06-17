package main

import (
	"fmt"
)

func main() {
	s := "Hello, playground"
	sb := []byte(s)

	for _, v := range sb {
		fmt.Printf("%v - %T -%#U - %#x\n", v, v, v, v)
	}

}
