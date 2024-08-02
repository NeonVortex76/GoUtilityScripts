package main

import (
    "fmt"
)

func main() {
    greet("Alice")
    fmt.Println("Hello, World!")
}

func greet(name string) {
    fmt.Printf("Hello, %s!\n", name)
}
