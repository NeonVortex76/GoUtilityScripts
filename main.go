func removeDuplicateOperations() {
    file, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    uniqueLines := make(map[string]bool)
    var updatedLines []string

    for scanner.Scan() {
        line := scanner.Text()
        if _, exists := uniqueLines[line]; !exists {
            uniqueLines[line] = true
            updatedLines = append(updatedLines, line)
        }
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
        return
    }

    file.Close()

    newFile, err := os.Create("results.txt")
    if err != nil {
        log.Println("Error creating file:", err)
        return
    }
    defer newFile.Close()

    writer := bufio.NewWriter(newFile)
    for _, line := range updatedLines {
        _, err := writer.WriteString(line + "\n")
        if err != nil {
            log.Println("Error writing to file:", err)
            return
        }
    }

    if err := writer.Flush(); err != nil {
        log.Println("Error flushing data to file:", err)
    }

    fmt.Println("Duplicate operations have been removed.")
}
