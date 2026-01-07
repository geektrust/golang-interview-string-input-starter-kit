package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		panic(fmt.Errorf("no command line arguments passed"))
	}

	/*
	 * Format of the 'args' slice:
	 * [
	 *   "<INPUT_STRING_1>",
	 *   "<INPUT_STRING_2>",
	 *   ...
	 * ]
	 *
	 * Each element in 'args' represents one complete input string.
	 * We iterate over the inputs and process them one by one.
	 */
	for _, arg := range args {
		handle(arg)
	}
}

/*
 * This function parses the input string and extracts the required value.
 *
 * Format of input:
 * "<INPUT_STRING>"
 * Example:
 * "HelloWorld"
 *
 * The input string is trimmed and split using space as the delimiter.
 * The first token is retrieved and stored for further processing.
 *
 * You may extend this function to extract multiple values depending
 * on the problem statement.
 */
func handle(input string) {
	inputList := strings.Split(strings.TrimSpace(input), " ")
	value := inputList[0]

	// Start your implementation from here.
	fmt.Println("Input String:", value)
}