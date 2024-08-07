package main

import (
    "fmt"
)

func main() {
    greet("Alice")
    fmt.Println("Hello, World!")
    sum(5, 3)
    subtract(10, 4)
    multiply(7, 6)
    divide(20, 4)
}

func greet(name string) {
    fmt.Printf("Hello, %s!\n", name)
}

func sum(a, b int) {
    result := a + b
    fmt.Printf("The sum of %d and %d is %d\n", a, b, result)
}

func subtract(a, b int) {
    result := a - b
    fmt.Printf("The difference between %d and %d is %d\n", a, b, result)
}

func multiply(a, b int) {
    result := a * b
    fmt.Printf("The product of %d and %d is %d\n", a, b, result)
}

func divide(a, b int) {
    if b == 0 {
        fmt.Println("Error: Division by zero")
        return
    }
    result := a / b
    fmt.Printf("The division of %d by %d is %d\n", a, b, result)
}
