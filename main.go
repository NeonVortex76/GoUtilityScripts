func reverseOperationLines() {
    inputFile, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer inputFile.Close()

    var lines []string
    scanner := bufio.NewScanner(inputFile)

    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
    }

    fmt.Println("Reversed operation lines:")
    for i := len(lines) - 1; i >= 0; i-- {
        fmt.Println(lines[i])
    }
}
