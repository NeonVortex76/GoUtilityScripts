func renameKeywordInOperations() {
    inputFile, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer inputFile.Close()

    var oldKeyword, newKeyword string
    fmt.Print("Enter the keyword to rename: ")
    fmt.Scanln(&oldKeyword)
    fmt.Print("Enter the new keyword: ")
    fmt.Scanln(&newKeyword)

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
        modifiedLine := strings.ReplaceAll(line, oldKeyword, newKeyword)
        _, err := writer.WriteString(modifiedLine + "\n")
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

    fmt.Println("Keyword renamed in operations.")
}
