package tokens

import (
    "sync"
    "strings"
    "strconv"
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
    Value string
    Variant Variant
}

func Tokenize(math string) []Token {
    var wg sync.WaitGroup
    defer wg.Wait()
    elements := strings.Split(math, " ")
    result := make([]Token, len(elements))
    for i, n := range elements {
        wg.Add(1)
        go func(i int, n string) {
            defer wg.Done()
            switch n {
                case "+":
                    result[i] = Token{Value: n, Variant: Plus}
                    return
                case "-":
                    result[i] = Token{Value: n, Variant: Minus}
                    return
                case "*":
                    result[i] = Token{Value: n, Variant: Multiply}
                    return
                case "/":
                    result[i] = Token{Value: n, Variant: Divide}
                    return
            }     
            result[i] = Token{Value: n, Variant: Number}
        }(i, n)
    }

    return result
}

func Operation(operator Token, lhs Token, rhs Token) (int, error) { 
    lhsValue, err := strconv.Atoi(lhs.Value)
    if err != nil {
        return -1, err
    }
    rhsValue, err := strconv.Atoi(rhs.Value)
    if err != nil {
        return -1, err
    }
    switch operator.Variant {
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
