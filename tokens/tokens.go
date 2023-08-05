package tokens

import (
	"strconv"
	"strings"
	"sync"
)

type Var uint8

const (
    Number Var     = 0
    Plus Var       = 1
    Minus Var      = 2
    Multiply Var   = 3
    Divide Var     = 4
)

type Token struct {
    Value string
    Var Var
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
                    result[i] = Token{Value: n, Var: Plus}
                    return
                case "-":
                    result[i] = Token{Value: n, Var: Minus}
                    return
                case "*":
                    result[i] = Token{Value: n, Var: Multiply}
                    return
                case "/":
                    result[i] = Token{Value: n, Var: Divide}
                    return
            }     
            result[i] = Token{Value: n, Var: Number}
        }(i, n)
    }

    return result
}

func operation(operator Token, lhs Token, rhs Token) (int, error) { 
    lhsValue, err := strconv.Atoi(lhs.Value)
    if err != nil {
        return -1, err
    }
    rhsValue, err := strconv.Atoi(rhs.Value)
    if err != nil {
        return -1, err
    }
    switch operator.Var {
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

func Compute(t []Token) int {
    abc := 0
    firstOperations := make(map[int]Token)
    iterations := t
    for i, n := range t {
        if n.Var == Multiply || n.Var == Divide {
            firstOperations[i] = n 
        } 
    }
    
    for key, value := range firstOperations {
        key = key - (2 * abc)
        result, _ := operation(value, t[key-1], t[key+1])
        tk := Token{Value: strconv.Itoa(result), Var: Number}
        iterations[key-1] = tk
        lhs := iterations[:key]
        if key+2 > len(t) {
            break
        }
        rhs := iterations[key+2:]

        iterations = append(lhs, rhs...)
        abc += 1
    }
    
    if len(iterations) == 1 {
        result, _ := strconv.Atoi(iterations[0].Value)
        return result
    }
    return -69
}
