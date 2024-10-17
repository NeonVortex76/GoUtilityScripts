func removeDuplicateOperations() {
    inputFile, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer inputFile.Close()

    operationsMap := make(map[string]bool)
    scanner := bufio.NewScanner(inputFile)
    uniqueOperations := []string{}

    for scanner.Scan() {
        line := scanner.Text()
        if !operationsMap[line] {
            operationsMap[line] = true
            uniqueOperations = append(uniqueOperations, line)
        }
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
        return
    }

    tempFile, err := os.Create("temp_results.txt")
    if err != nil {
        log.Println("Error creating temp file:", err)
        return
    }
    defer tempFile.Close()

    writer := bufio.NewWriter(tempFile)
    for _, operation := range uniqueOperations {
        _, err := writer.WriteString(operation + "\n")
        if err != nil {
            log.Println("Error writing to temp file:", err)
        }
    }

    err = writer.Flush()
    if err != nil {
        log.Println("Error flushing to temp file:", err)
    }

    os.Remove("results.txt")
    os.Rename("temp_results.txt", "results.txt")

    fmt.Println("Duplicate operations removed.")
}
