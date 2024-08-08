package main

import (
    "fmt"
)

func main() {
    greet("Alice")
    fmt.Println("Hello, World!")
    
    a, b := 5, 3
    sumResult := sum(a, b)
    subtractResult := subtract(10, 4)
    multiplyResult := multiply(7, 6)
    divideResult := divide(20, 4)

    fmt.Printf("Results: Sum=%d, Subtract=%d, Multiply=%d, Divide=%s\n", sumResult, subtractResult, multiplyResult, divideResult)
}

func greet(name string) {
    fmt.Printf("Hello, %s!\n", name)
}

func sum(a, b int) int {
    return a + b
}

func subtract(a, b int) int {
    return a - b
}

func multiply(a, b int) int {
    return a * b
}

func divide(a, b int) string {
    if b == 0 {
        return "Error: Division by zero"
    }
    return fmt.Sprintf("%d", a/b)
}
