package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	firstName, lastName, s string
)
var inputReader *bufio.Reader
var input string
var err error

func main() {
	worldLetterCount()
	json1()
}

func func1() {
	fmt.Println("Please enter your full name: ")
	fmt.Scanln(&firstName, &lastName)
	// fmt.Scanf("%s %s", &firstName, &lastName)
	fmt.Printf("Hi %s %s!\n", firstName, lastName)
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input: ")
	input, err = inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("The input was: %s\n", input)
	}
}

func worldLetterCount() {
	nrchars, nrwords, nrlines := 0, 0, 0
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input, type S to stop: ")
	for {
		input, err = inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error occured: %s\n", err)
		}
		if input == "S\n" {
			fmt.Println("Here are the statistics:")
			fmt.Printf("Number for chars: %d\n", nrchars)
			fmt.Printf("Number for words: %d\n", nrwords)
			fmt.Printf("Number for lines: %d\n", nrlines)
			break
		} else {
			nrchars += len(input) - 1
			nrwords += len(strings.Fields(input))
			nrlines++
		}
	}
}

func json1() {
	b := []byte(`{"Name": "Wednesday", "Age": 6, "Parents": ["Gomez", "Morticia"]}`)
	fmt.Printf("%s", b)
}
