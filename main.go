func truncateFileAfterLines(maxLines int) {
    inputFile, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer inputFile.Close()

    tempFile, err := os.Create("results_temp.txt")
    if err != nil {
        log.Println("Error creating temporary file:", err)
        return
    }
    defer tempFile.Close()

    scanner := bufio.NewScanner(inputFile)
    lineCount := 0

    for scanner.Scan() {
        if lineCount >= maxLines {
            break
        }
        _, err := tempFile.WriteString(scanner.Text() + "\n")
        if err != nil {
            log.Println("Error writing to temporary file:", err)
        }
        lineCount++
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
        return
    }

    err = os.Rename("results_temp.txt", "results.txt")
    if err != nil {
        log.Println("Error replacing original file:", err)
    }

    fmt.Printf("Truncated file to %d lines successfully.\n", maxLines)
}
