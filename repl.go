package main

import "strings"

/*

The purpose of this function will be to split the user's input into "words" based on whitespace.
It should also lowercase the input and trim any leading or trailing whitespace. For example:

E.G

hello world -> ["hello", "world"]
Charmander Bulbasaur PIKACHU -> ["charmander", "bulbasaur", "pikachu"]


*/

func cleanInput(text string) []string {

	lowered := strings.ToLower(text)
	splitStrings := strings.Fields(lowered)

	return splitStrings

}

//Two main things, iterate over the string, and each time we get whitepsace and the prev char was a charatcer
//Means this is a new word, we split at this point, and make sure the whole string is lowercase

/*



 */
