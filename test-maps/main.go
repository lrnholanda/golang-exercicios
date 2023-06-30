package main

import (
	"fmt"
)

var contacts = make(map[string]string)

type UserContacts struct {
	contact string
	number  string
}

func main() {
	fmt.Println("Welcome to your phonebook.")
	var command string
	for {
		fmt.Print("Command >_ ")
		fmt.Scan(&command)
		if command == "store" {
			contact, number := getUserInput()
			addNewContact(contact, number)
		} else if command == "list" {
			listContacts()
		} else if command == "lookup" {
			fmt.Print("Enter name: ")
			var contact string
			fmt.Scan(&contact)
			value := checkForContact(contact, contacts)
			if value {
				fmt.Println(contacts[contact])
			} else {
				fmt.Println("Contact doesn't exist, do you want to add it? y/n:")
				var tecla string
				fmt.Scan(&tecla)
				if tecla == "y" {
					contact, number := getUserInput()
					addNewContact(contact, number)
				} else if tecla == "n" {
					return
				}
			}
		} else if command == "quit" {
			break
		} else {
			fmt.Println("Comando Não reconhecido: ", command)
		}
	}
}

func getUserInput() (string, string) {
	var contact string
	var number string

	fmt.Println("Digite o nome do contato: ")
	fmt.Scan(&contact)

	fmt.Println("Digite o número: ")
	fmt.Scan(&number)

	return contact, number
}

func addNewContact(contact string, number string) {
	contacts[contact] = number
	fmt.Println("Contact saved")
}

func listContacts() {
	for key, value := range contacts {
		fmt.Println(key, value)
	}
}

func checkForContact(contact string, contacts map[string]string) bool {
	for _, value := range contacts {
		if value == contact {
			return true
		}
	}
	return false
}
