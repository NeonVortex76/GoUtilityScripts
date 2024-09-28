func saveHistoryToFile() {
    file, err := os.Open("results.txt")
    if err != nil {
        log.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    var newFileName string
    fmt.Print("Enter the new file name to save the history (e.g., backup.txt): ")
    fmt.Scanln(&newFileName)

    newFile, err := os.Create(newFileName)
    if err != nil {
        log.Println("Error creating new file:", err)
        return
    }
    defer newFile.Close()

    scanner := bufio.NewScanner(file)
    writer := bufio.NewWriter(newFile)

    for scanner.Scan() {
        line := scanner.Text()
        _, err := writer.WriteString(line + "\n")
        if err != nil {
            log.Println("Error writing to new file:", err)
            return
        }
    }

    if err := writer.Flush(); err != nil {
        log.Println("Error flushing data to new file:", err)
    }

    if err := scanner.Err(); err != nil {
        log.Println("Error reading file:", err)
    } else {
        fmt.Println("History successfully saved to", newFileName)
    }
}
