package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
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

func cleanInput(text string) []string {

	lowered := strings.ToLower(text)
	splitStrings := strings.Fields(lowered)

	return splitStrings

}

func startRepl(cfg *config) {

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
			err := cmd.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}

		} else {
			fmt.Println("Unknown command")
		}

	}
}
