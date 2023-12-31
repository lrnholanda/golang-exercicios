package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"runtime/debug"
)

var NoTooSmall = errors.New("the number is to small")

func main() {
	f, err := os.OpenFile("logs", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Println(err)
	}
	log.SetOutput(f)
	// What you are seeing above is how we use an if clause to check for errors, if so, print out. On the else, we have our actual value.

	log.Println("starting program")
	no := divide(10, 2)
	fmt.Println(no)

	no = divide(10, 1)
	fmt.Println(no)
	f.Close()
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
		fmt.Println(r, string(debug.Stack()))
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
