package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func add(first int, second int) int {
	return first + second
}

// func main() {
// 	arg_len := len(os.Args[1:])
// 	fmt.Println("Argument length: ", arg_len)
// 	no1, _ := strconv.Atoi(os.Args[1])
// 	no2, _ := strconv.Atoi(os.Args[2])
// 	var sum = add(no1, no2)
// 	fmt.Printf("Sum equals to: %v", sum)
// }

// func main() {
// 	no, err := strconv.Atoi(os.Args[1])
// 	fmt.Println(no)
// 	if err != nil {
// 		fmt.Println(err)
// 		fmt.Println("Couldn't convert: " + os.Args[1])
// 	} else {
// 		fmt.Println(no)
// 	}
// }

func main() {
	var no int = 100
	fmt.Println(reflect.TypeOf(no))

	var intStr string = "100"
	q := strconv.QuoteToASCII("Hello, 世界")
	fmt.Println(q)
	fourBaseEightBitInt, _ := strconv.ParseInt(intStr, 4, 8)
	tenBaseSixteenBitInt, _ := strconv.ParseInt(intStr, 10, 16)
	fmt.Println(reflect.TypeOf(fourBaseEightBitInt))
	fmt.Println(reflect.TypeOf(tenBaseSixteenBitInt))
}
