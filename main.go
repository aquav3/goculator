package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
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

func tokenize(math string) []Token {
    elements := strings.Split(math, " ")
    result := make([]Token, len(elements))
    for i, n := range elements {
        go func(i int, n string) {
            switch n {
                case "+":
                    result[i] = Token{value: n, variant: Plus}
                    return
                case "-":
                    result[i] = Token{value: n, variant: Minus}
                    return
                case "*":
                    result[i] = Token{value: n, variant: Multiply}
                    return
                case "/":
                    result[i] = Token{value: n, variant: Divide}
                    return
            }     
            result[i] = Token{value: n, variant: Number}
        }(i, n)
    }

    return result
}

func operation(operator Token, lhs Token, rhs Token) (int, error) { 
    lhsValue, err := strconv.Atoi(lhs.value)
    if err != nil {
        return -1, err
    }
    rhsValue, err := strconv.Atoi(rhs.value)
    if err != nil {
        return -1, err
    }
    switch operator.variant {
        case Plus:
            return lhsValue + rhsValue, nil
        case Minus:
            return lhsValue - rhsValue, nil
        case Multiply:
            return lhsValue * rhsValue, nil
        case Divide:
            return lhsValue / rhsValue, nil
    }
    return 0, nil
}

func main() {
    input, err := getInput("Enter the math dude: ") 
    if err != nil {
        log.Fatalln("Error: ", err)
    }

    fmt.Println(input)
}
