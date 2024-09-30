func deleteOperationsByKeyword() {
    file, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    var keyword string
    fmt.Print("Enter the keyword to delete operations: ")
    fmt.Scanln(&keyword)

    var updatedLines []string
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()
        if !strings.Contains(line, keyword) {
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

    fmt.Println("Operations containing the keyword have been deleted.")
}
