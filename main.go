package main

import (
    "fmt"
    "log"
    "math"
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

        var operation string
        for {
            fmt.Println("Choose an operation: +, -, *, /, ^ (power)")
            fmt.Scan(&input)
            operation = input

            fmt.Printf("You chose operation '%s'. Confirm? (yes/no): ", operation)
            fmt.Scan(&input)
            if strings.ToLower(input) == "yes" {
                break
            }
        }

        switch operation {
        case "+":
            fmt.Printf("Result: %d\n", sum(a, b))
        case "-":
            fmt.Printf("Result: %d\n", subtract(a, b))
        case "*":
            fmt.Printf("Result: %d\n", multiply(a, b))
        case "/":
            if b == 0 {
                fmt.Println("Cannot divide by zero. Please enter a new second number.")
                continue
            }
            fmt.Printf("Result: %s\n", divide(a, b))
        case "^":
            fmt.Printf("Result: %.0f\n", power(a, b))
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
    return fmt.Sprintf("%d", a/b)
}

func power(a, b int) float64 {
    return math.Pow(float64(a), float64(b))
}

func parseInput(input string) (int, error) {
    var num int
    _, err := fmt.Sscan(input, &num)
    return num, err
}
