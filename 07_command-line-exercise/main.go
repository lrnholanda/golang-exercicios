package main

import (
	"os"
	"strconv"
)

func sun(numb1 int, numb2 int) int {
	return numb1 + numb2
}

func main() {
	no1, err := strconv.Atoi(os.Args[1])
	no2, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(err)
	} else {
		var result = sun(no1, no2)
		println(result)
	}

}
