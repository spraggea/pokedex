package main

import (
	"fmt"
)

func commandhelp(*config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n")

	for key, value := range getCommands() {
		fmt.Printf("%s: %s\n", key, value.description)

	}
	return nil
}
