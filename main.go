package main

import (
    "fmt"
    "log"
    "strings"
)

func main() {
    greet("Alice")
    fmt.Println("Hello, World!")

    for {
        var input string
        fmt.Print("Enter first number (or type 'exit' to quit): ")
        fmt.Scan(&input)

        if strings.ToLower(input) == "exit" {
            fmt.Println("Exiting program.")
            break
        }

        a, err := parseInput(input)
        if err != nil {
            log.Println("Invalid input. Please enter a valid number.")
            continue
        }

        fmt.Print("Enter second number: ")
        fmt.Scan(&input)
        b, err := parseInput(input)
        if err != nil {
            log.Println("Invalid input. Please enter a valid number.")
            continue
        }

        fmt.Println("Choose an operation: +, -, *, /")
        fmt.Scan(&input)
        operation := input

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

func parseInput(input string) (int, error) {
    var num int
    _, err := fmt.Sscan(input, &num)
    return num, err
}
