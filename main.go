func countLines() {
    inputFile, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer inputFile.Close()

    lineCount := 0
    scanner := bufio.NewScanner(inputFile)
    for scanner.Scan() {
        lineCount++
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
        return
    }

    fmt.Printf("Total number of lines: %d\n", lineCount)
}
