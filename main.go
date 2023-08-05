package main

import (
	"fmt"
	"log"
    "github.com/aquav3/goculator/utils"
	"github.com/aquav3/goculator/tokens"
)
 func main() {
    input, err := utils.GetInput("Enter the math dude: ") 
    if err != nil {
        log.Fatalln("Error: ", err)
    }
    tokeys := tokens.Tokenize(input)
    result, err := tokens.Operation(tokeys[1], tokeys[0], tokeys[2])
    if err != nil {
        log.Fatalln("Error:", err)
    }
    fmt.Println(result)
}
