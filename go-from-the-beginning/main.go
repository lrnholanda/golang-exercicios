package main

import "fmt"

func main() {
	//var response string
	//fmt.Scan(&response)
	//fmt.Println("User typed: ", response)

	// var a1, a2 string
	// // multiple input
	// fmt.Scan(&a1, &a2)

	// // formatted string to print out the user input
	// str := fmt.Sprintf("a1: %s a2: %s", a1, a2)

	// fmt.Println(str)

	// With formatters, you’re able to pick the part of the user input you are interested in and place that into the variable you want.

	// var prefix string
	// var no int
	// fmt.Scanf("%3s%d", &prefix, &no)
	// fmt.Printf("prefix: %s, invoice no: %d", prefix, no)
	var numberPlayers int
	var player string
	s, sep := "", ""
	fmt.Println("Enter number of players : ")
	fmt.Scan(&numberPlayers)
	for i := 0; i < numberPlayers; i++ {
		fmt.Printf("Enter Player %d name: ", i+1)
		fmt.Scan(&player)
		s += sep + player
		sep = " "
	}
	fmt.Println("Players are: ", s)
}
