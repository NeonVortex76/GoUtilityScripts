func removeDuplicateOperations() {
    inputFile, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer inputFile.Close()

    tempFile, err := os.Create("temp_results.txt")
    if err != nil {
        log.Println("Error creating temp file:", err)
        return
    }
    defer tempFile.Close()

    scanner := bufio.NewScanner(inputFile)
    writer := bufio.NewWriter(tempFile)

    operationsSet := make(map[string]bool)

    for scanner.Scan() {
        line := scanner.Text()
        if _, exists := operationsSet[line]; !exists {
            operationsSet[line] = true
            _, err := writer.WriteString(line + "\n")
            if err != nil {
                log.Println("Error writing to temp file:", err)
            }
        }
    }

    err = writer.Flush()
    if err != nil {
        log.Println("Error flushing to temp file:", err)
    }

    os.Remove("results.txt")
    os.Rename("temp_results.txt", "results.txt")

    fmt.Println("Duplicate operations have been removed.")
}
