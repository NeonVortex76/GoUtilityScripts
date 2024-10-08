func deleteOperationsByKeyword() {
    inputFile, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer inputFile.Close()

    var keyword string
    fmt.Print("Enter the keyword to delete operations: ")
    fmt.Scanln(&keyword)

    tempFile, err := os.Create("temp_results.txt")
    if err != nil {
        log.Println("Error creating temp file:", err)
        return
    }
    defer tempFile.Close()

    scanner := bufio.NewScanner(inputFile)
    writer := bufio.NewWriter(tempFile)

    for scanner.Scan() {
        line := scanner.Text()
        if !strings.Contains(line, keyword) {
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

    fmt.Printf("Operations containing the keyword '%s' were deleted.\n", keyword)
}
