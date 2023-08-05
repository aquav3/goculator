package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Variant uint8

const (
    Number Variant     = 0
    Plus Variant       = 1
    Minus Variant      = 2
    Multiply Variant   = 3
    Divide Variant     = 4
)

type Token struct {
    value string
    variant Variant
}

func getInput(prompt string) (string, error) {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print(prompt)
    input, err := reader.ReadString('\n')
    if err != nil {
        return "", err 
    }
    input = strings.TrimRight(input, "\n")

    return input, nil
}

func main() {
    input, err := getInput("Enter the math dude: ") 
    if err != nil {
        log.Fatalln("Error: ", err)
    }

    fmt.Println(input)
}
