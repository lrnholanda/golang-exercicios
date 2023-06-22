package main

import (
	"fmt"
)

//Loops: utilizando ascii

func main() {
	for i := 33; i <= 122; i++ {
		fmt.Printf("Decimal: %d\tHexadecimal: %#x\tUnicode: %c\n", i, i, i)
	}
}
