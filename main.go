package main

import (
    "fmt"
    "log"
)

func main() {
    greet("Alice")
    fmt.Println("Hello, World!")

    for {
        var a, b int
        fmt.Print("Enter first number (or type 'exit' to quit): ")
        if _, err := fmt.Scan(&a); err != nil {
            log.Println("Invalid input. Exiting.")
            break
        }

        fmt.Print("Enter second number: ")
        if _, err := fmt.Scan(&b); err != nil {
            log.Fatalf("Invalid input: %v", err)
        }

        fmt.Println("Choose an operation: +, -, *, /")
        var operation string
        fmt.Scan(&operation)

        switch operation {
        case "+":
            fmt.Printf("Result: %d\n", sum(a, b))
        case "-":
            fmt.Printf("Result: %d\n", subtract(a, b))
        case "*":
            fmt.Printf("Result: %d\n", multiply(a, b))
        case "/":
            fmt.Printf("Result: %s\n", divide(a, b))
        default:
            fmt.Println("Invalid operation")
        }
        fmt.Println("---------")
    }
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
