func countLineOccurrences() {
    inputFile, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer inputFile.Close()

    lineCounts := make(map[string]int)
    scanner := bufio.NewScanner(inputFile)

    for scanner.Scan() {
        line := scanner.Text()
        lineCounts[line]++
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
        return
    }

    tempFile, err := os.Create("results_temp.txt")
    if err != nil {
        log.Println("Error creating temporary file:", err)
        return
    }
    defer tempFile.Close()

    for line, count := range lineCounts {
        _, err := tempFile.WriteString(fmt.Sprintf("%s (x%d)\n", line, count))
        if err != nil {
            log.Println("Error writing to temporary file:", err)
        }
    }

    err = os.Rename("results_temp.txt", "results.txt")
    if err != nil {
        log.Println("Error replacing original file:", err)
    }

    fmt.Println("Line occurrences counted successfully.")
}
