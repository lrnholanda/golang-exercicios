package main

import (
	"errors"
	"fmt"
)

var NoTooSmall = errors.New("the number is to small")

func main() {
	fmt.Println("Olá")
	no, err := returnPositive(-2)
	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println("value: ", no)
	}

	// What you are seeing above is how we use an if clause to check for errors, if so, print out. On the else, we have our actual value.
}

func divide(nominator int, divider int) float32 {
	defer errorHandler()
	if divider == 0 {
		panic("can't divede by 0")
	}
	return float32(nominator) / float32(divider)
}

func errorHandler() {
	if r := recover(); r != nil {
		fmt.Println("Recovered ", r)
	}
}

// We can improve this function, by ensuring it always returns the result and an error, like so

func returnPositive(no int) (int, error) {
	if no > 0 {
		return no, nil
	} else {
		return 0, NoTooSmall
	}
}
