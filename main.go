package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type config struct {
	Next     *string
	Previous *string
}

func getCommands() map[string]cliCommand { // Storing commands in a map that we can access via this func
	//Helps get past the cycle initialization of var names the callback func.

	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},

		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandhelp,
		},
		"map": {
			name:        "map",
			description: "Displays 20 new map locations",
			callback:    commandmap,
		},
	}
}

func commandExit(*config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil

}

func commandhelp(*config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n")

	for key, value := range getCommands() {
		fmt.Printf("%s: %s\n", key, value.description)

	}
	return nil
}

func commandmap(cfg *config) error {
	/* user types "map"
	-> command reads config.Next (or default URL if empty)
	-> GET request
	-> unmarshal JSON
	-> print names
	-> update config.Next and config.Previous from the response
	*/

}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex >")

		success := scanner.Scan()
		if !success {
			break

		}

		userInput := scanner.Text()

		cleanedInput := cleanInput(userInput)
		if len(cleanedInput) == 0 {
			continue
		}
		command := cleanedInput[0]

		cmd, exists := getCommands()[command]
		if exists {
			err := cmd.callback()
			if err != nil {
				fmt.Println(err)
			}

		} else {
			fmt.Println("Unknown command")
		}

	}

}
