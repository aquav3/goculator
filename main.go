package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

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
