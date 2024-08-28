package main

import (
    "bufio"
    "fmt"
    "log"
    "math"
    "os"
    "strconv"
    "strings"
    "time"
)

var lastResult int

func main() {
    greet("Alice")
    fmt.Println("Hello, World!")

    file, err := os.OpenFile("results.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    for {
        var choice string
        fmt.Print("Choose an option: (1) New operation (2) View history (3) Clear history (4) Exit: ")
        fmt.Scanln(&choice)

        switch choice {
        case "1":
            performNewOperation(file)
        case "2":
            viewHistory()
        case "3":
            clearHistory()
        case "4":
            fmt.Println("Exiting program.")
            return
        default:
            fmt.Println("Invalid choice. Please enter 1, 2, 3, or 4.")
        }
        fmt.Println("---------")
    }
}

func greet(name string) {
    fmt.Printf("Hello, %s!\n", name)
}

func performNewOperation(file *os.File) {
    var input string
    fmt.Print("Enter operation (e.g., '10 + 20') or type 'last' to use last result: ")
    fmt.Scanln(&input)

    if strings.Contains(input, "last") {
        input = strings.Replace(input, "last", strconv.Itoa(lastResult), 1)
    }

    a, b, operation, err := parseInputWithOperation(input)
    if err != nil {
        log.Println("Invalid input:", err)
        return
    }

    result, err := performOperation(a, b, operation)
    if err != nil {
        log.Println("Error:", err)
        return
    }

    resultStr := fmt.Sprintf("Result: %v\n", result)
    fmt.Print(resultStr)

    lastResult = result.(int) // Сохраняем последний результат

    var comment string
    fmt.Print("Add a comment (optional): ")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    comment = scanner.Text()

    timestamp := time.Now().Format(time.RFC3339)
    entry := fmt.Sprintf("%s - %s - %v - Comment: %s\n", timestamp, input, result, comment)
    if _, err := file.WriteString(entry); err != nil {
        log.Println("Error writing to file:", err)
    }
}

func viewHistory() {
    file, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    fmt.Println("Calculation History:")
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
    }
}

func clearHistory() {
    err := os.Truncate("results.txt", 0)
    if err != nil {
        log.Println("Error clearing history:", err)
    } else {
        fmt.Println("History cleared.")
    }
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
