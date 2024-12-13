func removeDuplicateLines() {
    inputFile, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer inputFile.Close()

    lineSet := make(map[string]struct{})
    scanner := bufio.NewScanner(inputFile)
    for scanner.Scan() {
        lineSet[scanner.Text()] = struct{}{}
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

    for line := range lineSet {
        _, err := tempFile.WriteString(line + "\n")
        if err != nil {
            log.Println("Error writing to temporary file:", err)
        }
    }

    err = os.Rename("results_temp.txt", "results.txt")
    if err != nil {
        log.Println("Error replacing original file:", err)
    }

    fmt.Println("Removed duplicate lines successfully.")
}
