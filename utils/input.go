package utils

import (
    "bufio"
    "fmt"
    "strings"
    "os"
)

func GetInput(prompt string) (string, error) {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print(prompt)
    input, err := reader.ReadString('\n')
    if err != nil {
        return "", err 
    }
    input = strings.TrimRight(input, "\n")
    return input, nil
}


