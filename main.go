package main

import (
    "fmt"
    "log"
    "math"
    "os"
    "strconv"
    "strings"
    "time"
)

func main() {
    greet("Alice")
    fmt.Println("Hello, World!")

    file, err := os.OpenFile("results.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    for {
        var input string
        fmt.Print("Enter operation (e.g., '10 + 20') or type 'exit' to quit: ")
        fmt.Scanln(&input)

        if strings.ToLower(input) == "exit" {
            fmt.Println("Exiting program.")
            break
        }

        a, b, operation, err := parseInputWithOperation(input)
        if err != nil {
            log.Println("Invalid input:", err)
            continue
        }

        result, err := performOperation(a, b, operation)
        if err != nil {
            log.Println("Error:", err)
            continue
        }

        resultStr := fmt.Sprintf("Result: %v\n", result)
        fmt.Print(resultStr)

        timestamp := time.Now().Format(time.RFC3339)
        entry := fmt.Sprintf("%s - %s - %v\n", timestamp, input, result)
        if _, err := file.WriteString(entry); err != nil {
            log.Println("Error writing to file:", err)
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

func divide(a, b int) (string, error) {
    if b == 0 {
        return "", fmt.Errorf("cannot divide by zero")
    }
    return fmt.Sprintf("%d", a/b), nil
}

func power(a, b int) float64 {
    return math.Pow(float64(a), float64(b))
}

func modulus(a, b int) int {
    return a % b
}

func parseInputWithOperation(input string) (int, int, string, error) {
    parts := strings.Fields(input)
    if len(parts) != 3 {
        return 0, 0, "", fmt.Errorf("input must contain exactly three parts")
    }

    a, err1 := strconv.Atoi(parts[0])
    operation := parts[1]
    b, err2 := strconv.Atoi(parts[2])

    if err1 != nil || err2 != nil {
        return 0, 0, "", fmt.Errorf("invalid numbers")
    }

    return a, b, operation, nil
}

func performOperation(a, b int, operation string) (interface{}, error) {
    switch operation {
    case "+":
        return sum(a, b), nil
    case "-":
        return subtract(a, b), nil
    case "*":
        return multiply(a, b), nil
    case "/":
        return divide(a, b)
    case "^":
        return power(a, b), nil
    case "%":
        return modulus(a, b), nil
    default:
        return nil, fmt.Errorf("invalid operation")
    }
}
