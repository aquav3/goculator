package tokens

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

type Var string

const (
    Number Var     = "Number"
    Plus Var       = "Plus"
    Minus Var      = "Minus"
    Multiply Var   = "Multiply"
    Divide Var     = "Divide"
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
    result := 0
    firstOperations := make(map[int]Token)
    secondOperations := make(map[int]Token)
    iterations := t
    for i, n := range t {
        if n.Var == Multiply || n.Var == Divide {
            firstOperations[i] = n 
        }
        if n.Var == Plus || n.Var == Minus {
            secondOperations[i] = n       
        }
    }
    
    for key, value := range firstOperations {
        fmt.Println("Left hand side value: ", t[key-1].Value, "Variant: ", t[key-1].Var)
        fmt.Println("Operator value: ", t[key].Value, "Variant: ", t[key].Var)
        fmt.Println("Right hand side value: ", t[key+1].Value, "Variant: ", t[key+1].Var)
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
    
    for key, value := range secondOperations {
        fmt.Println("Left hand side value: ", t[key-1].Value, "Variant: ", t[key-1].Var)
        fmt.Println("Operator value: ", t[key].Value, "Variant: ", t[key].Var)
        fmt.Println("Right hand side value: ", t[key+1].Value, "Variant: ", t[key+1].Var)
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

    return result
}
