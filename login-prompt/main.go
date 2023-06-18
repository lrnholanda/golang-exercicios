package main

import (
	"flag"
	"fmt"
)

func main() {
	var username string
	var password string

	flag.StringVar(&username, "u", "root", "Default username: root")
	flag.StringVar(&password, "p", "mysql", "Default password: mysql")

	if username == "root" && password == "mysql" {
		fmt.Println("MySQL>")
	} else {
		fmt.Println("An Error Occourred.")
	}
}
