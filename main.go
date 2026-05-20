package main

import (
	"bufio"
	"fmt"
	"os"
)

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

		command := cleanedInput[0]

		if len(command) == 0 {
			continue
		}

		fmt.Printf("Your command was: %s\n", command)

	}

}
