package main

import (
    "fmt"
)

func main() {
    greet("Alice")
    fmt.Println("Hello, World!")

    var a, b int
    fmt.Print("Enter first number: ")
    fmt.Scan(&a)
    fmt.Print("Enter second number: ")
    fmt.Scan(&b)

    sumResult := sum(a, b)
    subtractResult := subtract(a, b)
    multiplyResult := multiply(a, b)
    divideResult := divide(a, b)

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
