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
    if err != nil {
        log.Fatalln("Error:", err)
    }
    fmt.Println(result)
}
